package resourcemanager

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"github.com/hashicorp/go-azure-sdk/client/base"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/odata"
)

type Client struct {
	// Client is the base client.Client for making API requests
	Client *base.Client

	// DisableResourceProviderRegistration prevents the client from attempting to register an RP for the subscription whenever a request
	// fails due to absent subscription registration. When false (the default), Client.DisableRetries must also be false for this to work.
	DisableResourceProviderRegistration bool

	// FallbackPollingDuration is the maximum length of time for which LRO polling can take place. Used primarily when no deadline is provided when polling.
	FallbackPollingDuration time.Duration
}

func NewResourceManagerClient(endpoint environments.ApiEndpoint) *Client {
	return &Client{
		Client:                  base.NewClient(endpoint),
		FallbackPollingDuration: 60 * time.Minute,
	}
}

func (c *Client) NewDeleteRequest(ctx context.Context, path string, apiVersion string, odata odata.Query) (req *base.Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodDelete, "application/json; charset=utf-8", path, base.QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		req.RetryFunc = base.RequestRetryAny(base.RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK, http.StatusNoContent}
	}

	return
}

func (c *Client) NewGetRequest(ctx context.Context, path string, apiVersion string, odata odata.Query) (req *base.Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodGet, "application/json; charset=utf-8", path, base.QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		req.RetryFunc = base.RequestRetryAny(base.RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK}
	}

	return
}

func (c *Client) NewPatchRequest(ctx context.Context, path string, apiVersion string, odata odata.Query, model interface{}) (req *base.Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodPatch, "application/json; charset=utf-8", path, base.QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.RetryFunc = base.RequestRetryAny(base.RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK, http.StatusNoContent}
	}

	return
}

func (c *Client) NewPostRequest(ctx context.Context, path string, apiVersion string, odata odata.Query, model interface{}) (req *base.Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodPost, "application/json; charset=utf-8", path, base.QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}

		retryFuncs := []base.RequestRetryFunc{base.RetryOn404ConsistencyFailureFunc}
		if !c.DisableResourceProviderRegistration {
			retryFuncs = append(retryFuncs, retryWithRegistration)
		}
		req.RetryFunc = base.RequestRetryAny(retryFuncs...)
	}

	return
}

func (c *Client) NewPutRequest(ctx context.Context, path string, apiVersion string, odata odata.Query, model interface{}) (req *base.Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodPut, "application/json; charset=utf-8", path, base.QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusAccepted, http.StatusNoContent}

		retryFuncs := []base.RequestRetryFunc{base.RetryOn404ConsistencyFailureFunc}
		if !c.DisableResourceProviderRegistration {
			retryFuncs = append(retryFuncs, retryWithRegistration)
		}
		req.RetryFunc = base.RequestRetryAny(retryFuncs...)
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
