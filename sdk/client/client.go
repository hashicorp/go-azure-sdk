// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"math"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strconv"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/accept"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-retryablehttp"
)

// RetryOn404ConsistencyFailureFunc can be used to retry a request when a 404 response is received
func RetryOn404ConsistencyFailureFunc(resp *http.Response, _ *odata.OData) (bool, error) {
	return resp != nil && resp.StatusCode == http.StatusNotFound, nil
}

// RequestRetryAny wraps multiple RequestRetryFuncs and calls them in turn, returning true if any func returns true
func RequestRetryAny(retryFuncs ...RequestRetryFunc) func(resp *http.Response, o *odata.OData) (bool, error) {
	return func(resp *http.Response, o *odata.OData) (retry bool, err error) {
		for _, retryFunc := range retryFuncs {
			if retryFunc != nil {
				retry, err = retryFunc(resp, o)
				if err != nil {
					return
				}
				if retry {
					return
				}
			}
		}
		return false, nil
	}
}

// RequestRetryAll wraps multiple RequestRetryFuncs and calls them in turn, only returning true if all funcs return true
func RequestRetryAll(retryFuncs ...RequestRetryFunc) func(resp *http.Response, o *odata.OData) (bool, error) {
	return func(resp *http.Response, o *odata.OData) (retry bool, err error) {
		for _, retryFunc := range retryFuncs {
			if retryFunc != nil {
				retry, err = retryFunc(resp, o)
				if err != nil {
					return
				}
				if !retry {
					return
				}
			}
		}
		return true, nil
	}
}

// RetryableErrorHandler ensures that the response is returned after exhausting retries for a request
// We mustn't return an error here, or net/http will not return the response
func RetryableErrorHandler(resp *http.Response, _ error, _ int) (*http.Response, error) {
	return resp, nil
}

// Request embeds *http.Request and adds useful metadata
type Request struct {
	RetryFunc        RequestRetryFunc
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc

	Client BaseClient

	// Embed *http.Request so that we can send this to an *http.Client
	*http.Request
}

func (r *Request) Marshal(payload interface{}) error {
	contentType := strings.ToLower(r.Header.Get("Content-Type"))

	if strings.Contains(contentType, "application/json") {
		body, err := json.Marshal(payload)
		if err == nil {
			r.ContentLength = int64(len(body))
			r.Body = io.NopCloser(bytes.NewReader(body))
		}
		return nil
	}

	if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		body, err := xml.Marshal(payload)
		if err == nil {
			r.ContentLength = int64(len(body))
			r.Body = io.NopCloser(bytes.NewReader(body))
		}
		return nil
	}

	if strings.Contains(contentType, "application/octet-stream") {
		v, ok := payload.([]byte)
		if !ok {
			return fmt.Errorf("internal-error: `payload` must be []byte but got %+v", payload)
		}

		r.ContentLength = int64(len(v))
		r.Body = io.NopCloser(bytes.NewReader(v))
	}

	return fmt.Errorf("internal-error: unimplemented marshal function for content type %q", contentType)
}

func (r *Request) Execute(ctx context.Context) (*Response, error) {
	return r.Client.Execute(ctx, r)
}

func (r *Request) ExecutePaged(ctx context.Context) (*Response, error) {
	return r.Client.ExecutePaged(ctx, r)
}

func (r *Request) IsIdempotent() bool {
	switch strings.ToUpper(r.Method) {
	case http.MethodGet, http.MethodHead, http.MethodOptions:
		return true
	}
	return false
}

// Response embeds *http.Response and adds useful methods
type Response struct {
	OData *odata.OData

	// Embed *http.Response
	*http.Response
}

func (r *Response) Unmarshal(model interface{}) error {
	if model == nil {
		return fmt.Errorf("model was nil")
	}

	contentType := strings.ToLower(r.Header.Get("Content-Type"))
	if contentType == "" {
		// some APIs (e.g. Storage Data Plane) don't return a content type... so we'll assume from the Accept header
		acc, err := accept.FromString(r.Request.Header.Get("Accept"))
		if err != nil {
			if preferred := acc.FirstChoice(); preferred != nil {
				contentType = preferred.ContentType
			}
		}
		if contentType == "" {
			// fall back on request media type
			contentType = strings.ToLower(r.Request.Header.Get("Content-Type"))
		}
	}
	if strings.Contains(contentType, "application/json") {
		// Read the response body and close it
		respBody, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("parsing response body: %+v", err)
		}
		r.Body.Close()

		// Trim away a BOM if present
		respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

		// Unmarshal into provided model
		if err := json.Unmarshal(respBody, model); err != nil {
			return fmt.Errorf("unmarshaling response body: %+v", err)
		}

		// Reassign the response body as downstream code may expect it
		r.Body = io.NopCloser(bytes.NewBuffer(respBody))
	}

	if strings.Contains(contentType, "application/xml") || strings.Contains(contentType, "text/xml") {
		// Read the response body and close it
		respBody, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("could not parse response body")
		}
		r.Body.Close()

		// Trim away a BOM if present
		respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

		// Unmarshal into provided model
		if err := xml.Unmarshal(respBody, model); err != nil {
			return err
		}

		// Reassign the response body as downstream code may expect it
		r.Body = io.NopCloser(bytes.NewBuffer(respBody))
	}

	if strings.Contains(contentType, "application/octet-stream") {
		if _, ok := model.([]byte); !ok {
			return fmt.Errorf("internal-error: `model` must be []byte but got %+v", model)
		}

		// Read the response body and close it
		respBody, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("could not parse response body")
		}
		r.Body.Close()

		// Trim away a BOM if present
		respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

		// copy the byte stream across
		model = respBody

		// Reassign the response body as downstream code may expect it
		r.Body = io.NopCloser(bytes.NewBuffer(respBody))
	}

	return nil
}

// Client is a base client to be used by API-specific clients. It satisfies the BaseClient interface.
type Client struct {
	// BaseUri is the base endpoint for this API.
	BaseUri string

	// UserAgent is the HTTP user agent string to send in requests.
	UserAgent string

	// CorrelationId is a custom correlation ID which can be added to requests for tracing purposes
	CorrelationId string

	// Authorizer is anything that can provide an access token with which to authorize requests.
	Authorizer auth.Authorizer

	// DisableRetries prevents the client from reattempting failed requests (which it does to work around eventual consistency issues).
	// This does not impact handling of retries related to rate limiting, which are always performed.
	DisableRetries bool

	// RequestMiddlewares is a slice of functions that are called in order before a request is sent
	RequestMiddlewares *[]RequestMiddleware

	// ResponseMiddlewares is a slice of functions that are called in order before a response is parsed and returned
	ResponseMiddlewares *[]ResponseMiddleware
}

// NewClient returns a new Client configured with sensible defaults
func NewClient(baseUri string, serviceName, apiVersion string) *Client {
	segments := []string{
		"Go-http-Client/1.1",
		fmt.Sprintf("%s/%s", serviceName, apiVersion),
	}
	return &Client{
		BaseUri:   baseUri,
		UserAgent: fmt.Sprintf("HashiCorp/go-azure-sdk (%s)", strings.Join(segments, " ")),
	}
}

// NewRequest configures a new *Request
func (c *Client) NewRequest(ctx context.Context, input RequestOptions) (*Request, error) {
	req := (&http.Request{}).WithContext(ctx)

	req.Method = input.HttpMethod

	req.Header = make(http.Header)
	req.Header.Add("Content-Type", input.ContentType)

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}
	if c.CorrelationId != "" {
		req.Header.Add("X-Ms-Correlation-Request-Id", c.CorrelationId)
	}

	path := strings.TrimPrefix(input.Path, "/")
	u, err := url.ParseRequestURI(fmt.Sprintf("%s/%s", c.BaseUri, path))
	if err != nil {
		return nil, err
	}

	req.Host = u.Host
	req.URL = u

	ret := Request{
		Client:           c,
		Request:          req,
		ValidStatusCodes: input.ExpectedStatusCodes,
	}

	return &ret, nil
}

// Execute is used by the package to send an HTTP request to the API
func (c *Client) Execute(ctx context.Context, req *Request) (*Response, error) {
	if req.Request == nil {
		return nil, fmt.Errorf("req.Request was nil")
	}

	// at this point we're ready to send the HTTP Request, as such let's get the Authorization token
	// and add that to the request
	if c.Authorizer != nil {
		token, err := c.Authorizer.Token(ctx, req.Request)
		if err != nil {
			return nil, err
		}
		token.SetAuthHeader(req.Request)
	}

	var err error

	// Check we can read the request body and set a default empty body
	var reqBody []byte
	if req.Body != nil {
		reqBody, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, fmt.Errorf("reading request body: %v", err)
		}
		req.Body = io.NopCloser(bytes.NewBuffer(reqBody))
		req.Header.Set("Content-Length", fmt.Sprintf("%d", req.ContentLength))
	}

	// Instantiate a RetryableHttp client and configure its CheckRetry func
	r := c.retryableClient(func(ctx context.Context, r *http.Response, err error) (bool, error) {
		// First check for badly malformed responses
		if r == nil {
			if req.IsIdempotent() {
				return true, nil
			}
			return false, fmt.Errorf("HTTP response was nil; connection may have been reset")
		}

		// Eventual consistency checks
		if !c.DisableRetries {
			if r.StatusCode == http.StatusFailedDependency {
				return true, nil
			}

			o, err := odata.FromResponse(r)
			if err != nil {
				return false, err
			}

			if f := req.RetryFunc; f != nil {
				shouldRetry, err := f(r, o)
				if err != nil || shouldRetry {
					return shouldRetry, err
				}
			}
		}

		// Fall back to default retry policy to handle rate limiting, server errors etc.
		return retryablehttp.DefaultRetryPolicy(ctx, r, err)
	})

	// Derive an *http.Client for sending the request
	client := r.StandardClient()

	// Configure any RequestMiddlewares
	if c.RequestMiddlewares != nil {
		for _, m := range *c.RequestMiddlewares {
			r, err := m(req.Request)
			if err != nil {
				return nil, err
			}
			req.Request = r
		}
	}

	// Send the request
	resp := &Response{}
	resp.Response, err = client.Do(req.Request)
	if err != nil {
		return resp, err
	}
	if resp.Response == nil {
		return resp, fmt.Errorf("HTTP response was nil; connection may have been reset")
	}

	// Configure any ResponseMiddlewares
	if c.ResponseMiddlewares != nil {
		for _, m := range *c.ResponseMiddlewares {
			r, err := m(req.Request, resp.Response)
			if err != nil {
				return resp, err
			}
			resp.Response = r
		}
	}

	// Extract OData from response
	var o *odata.OData
	resp.OData, err = odata.FromResponse(resp.Response)
	if err != nil {
		return resp, err
	}
	if resp == nil {
		return resp, fmt.Errorf("nil response received")
	}

	// Determine whether response status is valid
	if !containsStatusCode(req.ValidStatusCodes, resp.StatusCode) {
		if f := req.ValidStatusFunc; f != nil && f(resp.Response, o) {
			return resp, nil
		}

		// Determine suitable error text
		var errText string
		switch {
		case o != nil && o.Error != nil && o.Error.String() != "":
			errText = fmt.Sprintf("error: %s", o.Error)

		default:
			defer resp.Body.Close()

			respBody, err := io.ReadAll(resp.Body)
			if err != nil {
				return resp, fmt.Errorf("unexpected status %d, could not read response body", resp.StatusCode)
			}
			if len(respBody) == 0 {
				return resp, fmt.Errorf("unexpected status %d received with no body", resp.StatusCode)
			}

			errText = fmt.Sprintf("response: %s", respBody)
		}

		return resp, fmt.Errorf("unexpected status %d with %s", resp.StatusCode, errText)
	}

	return resp, nil
}

// ExecutePaged automatically pages through the results of Execute
func (c *Client) ExecutePaged(ctx context.Context, req *Request) (*Response, error) {
	// Perform the request
	resp, err := c.Execute(ctx, req)
	if err != nil {
		return resp, err
	}

	// Check for json content before handling pagination
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if !strings.HasPrefix(contentType, "application/json") {
		return resp, fmt.Errorf("unsupported content-type %q received, only application/json is supported for paged results", contentType)
	}

	// Read the response body and close it
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return resp, fmt.Errorf("could not parse response body")
	}
	resp.Body.Close()

	// Unmarshal firstOdata
	var firstOdata odata.OData
	if err := json.Unmarshal(respBody, &firstOdata); err != nil {
		return resp, err
	}

	firstValue, ok := firstOdata.Value.([]interface{})
	if firstOdata.NextLink == nil || firstValue == nil || !ok {
		// No more pages, reassign response body and return
		resp.Body = io.NopCloser(bytes.NewBuffer(respBody))
		return resp, nil
	}

	// Get the next page, recursively
	// TODO: may have to accommodate APIs with nonstandard paging
	nextReq := req
	u, err := url.Parse(string(*firstOdata.NextLink))
	if err != nil {
		return resp, err
	}
	nextReq.URL = u
	nextResp, err := c.ExecutePaged(ctx, req)
	if err != nil {
		return resp, err
	}

	// Read the next page response body and close it
	nextRespBody, err := io.ReadAll(nextResp.Body)
	if err != nil {
		return resp, fmt.Errorf("could not parse response body")
	}
	nextResp.Body.Close()

	// Unmarshal nextOdata from the next page
	var nextOdata odata.OData
	if err := json.Unmarshal(nextRespBody, &nextOdata); err != nil {
		return nextResp, err
	}

	// When next page has results, append to current page
	if nextValue, ok := nextOdata.Value.([]interface{}); ok {
		value := append(firstValue, nextValue...)
		nextOdata.Value = &value
	}

	// Marshal the entire result, along with fields from the final page
	newJson, err := json.Marshal(nextOdata)
	if err != nil {
		return nextResp, err
	}

	// Reassign the response body
	resp.Body = io.NopCloser(bytes.NewBuffer(newJson))

	return resp, nil
}

// retryableClient instantiates a new *retryablehttp.Client having the provided checkRetry func
func (c *Client) retryableClient(checkRetry retryablehttp.CheckRetry) (r *retryablehttp.Client) {
	r = retryablehttp.NewClient()

	r.Backoff = func(min, max time.Duration, attemptNum int, resp *http.Response) time.Duration {
		if resp != nil {
			// Always look for Retry-After header
			if s, ok := resp.Header["Retry-After"]; ok {
				if sleep, err := strconv.ParseInt(s[0], 10, 64); err == nil {
					return time.Second * time.Duration(sleep)
				}
			}
		}

		// Default exponential backoff
		mult := math.Pow(2, float64(attemptNum)) * float64(min)
		sleep := time.Duration(mult)
		if float64(sleep) != mult || sleep > max {
			sleep = max
		}
		return sleep
	}

	r.CheckRetry = checkRetry
	r.ErrorHandler = RetryableErrorHandler
	r.Logger = nil

	r.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
				d := &net.Dialer{Resolver: &net.Resolver{}}
				return d.DialContext(ctx, network, addr)
			},
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
			ForceAttemptHTTP2:     true,
			MaxIdleConnsPerHost:   runtime.GOMAXPROCS(0) + 1,
		},
	}

	return
}

// containsStatusCode determines whether the returned status code is in the []int of expected status codes.
func containsStatusCode(expected []int, actual int) bool {
	for _, v := range expected {
		if actual == v {
			return true
		}
	}

	return false
}
