// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
	"github.com/hashicorp/go-retryablehttp"
)

var _ BaseClient = &testClient{}

type testClient struct {
	*Client
}

func (c *testClient) NewRequest(ctx context.Context, input RequestOptions) (*Request, error) {
	req, err := c.Client.NewRequest(ctx, input)
	if err != nil {
		return nil, fmt.Errorf("building %s request: %+v", input.HttpMethod, err)
	}

	req.Client = c
	query := url.Values{}

	if input.OptionsObject != nil {
		if h := input.OptionsObject.ToHeaders(); h != nil {
			for k, v := range h.Headers() {
				req.Header[k] = v
			}
		}

		if q := input.OptionsObject.ToQuery(); q != nil {
			for k, v := range q.Values() {
				// we intentionally only add one of each type
				query.Del(k)
				query.Add(k, v[0])
			}
		}

		if o := input.OptionsObject.ToOData(); o != nil {
			req.Header = o.AppendHeaders(req.Header)
			query = o.AppendValues(query)
		}
	}

	req.URL.RawQuery = query.Encode()
	req.ValidStatusCodes = input.ExpectedStatusCodes

	return req, nil
}

func TestAccClient(t *testing.T) {
	test.AccTest(t)

	ctx := context.TODO()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.MicrosoftGraph
	endpoint, ok := api.Endpoint()
	if !ok {
		t.Fatalf("missing endpoint for microsoft graph for this environment")
	}
	conn.Authorize(ctx, t, api)

	c := &testClient{
		Client: NewClient(*endpoint, "example", "2020-01-01"),
	}
	c.SetAuthorizer(conn.Authorizer)

	path := fmt.Sprintf("/v1.0/servicePrincipals/%s", conn.Claims.ObjectId)
	reqOpts := RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod:    http.MethodGet,
		OptionsObject: nil,
		Path:          path,
	}
	req, err := c.NewRequest(ctx, reqOpts)
	if err != nil {
		t.Fatal(err)
	}

	resp, err := req.Execute(ctx)
	if err != nil {
		t.Fatalf("Execute(): %v", err)
	}

	fmt.Printf("%#v", resp)
}

var _ Options = &requestOptions{}

type requestOptions struct {
	query *odata.Query
}

func (r *requestOptions) ToHeaders() *Headers   { return nil }
func (r *requestOptions) ToOData() *odata.Query { return r.query }
func (r *requestOptions) ToQuery() *QueryParams { return nil }

func TestAccClient_Paged(t *testing.T) {
	test.AccTest(t)

	ctx := context.TODO()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.MicrosoftGraph
	endpoint, ok := api.Endpoint()
	if !ok {
		t.Fatalf("missing endpoint for microsoft graph for this environment")
	}
	conn.Authorize(ctx, t, api)

	c := &testClient{
		Client: NewClient(*endpoint, "example", "2020-01-01"),
	}
	c.SetAuthorizer(conn.Authorizer)

	path := "/v1.0/applications"
	reqOpts := RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		OptionsObject: &requestOptions{
			query: &odata.Query{
				Filter: "startsWith(displayName,'acctest')",
				Select: []string{"appId", "displayName"},
				Top:    10,
			},
		},
		Path: path,
	}
	req, err := c.NewRequest(ctx, reqOpts)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = req.ExecutePaged(ctx); err != nil {
		t.Fatalf("ExecutePaged(): %v", err)
	}
}

var _ odata.CustomPager = &pager{}

type pager struct {
	NextLink *odata.Link `json:"@odata.nextLink"`
}

func (p *pager) NextPageLink() *odata.Link {
	if p == nil {
		log.Fatalf("pager: p was nil")
	}
	if p.NextLink == nil {
		log.Printf("[DEBUG] pager: nextLink was nil")
	} else {
		log.Printf("[DEBUG] pager: found custom nextLink %q", *p.NextLink)
	}
	defer func() {
		p.NextLink = nil
	}()
	return p.NextLink
}

func TestAccClient_CustomPaged(t *testing.T) {
	test.AccTest(t)

	ctx := context.TODO()
	conn := test.NewConnection(t)
	api := conn.AuthConfig.Environment.MicrosoftGraph
	endpoint, ok := api.Endpoint()
	if !ok {
		t.Fatalf("missing endpoint for microsoft graph for this environment")
	}
	conn.Authorize(ctx, t, api)

	c := &testClient{
		Client: NewClient(*endpoint, "example", "2020-01-01"),
	}
	c.SetAuthorizer(conn.Authorizer)

	path := "/v1.0/applications"
	reqOpts := RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		OptionsObject: &requestOptions{
			query: &odata.Query{
				Filter: "startsWith(displayName,'acctest')",
				Select: []string{"appId", "displayName"},
				Top:    10,
			},
		},
		Pager: &pager{},
		Path:  path,
	}
	req, err := c.NewRequest(ctx, reqOpts)
	if err != nil {
		t.Fatal(err)
	}

	if _, err = req.ExecutePaged(ctx); err != nil {
		t.Fatalf("ExecutePaged(): %v", err)
	}
}

func TestMarshalByteStreamAndPowerShell(t *testing.T) {
	contentTypes := []string{
		"application/octet-stream",
		"text/powershell",
	}
	value := []byte("What is my purpose?")
	for _, contentType := range contentTypes {
		r := &Request{
			Request: &http.Request{
				Header: map[string][]string{
					"Content-Type": {contentType},
				},
			},
		}
		if err := r.Marshal(&value); err != nil {
			t.Fatalf("marshaling: %+v", err)
		}

		err := unmarshalResponse(r.Body, func(in []byte) error {
			out := string(in)
			if out != "What is my purpose?" {
				return fmt.Errorf("expected the marshalled response to match `What is my purpose?` but got %q", out)
			}
			return nil
		})
		if err != nil {
			t.Fatalf("validating marshaled value: %+v", err)
		}
	}
}

func TestMarshalJson(t *testing.T) {
	type sampleObj struct {
		Name   string `json:"name"`
		Animal string `json:"animal"`
	}
	type payload struct {
		Inner sampleObj `json:"inner"`
	}

	val := payload{
		Inner: sampleObj{
			Name:   "tabatha",
			Animal: "cat",
		},
	}
	r := &Request{
		Request: &http.Request{
			Header: map[string][]string{
				"Content-Type": {"application/json"},
			},
		},
	}
	if err := r.Marshal(&val); err != nil {
		t.Fatalf("marshaling: %+v", err)
	}

	var unmarshaled payload
	err := unmarshalResponse(r.Body, func(in []byte) error {
		if e := json.Unmarshal(in, &unmarshaled); e != nil {
			return e
		}
		return nil
	})
	if err != nil {
		t.Fatalf("unmarshaling value: %+v", err)
	}
	if !reflect.DeepEqual(val, unmarshaled) {
		t.Fatalf("unexpected difference in encoded objects. First: %+v\n\nSecond: %+v", val, unmarshaled)
	}
}

func TestMarshalXml(t *testing.T) {
	type sampleObj struct {
		Name   string `xml:"name"`
		Animal string `xml:"animal"`
	}
	type payload struct {
		Inner sampleObj `xml:"inner"`
	}

	contentTypes := []string{
		"application/xml",
		"text/xml",
	}
	for _, contentType := range contentTypes {
		val := payload{
			Inner: sampleObj{
				Name:   "tabatha",
				Animal: "cat",
			},
		}
		r := &Request{
			Request: &http.Request{
				Header: map[string][]string{
					"Content-Type": {contentType},
				},
			},
		}
		if err := r.Marshal(&val); err != nil {
			t.Fatalf("marshaling: %+v", err)
		}

		var unmarshaled payload
		err := unmarshalResponse(r.Body, func(in []byte) error {
			if e := xml.Unmarshal(in, &unmarshaled); e != nil {
				return e
			}
			return nil
		})
		if err != nil {
			t.Fatalf("unmarshaling value: %+v", err)
		}
		if !reflect.DeepEqual(val, unmarshaled) {
			t.Fatalf("unexpected difference in encoded objects. First: %+v\n\nSecond: %+v", val, unmarshaled)
		}
	}
}

func TestUnmarshalByteStreamAndPowerShell(t *testing.T) {
	contentTypes := []string{
		"application/octet-stream",
		"text/powershell",
	}
	expected := []byte("you serve butter")
	for _, contentType := range contentTypes {
		r := &Response{
			Response: &http.Response{
				Header: map[string][]string{
					"Content-Type": {contentType},
				},
				Body: io.NopCloser(bytes.NewReader(expected)),
			},
		}
		var unmarshaled = make([]byte, 0)
		if err := r.Unmarshal(&unmarshaled); err != nil {
			t.Fatalf("unmarshaling: %+v", err)
		}
		if string(unmarshaled) != "you serve butter" {
			t.Fatalf("unexpected difference in decoded objects. Expected %q\n\nGot: %q", string(expected), string(unmarshaled))
		}
	}
}

func TestUnmarshalByteStreamAndPowerShellWithModel(t *testing.T) {
	contentTypes := []string{
		"application/octet-stream",
		"text/powershell",
	}
	var respModel = struct {
		HttpResponse *http.Response
		Model        *[]byte
	}{
		Model: pointer.To(make([]byte, 0)),
	}
	expected := []byte("you serve butter")
	for _, contentType := range contentTypes {
		r := &Response{
			Response: &http.Response{
				Header: map[string][]string{
					"Content-Type": {contentType},
				},
				Body: io.NopCloser(bytes.NewReader(expected)),
			},
		}
		if err := r.Unmarshal(respModel.Model); err != nil {
			t.Fatalf("unmarshaling: %+v", err)
		}
		if string(*respModel.Model) != "you serve butter" {
			t.Fatalf("unexpected difference in decoded objects. Expected %q\n\nGot: %q", string(expected), string(*respModel.Model))
		}
	}
}

func TestUnmarshalJson(t *testing.T) {
	type sampleObj struct {
		Name   string `json:"name"`
		Animal string `json:"animal"`
	}
	type payload struct {
		Inner sampleObj `json:"inner"`
	}
	expected := payload{
		Inner: sampleObj{
			Name:   "tabatha",
			Animal: "cat",
		},
	}
	input, err := json.Marshal(expected)
	if err != nil {
		t.Fatalf("preparing source data for unmarshaling: %+v", err)
	}
	r := &Response{
		Response: &http.Response{
			Header: map[string][]string{
				"Content-Type": {"application/json"},
			},
			Body: io.NopCloser(bytes.NewReader(input)),
		},
	}
	var unmarshaled payload
	if err := r.Unmarshal(&unmarshaled); err != nil {
		t.Fatalf("unmarshaling value: %+v", err)
	}
	if !reflect.DeepEqual(expected, unmarshaled) {
		t.Fatalf("unexpected difference in decoded objects. Expected %q\n\nGot: %+v", expected, unmarshaled)
	}
}

func TestUnmarshalXml(t *testing.T) {
	type sampleObj struct {
		Name   string `xml:"name"`
		Animal string `xml:"animal"`
	}
	type payload struct {
		Inner sampleObj `xml:"inner"`
	}
	contentTypes := []string{
		"application/xml",
		"text/xml",
	}
	expected := payload{
		Inner: sampleObj{
			Name:   "tabatha",
			Animal: "cat",
		},
	}
	for _, contentType := range contentTypes {
		input, err := xml.Marshal(expected)
		if err != nil {
			t.Fatalf("preparing source data for unmarshaling: %+v", err)
		}
		r := &Response{
			Response: &http.Response{
				Header: map[string][]string{
					"Content-Type": {contentType},
				},
				Body: io.NopCloser(bytes.NewReader(input)),
			},
		}
		var unmarshaled payload
		if err := r.Unmarshal(&unmarshaled); err != nil {
			t.Fatalf("unmarshaling value: %+v", err)
		}
		if !reflect.DeepEqual(expected, unmarshaled) {
			t.Fatalf("unexpected difference in decoded objects. Expected %q\n\nGot: %+v", expected, unmarshaled)
		}
	}
}

func TestUnmarshalNilHeaders(t *testing.T) {
	expected := []byte("any payload")
	r := &Response{
		Response: &http.Response{
			Header: nil,
			Body:   io.NopCloser(bytes.NewReader(expected)),
		},
	}
	var unmarshaled []byte
	if err := r.Unmarshal(&unmarshaled); err != nil {
		if err.Error() != "could not determine Content-Type for response" {
			t.Fatalf("unexpected error when unmarshaling: %+v", err)
		}
	} else {
		t.Fatalf("expected an error but got no error")
	}
}

func TestUnmarshalNilResponse(t *testing.T) {
	r := &Response{
		Response: nil,
	}
	var unmarshaled = make([]byte, 0)
	if err := r.Unmarshal(&unmarshaled); err != nil {
		if err.Error() != "could not unmarshal as the HTTP response was nil" {
			t.Fatalf("unexpected error when unmarshaling: %+v", err)
		}
	} else {
		t.Fatalf("expected an error but got no error")
	}
}

func unmarshalResponse(body io.ReadCloser, unmarshal func(in []byte) error) error {
	respBody, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("parsing response body: %+v", err)
	}
	body.Close()

	return unmarshal(respBody)
}

type roundTripperMock struct {
	roundTripFunc func(*http.Request) (*http.Response, error)
}

func (rt *roundTripperMock) RoundTrip(req *http.Request) (*http.Response, error) {
	if rt.roundTripFunc != nil {
		return rt.roundTripFunc(req)
	}
	return nil, fmt.Errorf("roundTripFunc missing from mock")
}

func TestClient_CustomTransport(t *testing.T) {
	ctx := context.TODO()

	c := NewClient("http://localhost", "testService", "v1.0")
	c.DisableRetries = true

	hitCount := 0
	mockTransport := &roundTripperMock{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			hitCount++
			return &http.Response{
				StatusCode: http.StatusOK,
				Body:       io.NopCloser(bytes.NewReader([]byte(`{"success": true}`))),
				Header: map[string][]string{
					"Content-Type": {"application/json"},
				},
				Request: req,
			}, nil
		},
	}
	c.Transport = mockTransport

	reqOpts := RequestOptions{
		ContentType: "application/json",
		ExpectedStatusCodes: []int{
			http.StatusOK,
		},
		HttpMethod: http.MethodGet,
		Path:       "/test",
	}

	req, err := c.NewRequest(ctx, reqOpts)
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	resp, err := req.Execute(ctx)
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", resp.StatusCode)
	}

	if hitCount != 1 {
		t.Errorf("expected transport to be hit 1 time, got %d", hitCount)
	}
}

// TestClient_RetryAfterNotLimitedByRetryCount ensures a short Retry-After is not starved of
// retries by a count precomputed from the deadline assuming the slower exponential schedule.
func TestClient_RetryAfterNotLimitedByRetryCount(t *testing.T) {
	var requestCount atomic.Int32

	const throttledResponses = 5

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		if requestCount.Add(1) <= throttledResponses {
			w.Header().Set("Retry-After", "1")
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"success": true}`))
	}))
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	resp, err := req.Execute(ctx)
	if err != nil {
		t.Fatalf("Execute error: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK, got %d", resp.StatusCode)
	}

	// Well beyond the 2 retries the old precomputed formula allowed for a 10s deadline.
	if got := requestCount.Load(); got != throttledResponses+1 {
		t.Errorf("expected exactly %d requests, got %d", throttledResponses+1, got)
	}
}

// TestClient_RetryStopsBeforeDeadline ensures retrying gives up before the deadline, so the
// caller receives the real response rather than context.DeadlineExceeded.
func TestClient_RetryStopsBeforeDeadline(t *testing.T) {
	var requestCount atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestCount.Add(1)
		w.Header().Set("Retry-After", "2")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	start := time.Now()
	resp, err := req.Execute(ctx)
	elapsed := time.Since(start)

	if elapsed >= 3*time.Second {
		t.Errorf("expected Execute to return before the 3s deadline, took %s", elapsed)
	}

	if err == nil {
		t.Fatalf("expected an error, got nil")
	}
	if errors.Is(err, context.DeadlineExceeded) {
		t.Errorf("expected error not to be context.DeadlineExceeded, got: %v", err)
	}

	if resp == nil || resp.Response == nil {
		t.Fatalf("expected a non-nil response")
	}
	if resp.StatusCode != http.StatusTooManyRequests {
		t.Errorf("expected 429, got %d", resp.StatusCode)
	}

	if got := requestCount.Load(); got < 1 {
		t.Errorf("expected at least 1 request, got %d", got)
	}
}

// TestClient_RetryMaxWithoutDeadline ensures the fixed RetryMax of 16 still applies when the
// context has no deadline.
func TestClient_RetryMaxWithoutDeadline(t *testing.T) {
	var requestCount atomic.Int32

	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		requestCount.Add(1)
		w.Header().Set("Retry-After", "1")
		w.WriteHeader(http.StatusTooManyRequests)
	}))
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")

	ctx := context.Background()

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/test",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	_, _ = req.Execute(ctx)

	if got := requestCount.Load(); got != 17 {
		t.Errorf("expected exactly 17 requests (1 + RetryMax 16), got %d", got)
	}
}

// TestClient_RetryAfterHotLoopRegression ensures a non-positive or overflowing Retry-After
// cannot produce a zero wait. Such values were passed through as the sleep duration, hot-looping
// the retries until the deadline (~45,000 requests in 2s).
func TestClient_RetryAfterHotLoopRegression(t *testing.T) {
	testCases := []struct {
		name       string
		retryAfter string
	}{
		{name: "zero", retryAfter: "0"},
		{name: "negative", retryAfter: "-1"},
		{name: "overflowing", retryAfter: "9223372036854775807"},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var requestCount atomic.Int32

			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
				requestCount.Add(1)
				w.Header().Set("Retry-After", tc.retryAfter)
				w.WriteHeader(http.StatusTooManyRequests)
			}))
			defer server.Close()

			c := NewClient(server.URL, "testService", "v1.0")

			ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
			defer cancel()

			req, err := c.NewRequest(ctx, RequestOptions{
				ContentType:         "application/json",
				ExpectedStatusCodes: []int{http.StatusOK},
				HttpMethod:          http.MethodGet,
				Path:                "/test",
			})
			if err != nil {
				t.Fatalf("NewRequest error: %v", err)
			}

			start := time.Now()
			_, _ = req.Execute(ctx)
			elapsed := time.Since(start)

			// Tens of thousands before the fix; RetryWaitMin bounds this to a handful.
			if got := requestCount.Load(); got >= 100 {
				t.Errorf("expected a small bounded request count under a 2s deadline, got %d requests in %s (hot loop?)", got, elapsed)
			}
		})
	}
}

// TestClient_RetryStateResetAcrossRepeatedDo ensures the attempt counter and cached backoff are
// reset per Do call, not per retryableClient call: carrying an attempt count of 2 into the second
// Do would compute a 4s wait, cross the 5s deadline and give up on a retry that should succeed.
// It drives Do directly, as this package's redirect handling does not re-enter Do today.
func TestClient_RetryStateResetAcrossRepeatedDo(t *testing.T) {
	var xRequests, yRequests atomic.Int32

	mux := http.NewServeMux()
	mux.HandleFunc("/x", func(w http.ResponseWriter, _ *http.Request) {
		if xRequests.Add(1) <= 2 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	})
	mux.HandleFunc("/y", func(w http.ResponseWriter, _ *http.Request) {
		if yRequests.Add(1) <= 1 {
			w.WriteHeader(http.StatusTooManyRequests)
			return
		}
		w.WriteHeader(http.StatusOK)
	})

	server := httptest.NewServer(mux)
	defer server.Close()

	c := NewClient(server.URL, "testService", "v1.0")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	r := c.retryableClient(ctx, retryablehttp.DefaultRetryPolicy)

	reqX, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/x", nil)
	if err != nil {
		t.Fatalf("building request for /x: %v", err)
	}
	respX, err := r.Do(reqX)
	if err != nil {
		t.Fatalf("Do(/x) error: %v", err)
	}
	if respX.StatusCode != http.StatusOK {
		t.Fatalf("expected /x to eventually succeed, got status %d", respX.StatusCode)
	}

	reqY, err := retryablehttp.NewRequestWithContext(ctx, http.MethodGet, server.URL+"/y", nil)
	if err != nil {
		t.Fatalf("building request for /y: %v", err)
	}
	respY, err := r.Do(reqY)
	if err != nil {
		t.Fatalf("Do(/y) error: %v", err)
	}

	if respY.StatusCode != http.StatusOK {
		t.Errorf("expected 200 OK for /y (retry state reset across the second Do() call), got %d", respY.StatusCode)
	}
}
