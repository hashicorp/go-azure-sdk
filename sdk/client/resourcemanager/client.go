package resourcemanager

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
)

type Options interface {
	ToHeaders() *Headers
	ToOData() *odata.Query
	ToQuery() *QueryParams
}

type Headers map[string]string

func (h Headers) Append(he http.Header) http.Header {
	for k, v := range h {
		he.Set(k, v)
	}
	return he
}

func (h Headers) Headers() http.Header {
	return h.Append(make(http.Header))
}

type QueryParams map[string]string

func (q QueryParams) Values() url.Values {
	va := make(url.Values)
	for k, v := range q {
		va.Set(k, v)
	}
	return va
}

type Client struct {
	// DisableResourceProviderRegistration prevents the client from attempting to register an RP for the subscription whenever a request
	// fails due to absent subscription registration. When false (the default), Client.DisableRetries must also be false for this to work.
	DisableResourceProviderRegistration bool

	// FallbackPollingDuration is the maximum length of time for which LRO polling can take place. Used primarily when no deadline is provided when polling.
	FallbackPollingDuration time.Duration

	*client.Client
}

func NewResourceManagerClient(endpoint environments.ApiEndpoint) *Client {
	return &Client{
		Client:                  client.NewClient(endpoint),
		FallbackPollingDuration: 60 * time.Minute,
	}
}

func (c *Client) newRequest(ctx context.Context, method string, contentType string, path string, apiVersion string, options Options) (req *client.Request, err error) {
	req, err = c.Client.NewRequest(ctx, method, contentType, path)
	if err != nil {
		return
	}

	req.Client = c
	query := url.Values{}
	query.Set("api-version", apiVersion)

	if options != nil {
		if h := options.ToHeaders(); h != nil {
			req.Header = h.Append(req.Header)
		}

		if q := options.ToQuery(); q != nil {
			query = q.Values()
		}

		if o := options.ToOData(); o != nil {
			req.Header = o.AppendHeaders(req.Header)
			query = o.AppendValues(query)
		}
	}

	req.URL.RawQuery = query.Encode()

	return
}

func (c *Client) NewDeleteRequest(ctx context.Context, path string, apiVersion string, options Options) (req *client.Request, err error) {
	req, err = c.newRequest(ctx, http.MethodDelete, "application/json; charset=utf-8", path, apiVersion, options)

	if req != nil {
		req.RetryFunc = client.RequestRetryAny(client.RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK, http.StatusNoContent}
	}

	return
}

func (c *Client) NewGetRequest(ctx context.Context, path string, apiVersion string, options Options) (req *client.Request, err error) {
	req, err = c.newRequest(ctx, http.MethodGet, "application/json; charset=utf-8", path, apiVersion, options)

	if req != nil {
		req.RetryFunc = client.RequestRetryAny(client.RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK}
	}

	return
}

func (c *Client) NewPatchRequest(ctx context.Context, path string, apiVersion string, options Options, model interface{}) (req *client.Request, err error) {
	req, err = c.newRequest(ctx, http.MethodPatch, "application/json; charset=utf-8", path, apiVersion, options)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.RetryFunc = client.RequestRetryAny(client.RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK, http.StatusNoContent}
	}

	return
}

func (c *Client) NewPostRequest(ctx context.Context, path string, apiVersion string, options Options, model interface{}) (req *client.Request, err error) {
	req, err = c.newRequest(ctx, http.MethodPost, "application/json; charset=utf-8", path, apiVersion, options)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}

		retryFuncs := []client.RequestRetryFunc{client.RetryOn404ConsistencyFailureFunc}
		if !c.DisableResourceProviderRegistration {
			retryFuncs = append(retryFuncs, retryWithRegistration)
		}
		req.RetryFunc = client.RequestRetryAny(retryFuncs...)
	}

	return
}

func (c *Client) NewPutRequest(ctx context.Context, path string, apiVersion string, options Options, model interface{}) (req *client.Request, err error) {
	req, err = c.newRequest(ctx, http.MethodPut, "application/json; charset=utf-8", path, apiVersion, options)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}

		retryFuncs := []client.RequestRetryFunc{client.RetryOn404ConsistencyFailureFunc}
		if !c.DisableResourceProviderRegistration {
			retryFuncs = append(retryFuncs, retryWithRegistration)
		}
		req.RetryFunc = client.RequestRetryAny(retryFuncs...)
	}

	return
}

func retryWithRegistration(r *http.Response, o *odata.OData) (bool, error) {
	if o.Error != nil && o.Error.Code != nil && *o.Error.Code == "MissingSubscriptionRegistration" {
		if err := register(); err != nil {
			return false, err
		}
		return true, nil
	}

	return false, nil
}

func register() error {
	return fmt.Errorf("not implemented")
}
