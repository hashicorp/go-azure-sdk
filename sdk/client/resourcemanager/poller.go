package resourcemanager

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type Operation struct {
	Name      *string    `json:"name"`
	StartTime *time.Time `json:"startTime"`
	Status    status     `json:"status"`
}

type status string

const (
	statusCanceled   status = "Canceled"
	statusFailed     status = "Failed"
	statusInProgress status = "InProgress"
	statusSucceeded  status = "Succeeded"
)

type LongRunningPoller struct {
	// Response is the latest Response
	Response *client.Response

	ctx               context.Context
	client            *Client
	count             int
	initialRetryAfter time.Time
	originalUrl       *url.URL
	pollingUrl        *url.URL
}

// NewPollerFromResponse creates a new LongRunningPoller from the HTTP Response
func NewPollerFromResponse(ctx context.Context, resp *client.Response, client *Client) (*LongRunningPoller, error) {
	pollingUrl := resp.Header.Get(http.CanonicalHeaderKey("Azure-AsyncOperation"))
	if pollingUrl == "" {
		pollingUrl = resp.Header.Get("Location")
	}
	if pollingUrl == "" {
		return nil, fmt.Errorf("no polling URL found in response")
	}

	u, err := url.Parse(pollingUrl)
	if err != nil {
		return nil, fmt.Errorf("invalid polling URL %q in response: %v", pollingUrl, err)
	}
	if !u.IsAbs() {
		return nil, fmt.Errorf("invalid polling URL %q in response: URL was not absolute", pollingUrl)
	}
	if e, err := url.Parse(string(client.Client.Endpoint)); err == nil && u.Host != e.Host {
		return nil, fmt.Errorf("unsupported polling URL %q: client endpoint is different", pollingUrl)
	}

	var originalUrl *url.URL
	if resp.Request != nil {
		originalUrl = resp.Request.URL
	}

	poller := LongRunningPoller{
		ctx:               ctx,
		client:            client,
		initialRetryAfter: time.Now().Add(10 * time.Second),
		originalUrl:       originalUrl,
		pollingUrl:        u,
	}

	if s, ok := resp.Header["Retry-After"]; ok {
		if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
			poller.initialRetryAfter = time.Now().Add(time.Second * time.Duration(sleep))
		}
	}

	return &poller, nil
}

// PollUntilDone polls until the operation is completed, cancelled or failed
func (p *LongRunningPoller) PollUntilDone() error {
	if _, hasDeadline := p.ctx.Deadline(); !hasDeadline && p.client.FallbackPollingDuration != 0 {
		var cancel context.CancelFunc
		p.ctx, cancel = context.WithTimeout(p.ctx, p.client.FallbackPollingDuration)
		defer cancel()
	}

	// TODO: make this deadline-aware
	time.Sleep(time.Until(p.initialRetryAfter))

	req, err := p.client.NewRequest(p.ctx, http.MethodGet, "application/json; charset=utf-8", p.pollingUrl.Path)
	if err != nil {
		return err
	}
	req.URL.RawQuery = p.pollingUrl.RawQuery
	req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}

	// Custom RetryFunc to inspect the operation payload and check the status
	req.RetryFunc = func(r *http.Response, o *odata.OData) (bool, error) {
		if strings.HasPrefix(strings.ToLower(r.Header.Get("Content-Type")), "application/json") {
			// Read the response body and close it
			respBody, err := io.ReadAll(r.Body)
			if err != nil {
				return false, fmt.Errorf("could not parse response body")
			}
			r.Body.Close()

			// Trim away a BOM if present
			respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

			// Reassign the response body as downstream code may expect it
			r.Body = io.NopCloser(bytes.NewBuffer(respBody))

			// Unmarshal and inspect status
			var op Operation
			if err := json.Unmarshal(respBody, &op); err != nil {
				return false, err
			}

			switch op.Status {
			case statusCanceled, statusFailed:
				return false, fmt.Errorf("async operation cancelled")
			case statusSucceeded:
				return false, nil
			}
		}

		return true, nil
	}

	// Retries handled by retryablehttp
	p.Response, err = req.Execute()
	return err
}
