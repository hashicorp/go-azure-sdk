package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
	"runtime"
	"strings"
	"time"

	"github.com/hashicorp/go-azure-sdk/auth"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/odata"
	"github.com/hashicorp/go-retryablehttp"
	"github.com/miekg/dns"
)

// RequestRetryFunc is a function that determines whether an HTTP request has failed due to eventual consistency and should be retried
type RequestRetryFunc func(*http.Response, *odata.OData) (bool, error)

// RequestMiddleware can manipulate or log a request before it is sent
type RequestMiddleware func(*http.Request) (*http.Request, error)

// ResponseMiddleware can manipulate or log a response before it is parsed and returned
type ResponseMiddleware func(*http.Request, *http.Response) (*http.Response, error)

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
		return false, nil
	}
}

// ValidStatusFunc is a function that tests whether an HTTP response is considered valid for the particular request.
type ValidStatusFunc func(*http.Response, *odata.OData) bool

// RetryableErrorHandler ensures that the response is returned after exhausting retries for a request
// We can't return an error here, or net/http will not return the response
func RetryableErrorHandler(resp *http.Response, err error, numTries int) (*http.Response, error) {
	return resp, nil
}

type QueryParams map[string]string

func (q QueryParams) Values() url.Values {
	va := make(url.Values)
	for k, v := range q {
		va.Set(k, v)
	}
	return va
}

// Request embeds *http.Request and adds useful metadata
type Request struct {
	OData            odata.Query
	RetryFunc        RequestRetryFunc
	ValidStatusCodes []int
	ValidStatusFunc  ValidStatusFunc

	client *Client

	// Embed *http.Request so that we can send this to an *http.Client
	*http.Request
}

func (r *Request) Marshal(model interface{}) error {
	body, err := json.Marshal(model)
	if err == nil {
		r.ContentLength = int64(len(body))
		r.Body = ioutil.NopCloser(bytes.NewReader(body))
	}
	return nil
}

func (r *Request) Execute() (*Response, *odata.OData, int, error) {
	return r.client.Execute(r)
}

func (r *Request) ExecutePaged() (*Response, *odata.OData, int, error) {
	return r.client.ExecutePaged(r)
}

// Response embeds *http.Response and adds useful methods
type Response struct {
	// Embed *http.Response
	*http.Response
}

func (r *Response) Unmarshal(model interface{}) error {
	if model == nil {
		return fmt.Errorf("model was nil")
	}

	contentType := strings.ToLower(r.Header.Get("Content-Type"))
	if strings.HasPrefix(contentType, "application/json") {
		// Read the response body and close it
		respBody, err := io.ReadAll(r.Body)
		if err != nil {
			return fmt.Errorf("could not parse response body")
		}
		r.Body.Close()

		// Trim away a BOM if present
		respBody = bytes.TrimPrefix(respBody, []byte("\xef\xbb\xbf"))

		// Unmarshal into provided model
		if err := json.Unmarshal(respBody, model); err != nil {
			return err
		}
	}

	return nil
}

// Client is a base client to be used by clients for specific entities.
// It can send GET, POST, PUT, PATCH and DELETE requests to Microsoft Graph and is API version and tenant aware.
type Client struct {
	// Endpoint is the base endpoint for Microsoft Graph, usually "https://graph.microsoft.com".
	Endpoint environments.ApiEndpoint

	// UserAgent is the HTTP user agent string to send in requests.
	UserAgent string

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
func NewClient(endpoint environments.ApiEndpoint) *Client {
	return &Client{
		Endpoint:  endpoint,
		UserAgent: "HashiCorp/go-azure-sdk (Go-http-client/1.1)",
	}
}

// NewRequest configures an *http.Request
func (c *Client) NewRequest(ctx context.Context, method string, contentType string, path string, query QueryParams, odata odata.Query) (*Request, error) {
	req := (&http.Request{}).WithContext(ctx)

	req.Method = method
	req.Header = odata.AppendHeaders(req.Header)
	req.Header.Add("Content-Type", contentType)

	if c.UserAgent != "" {
		req.Header.Add("User-Agent", c.UserAgent)
	}

	if c.Authorizer != nil {
		token, err := c.Authorizer.Token()
		if err != nil {
			return nil, err
		}
		token.SetAuthHeader(req)
	}

	path = strings.TrimPrefix(path, "/")
	u, err := url.ParseRequestURI(fmt.Sprintf("%s/%s", c.Endpoint, path))
	if err != nil {
		return nil, err
	}

	u.RawQuery = query.Values().Encode()

	//host := u.Hostname()
	//port := u.Port()
	//
	//ip, err := lookup(host)
	//if err != nil {
	//	return nil, err
	//}
	//
	//req.Host = u.Host
	//u.Host = *ip
	////u.Host = "10.11.12.251"
	//if len(port) > 0 {
	//	u.Host += ":" + port
	//}

	req.Host = u.Host
	req.URL = u

	return &Request{client: c, Request: req}, nil
}

// Execute is used by the package to send an HTTP request to the API
func (c *Client) Execute(req *Request) (*Response, *odata.OData, int, error) {
	var status int
	var err error

	var reqBody []byte
	if req.Body != nil {
		reqBody, err = io.ReadAll(req.Body)
		if err != nil {
			return nil, nil, status, fmt.Errorf("reading request body: %v", err)
		}
	}
	req.Body = io.NopCloser(bytes.NewBuffer(reqBody))

	// Instantiate a RetryableHttp client and configure its CheckRetry func
	r := c.retryableClient(func(ctx context.Context, r *http.Response, err error) (bool, error) {
		if r != nil && !c.DisableRetries {
			if r.StatusCode == http.StatusFailedDependency {
				return true, nil
			}

			o, err := odata.FromResponse(r)
			if err != nil {
				return false, err
			}

			if f := req.RetryFunc; f != nil {
				return f(r, o)
			}
		}

		return retryablehttp.DefaultRetryPolicy(ctx, r, err)
	})

	// Derive an *http.Client for sending the request
	client := r.StandardClient()

	// Configure any RequestMiddlewares
	if c.RequestMiddlewares != nil {
		for _, m := range *c.RequestMiddlewares {
			r, err := m(req.Request)
			if err != nil {
				return nil, nil, status, err
			}
			req.Request = r
		}
	}

	// Send the request
	resp := &Response{}
	resp.Response, err = client.Do(req.Request)
	if err != nil {
		return resp, nil, status, err
	}
	status = resp.StatusCode

	// Configure any ResponseMiddlewares
	if c.ResponseMiddlewares != nil {
		for _, m := range *c.ResponseMiddlewares {
			r, err := m(req.Request, resp.Response)
			if err != nil {
				return resp, nil, status, err
			}
			resp.Response = r
		}
	}

	// Extract OData from response
	var o *odata.OData
	o, err = odata.FromResponse(resp.Response)
	if err != nil {
		return resp, o, status, err
	}
	if resp == nil {
		return resp, o, status, fmt.Errorf("nil response received")
	}

	// Determine whether response status is valid
	if !containsStatusCode(req.ValidStatusCodes, status) {
		if f := req.ValidStatusFunc; f != nil && f(resp.Response, o) {
			return resp, o, status, nil
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
				return resp, o, status, fmt.Errorf("unexpected status %d, could not read response body", status)
			}
			if len(respBody) == 0 {
				return resp, o, status, fmt.Errorf("unexpected status %d received with no body", status)
			}

			errText = fmt.Sprintf("response: %s", respBody)
		}

		return resp, o, status, fmt.Errorf("unexpected status %d with %s", status, errText)
	}

	return resp, o, status, nil
}

// ExecutePaged automatically pages through the results of Execute
func (c *Client) ExecutePaged(req *Request) (*Response, *odata.OData, int, error) {
	// Perform the request
	resp, o, status, err := c.Execute(req)
	if err != nil {
		return resp, o, status, err
	}

	// Check for json content before handling pagination
	contentType := strings.ToLower(resp.Header.Get("Content-Type"))
	if strings.HasPrefix(contentType, "application/json") {
		// Read the response body and close it
		respBody, err := io.ReadAll(resp.Body)
		if err != nil {
			return resp, o, status, fmt.Errorf("could not parse response body")
		}
		resp.Body.Close()

		// Unmarshal firstOdata
		var firstOdata odata.OData
		if err := json.Unmarshal(respBody, &firstOdata); err != nil {
			return resp, o, status, err
		}

		firstValue, ok := firstOdata.Value.([]interface{})
		if firstOdata.NextLink == nil || firstValue == nil || !ok {
			// No more pages, reassign response body and return
			resp.Body = io.NopCloser(bytes.NewBuffer(respBody))
			return resp, o, status, nil
		}

		// Get the next page, recursively
		nextReq := req
		u, err := url.Parse(*firstOdata.NextLink)
		if err != nil {
			return resp, o, status, err
		}
		nextReq.URL = u
		nextResp, o, status, err := c.ExecutePaged(req)
		if err != nil {
			return resp, o, status, err
		}

		// Read the next page response body and close it
		nextRespBody, err := io.ReadAll(nextResp.Body)
		if err != nil {
			return resp, o, status, fmt.Errorf("could not parse response body")
		}
		nextResp.Body.Close()

		// Unmarshal nextOdata from the next page
		var nextOdata odata.OData
		if err := json.Unmarshal(nextRespBody, &nextOdata); err != nil {
			return resp, o, status, err
		}

		if nextValue, ok := nextOdata.Value.([]interface{}); ok {
			// Next page has results, append to current page
			value := append(firstValue, nextValue...)
			nextOdata.Value = &value
		}

		// Marshal the entire result, along with fields from the final page
		newJson, err := json.Marshal(nextOdata)
		if err != nil {
			return resp, o, status, err
		}

		// Reassign the response body
		resp.Body = io.NopCloser(bytes.NewBuffer(newJson))
	}

	return resp, o, status, nil
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

func lookup(host string) (*string, error) {
	m := &dns.Msg{}
	m.SetQuestion(dns.Fqdn(host), dns.TypeA)
	d := &dns.Client{
		Dialer: &net.Dialer{
			Timeout: 2000 * time.Millisecond,
		},
		SingleInflight: true,
	}

	in, _, err := d.Exchange(m, "1.1.1.1:53")
	if err != nil {
		return nil, err
	}

	var ips, cnames []string

	for _, rr := range in.Answer {
		if v, ok := rr.(*dns.A); ok {
			ips = append(ips, v.A.String())
		} else if v, ok := rr.(*dns.CNAME); ok {
			cnames = append(cnames, v.Target)
		}
	}

	if len(ips) > 0 {
		return &ips[0], nil
	}

	for _, cname := range cnames {
		if ip, err := lookup(cname); err != nil {
			return ip, nil
		}
	}

	return nil, fmt.Errorf("host not found: %s", host)
}

// retryableClient instantiates a new *retryablehttp.Client having the provided checkRetry func
func (c *Client) retryableClient(checkRetry retryablehttp.CheckRetry) (r *retryablehttp.Client) {
	r = retryablehttp.NewClient()

	r.CheckRetry = checkRetry
	r.ErrorHandler = RetryableErrorHandler
	r.Logger = nil

	proxy := http.ProxyFromEnvironment

	//customDns := false
	//if u, e := proxy(&http.Request{}); u == nil && e == nil {
	//	customDns = true
	//}

	r.HTTPClient = &http.Client{
		Transport: &http.Transport{
			Proxy: proxy,
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
