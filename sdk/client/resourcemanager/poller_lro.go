package resourcemanager

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/client/pollers"
)

var _ pollers.PollerType = &longRunningOperationPoller{}

type longRunningOperationPoller struct {
	client               *client.Client
	count                int
	initialRetryDuration time.Duration
	originalUrl          *url.URL
	pollingUrl           *url.URL
}

func longRunningOperationPollerFromResponse(resp *client.Response, client *client.Client) (*longRunningOperationPoller, error) {
	poller := longRunningOperationPoller{
		client:               client,
		initialRetryDuration: 10 * time.Second,
	}

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
	poller.pollingUrl = u
	if endpoint, err := url.Parse(string(client.Endpoint)); err == nil && u.Host != endpoint.Host {
		return nil, fmt.Errorf("unsupported polling URL %q: client endpoint is different", pollingUrl)
	}

	if resp.Request != nil {
		poller.originalUrl = resp.Request.URL
	}

	if s, ok := resp.Header["Retry-After"]; ok {
		if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
			poller.initialRetryDuration = time.Second * time.Duration(sleep)
		}
	}

	return &poller, nil
}

func (p *longRunningOperationPoller) Poll(ctx context.Context) (result *pollers.PollResult, err error) {
	p.count++

	if p.pollingUrl == nil {
		return nil, fmt.Errorf("internal error: cannot poll without a pollingUrl")
	}

	reqOpts := client.RequestOptions{
		ContentType: "application/json; charset=utf-8",
		ExpectedStatusCodes: []int{
			http.StatusOK,
			http.StatusCreated,
			http.StatusAccepted,
			http.StatusNoContent,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: nil,
		Path:          p.pollingUrl.Path,
	}

	req, err := p.client.NewRequest(ctx, reqOpts)
	if err != nil {
		return nil, fmt.Errorf("building request for long-running-operation: %+v", err)
	}
	req.URL.RawQuery = p.pollingUrl.RawQuery

	// disable retries since the poller handles this for us
	req.RetryFunc = func(response *http.Response, data *odata.OData) (bool, error) {
		return false, nil
	}

	// Custom RetryFunc to inspect the operation payload and check the status
	req.RetryFunc = client.RequestRetryAny(defaultRetryFunctions...)

	result.HttpResponse, err = req.Execute()
	if err != nil {
		return nil, err
	}

	if result.HttpResponse != nil {
		var respBody []byte
		respBody, err = io.ReadAll(result.HttpResponse.Body)
		if err != nil {
			err = fmt.Errorf("parsing response body: %+v", err)
			return
		}
		result.HttpResponse.Body.Close()

		contentType := result.HttpResponse.Header.Get("Content-Type")

		var op operationResult
		if strings.Contains(strings.ToLower(contentType), "application/json") {
			if err = json.Unmarshal(respBody, &op); err != nil {
				err = fmt.Errorf("unmarshalling response body: %+v", err)
				return
			}
		} else {
			return nil, fmt.Errorf("internal-error: polling support for the Content-Type %q was not implemented: %+v", contentType, err)
		}

		// TODO: raising an error if this is Cancelled or Failed

		statuses := map[status]pollers.PollingStatus{
			statusCanceled:   pollers.PollingStatusCancelled,
			statusCancelled:  pollers.PollingStatusCancelled,
			statusFailed:     pollers.PollingStatusFailed,
			statusInProgress: pollers.PollingStatusInProgress,
			statusSucceeded:  pollers.PollingStatusSucceeded,
		}
		for k, v := range statuses {
			if strings.EqualFold(string(op.Status), string(k)) {
				result.Status = v
				return
			}
		}
	}

	return
}

type operationResult struct {
	Name      *string    `json:"name"`
	StartTime *time.Time `json:"startTime"`
	Status    status     `json:"status"`
}

type status string

const (
	statusCanceled   status = "Canceled"
	statusCancelled  status = "Cancelled"
	statusFailed     status = "Failed"
	statusInProgress status = "InProgress"
	statusSucceeded  status = "Succeeded"
)
