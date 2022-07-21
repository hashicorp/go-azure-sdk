package client

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/odata"
	"net/http"
)

type ResourceManagerClient struct {
	// SubscriptionID not yet implemented
	SubscriptionID string

	// DisableResourceProviderRegistration prevents the client from attempting to register an RP for the subscription whenever a request
	// fails due to absent subscription registration. When false (the default), Client.DisableRetries must also be false for this to work.
	DisableResourceProviderRegistration bool

	Client *Client
}

func NewResourceManagerClient(endpoint environments.ApiEndpoint, subscriptionId string) *ResourceManagerClient {
	return &ResourceManagerClient{
		SubscriptionID: subscriptionId,
		Client:         NewClient(endpoint),
	}
}

func (c *ResourceManagerClient) NewDeleteRequest(ctx context.Context, path string, apiVersion string, odata odata.Query) (req *Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodDelete, "application/json; charset=utf-8", path, QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		req.RetryFunc = RequestRetryAny(RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK, http.StatusNoContent}
	}

	return
}

func (c *ResourceManagerClient) NewGetRequest(ctx context.Context, path string, apiVersion string, odata odata.Query) (req *Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodGet, "application/json; charset=utf-8", path, QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		req.RetryFunc = RequestRetryAny(RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK}
	}

	return
}

func (c *ResourceManagerClient) NewPatchRequest(ctx context.Context, path string, apiVersion string, odata odata.Query, model interface{}) (req *Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodPatch, "application/json; charset=utf-8", path, QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.RetryFunc = RequestRetryAny(RetryOn404ConsistencyFailureFunc)
		req.ValidStatusCodes = []int{http.StatusOK, http.StatusNoContent}
	}

	return
}

func (c *ResourceManagerClient) NewPostRequest(ctx context.Context, path string, apiVersion string, odata odata.Query, model interface{}) (req *Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodPost, "application/json; charset=utf-8", path, QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}

		retryFuncs := []RequestRetryFunc{RetryOn404ConsistencyFailureFunc}
		if !c.DisableResourceProviderRegistration {
			retryFuncs = append(retryFuncs, retryWithRegistration)
		}
		req.RetryFunc = RequestRetryAny(retryFuncs...)
	}

	return
}

func (c *ResourceManagerClient) NewPutRequest(ctx context.Context, path string, apiVersion string, odata odata.Query, model interface{}) (req *Request, err error) {
	req, err = c.Client.NewRequest(ctx, http.MethodPut, "application/json; charset=utf-8", path, QueryParams{"api-version": apiVersion}, odata)

	if req != nil {
		if model != nil {
			if err = req.Marshal(model); err != nil {
				return req, err
			}
		}

		req.ValidStatusCodes = []int{http.StatusOK, http.StatusCreated, http.StatusNoContent}

		retryFuncs := []RequestRetryFunc{RetryOn404ConsistencyFailureFunc}
		if !c.DisableResourceProviderRegistration {
			retryFuncs = append(retryFuncs, retryWithRegistration)
		}
		req.RetryFunc = RequestRetryAny(retryFuncs...)
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
