package resourcemanager

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/client"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"net/url"
)

var _ client.BaseClient = &Client{}

type Client struct {
	*client.Client

	// apiVersion specifies the version of the API being used, which (by design) will be consistent across a client
	// as we intentionally split out multiple API Versions into different clients, rather than using composite API
	// Versions/packages which can cause confusion about which version is being used.
	apiVersion string
}

func NewResourceManagerClient(endpoint environments.ApiEndpoint, apiVersion string) *Client {
	return &Client{
		// TODO: we need a means of setting additional headers (e.g. the Correlation ID etc)
		Client:     client.NewClient(endpoint),
		apiVersion: apiVersion,
	}
}

//func (c *Client) NewDeleteRequest(ctx context.Context, path, contentType string, optionsObject Options) (*client.Request, error) {
//	if _, ok := ctx.Deadline(); !ok {
//		return nil, fmt.Errorf("the context used must have a deadline attached for polling purposes, but got no deadline")
//	}
//
//	opts := RequestOptions{
//		ContentType: contentType,
//		ExpectedStatusCodes: []int{
//			http.StatusOK,
//			http.StatusNoContent,
//		},
//		HttpMethod:    http.MethodDelete,
//		OptionsObject: optionsObject,
//		Path:          path,
//	}
//	req, err := c.NewRequest(ctx, opts)
//	if err != nil {
//		return nil, err
//	}
//	return req, nil
//}
//
//func (c *Client) NewGetRequest(ctx context.Context, path, contentType string, optionsObject Options) (*client.Request, error) {
//	if _, ok := ctx.Deadline(); !ok {
//		return nil, fmt.Errorf("the context used must have a deadline attached for polling purposes, but got no deadline")
//	}
//
//	opts := RequestOptions{
//		ContentType: contentType,
//		ExpectedStatusCodes: []int{
//			http.StatusOK,
//		},
//		HttpMethod:    http.MethodGet,
//		OptionsObject: optionsObject,
//		Path:          path,
//	}
//	req, err := c.NewRequest(ctx, opts)
//	if err != nil {
//		return nil, err
//	}
//	return req, nil
//}
//
//func (c *Client) NewPatchRequest(ctx context.Context, path, contentType string, optionsObject Options, payload interface{}) (*client.Request, error) {
//	if _, ok := ctx.Deadline(); !ok {
//		return nil, fmt.Errorf("the context used must have a deadline attached for polling purposes, but got no deadline")
//	}
//
//	opts := RequestOptions{
//		ContentType: contentType,
//		ExpectedStatusCodes: []int{
//			http.StatusOK,
//			http.StatusNoContent,
//		},
//		HttpMethod:    http.MethodPatch,
//		OptionsObject: optionsObject,
//		Path:          path,
//	}
//	req, err := c.NewRequest(ctx, opts)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := req.Marshal(payload); err != nil {
//		return nil, fmt.Errorf("marshaling payload for request: %+v", err)
//	}
//
//	return req, nil
//}
//
//func (c *Client) NewPostRequest(ctx context.Context, path, contentType string, optionsObject Options, payload interface{}) (*client.Request, error) {
//	if _, ok := ctx.Deadline(); !ok {
//		return nil, fmt.Errorf("the context used must have a deadline attached for polling purposes, but got no deadline")
//	}
//
//	opts := RequestOptions{
//		ContentType: contentType,
//		ExpectedStatusCodes: []int{
//			http.StatusOK,
//			http.StatusCreated,
//			http.StatusNoContent,
//		},
//		HttpMethod:    http.MethodPost,
//		OptionsObject: optionsObject,
//		Path:          path,
//	}
//	req, err := c.NewRequest(ctx, opts)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := req.Marshal(payload); err != nil {
//		return nil, fmt.Errorf("marshaling payload for request: %+v", err)
//	}
//
//	return req, nil
//}
//
//func (c *Client) NewPutRequest(ctx context.Context, path, contentType string, optionsObject Options, payload interface{}) (*client.Request, error) {
//	if _, ok := ctx.Deadline(); !ok {
//		return nil, fmt.Errorf("the context used must have a deadline attached for polling purposes, but got no deadline")
//	}
//
//	opts := RequestOptions{
//		ContentType: contentType,
//		ExpectedStatusCodes: []int{
//			http.StatusOK,
//			http.StatusCreated,
//			http.StatusAccepted,
//			http.StatusNoContent,
//		},
//		HttpMethod:    http.MethodPut,
//		OptionsObject: optionsObject,
//		Path:          path,
//	}
//	req, err := c.NewRequest(ctx, opts)
//	if err != nil {
//		return nil, err
//	}
//
//	if err := req.Marshal(payload); err != nil {
//		return nil, fmt.Errorf("marshaling payload for request: %+v", err)
//	}
//
//	return req, nil
//}

func (c *Client) NewRequest(ctx context.Context, input client.RequestOptions) (*client.Request, error) {
	if _, ok := ctx.Deadline(); !ok {
		return nil, fmt.Errorf("the context used must have a deadline attached for polling purposes, but got no deadline")
	}
	if err := input.Validate(); err != nil {
		return nil, fmt.Errorf("pre-validating request payload: %+v", err)
	}

	req, err := c.Client.NewRequest(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("building %s request: %+v", input.HttpMethod, err)
	}

	req.Client = c
	query := url.Values{}
	query.Set("api-version", c.apiVersion)

	if input.OptionsObject != nil {
		if h := input.OptionsObject.ToHeaders(); h != nil {
			req.Header = h.Append(req.Header)
		}

		if q := input.OptionsObject.ToQuery(); q != nil {
			query = q.Values()
		}

		if o := input.OptionsObject.ToOData(); o != nil {
			req.Header = o.AppendHeaders(req.Header)
			query = o.AppendValues(query)
		}
	}

	req.URL.RawQuery = query.Encode()
	req.RetryFunc = client.RequestRetryAny(defaultRetryFunctions...)
	req.ValidStatusCodes = input.ExpectedStatusCodes

	return req, nil
}
