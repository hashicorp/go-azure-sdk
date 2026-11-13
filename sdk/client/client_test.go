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
	"log"
	"net"
	"net/http"
	"net/http/httptest"
	"net/http/httptrace"
	"net/url"
	"reflect"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
	"github.com/hashicorp/go-azure-sdk/sdk/odata"
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
	_ = body.Close()

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

func TestClient_DefaultTransportReusesConnections(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	send := func(c *Client) bool {
		var reused bool
		trace := &httptrace.ClientTrace{
			GotConn: func(info httptrace.GotConnInfo) {
				reused = info.Reused
			},
		}
		ctx := httptrace.WithClientTrace(context.Background(), trace)
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
		var out map[string]bool
		if err := resp.Unmarshal(&out); err != nil {
			t.Fatalf("Unmarshal error: %v", err)
		}
		return reused
	}

	c1 := NewClient(srv.URL, "testService", "v1.0")
	c1.DisableRetries = true
	if send(c1) {
		t.Fatal("expected first connection not to be reused")
	}

	// A second distinct Client with Transport == nil should reuse the pooled connection
	c2 := NewClient(srv.URL, "testService", "v1.0")
	c2.DisableRetries = true
	if !send(c2) {
		t.Fatal("expected second request across clients to reuse pooled TCP connection")
	}
}

func TestClient_SetTransport(t *testing.T) {
	c := NewClient("http://localhost", "testService", "v1.0")
	custom := &roundTripperMock{}
	c.SetTransport(custom)
	if c.Transport != custom {
		t.Fatal("expected SetTransport to configure c.Transport")
	}
}

type trackingReadCloser struct {
	io.Reader
	closed bool
}

func (t *trackingReadCloser) Close() error {
	t.closed = true
	return nil
}

func TestExecutePaged_ClosesBodyOnUnsupportedContentType(t *testing.T) {
	ctx := context.Background()
	body := &trackingReadCloser{Reader: bytes.NewBufferString("plain text")}

	c := NewClient("http://localhost", "testService", "v1.0")
	c.DisableRetries = true
	c.SetTransport(&roundTripperMock{
		roundTripFunc: func(req *http.Request) (*http.Response, error) {
			return &http.Response{
				StatusCode: http.StatusOK,
				Header:     http.Header{"Content-Type": []string{"text/plain"}},
				Body:       body,
				Request:    req,
			}, nil
		},
	})

	req, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodGet,
		Path:                "/paged",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}

	if _, err := c.ExecutePaged(ctx, req); err == nil {
		t.Fatal("expected error for unsupported content-type")
	}
	if !body.closed {
		t.Fatal("expected response body to be closed on unsupported content-type")
	}
}

func TestResponseUnmarshal_ClosesZeroLengthBody(t *testing.T) {
	body := &trackingReadCloser{Reader: bytes.NewReader(nil)}
	resp := &Response{
		Response: &http.Response{
			StatusCode:    http.StatusOK,
			ContentLength: 0,
			Header:        http.Header{"Content-Type": []string{"application/json"}},
			Body:          body,
			Request:       &http.Request{Header: http.Header{}},
		},
	}

	var out map[string]interface{}
	if err := resp.Unmarshal(&out); err != nil {
		t.Fatalf("Unmarshal error: %v", err)
	}
	if !body.closed {
		t.Fatal("expected zero-length response body to be closed")
	}
}

func TestCloseIdleConnections(t *testing.T) {
	srv := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	defer srv.Close()

	send := func() bool {
		var reused bool
		ctx := httptrace.WithClientTrace(context.Background(), &httptrace.ClientTrace{
			GotConn: func(info httptrace.GotConnInfo) {
				reused = info.Reused
			},
		})
		c := NewClient(srv.URL, "testService", "v1.0")
		c.DisableRetries = true
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
		var out map[string]bool
		if err := resp.Unmarshal(&out); err != nil {
			t.Fatalf("Unmarshal error: %v", err)
		}
		return reused
	}

	_ = send()
	CloseIdleConnections()
	if send() {
		t.Fatal("expected connection not to be reused after CloseIdleConnections()")
	}
}

func TestClient_ConcurrentDefaultTransport(t *testing.T) {
	var newConns atomic.Int64
	srv := httptest.NewUnstartedServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte(`{"ok":true}`))
	}))
	srv.Config.ConnState = func(_ net.Conn, state http.ConnState) {
		if state == http.StateNew {
			newConns.Add(1)
		}
	}
	srv.Start()
	defer srv.Close()

	const workers = 32
	const requestsPerWorker = 10

	stopCloser := make(chan struct{})
	var closerWg sync.WaitGroup
	closerWg.Add(1)
	go func() {
		defer closerWg.Done()
		for {
			select {
			case <-stopCloser:
				return
			default:
				CloseIdleConnections()
				time.Sleep(time.Millisecond)
			}
		}
	}()

	var wg sync.WaitGroup
	errs := make(chan error, workers*requestsPerWorker)
	for i := 0; i < workers; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			c := NewClient(srv.URL, "testService", "v1.0")
			c.DisableRetries = true
			for j := 0; j < requestsPerWorker; j++ {
				req, err := c.NewRequest(context.Background(), RequestOptions{
					ContentType:         "application/json",
					ExpectedStatusCodes: []int{http.StatusOK},
					HttpMethod:          http.MethodGet,
					Path:                "/test",
				})
				if err != nil {
					errs <- err
					return
				}
				resp, err := req.Execute(context.Background())
				if err != nil {
					errs <- err
					return
				}
				var out map[string]bool
				if err := resp.Unmarshal(&out); err != nil {
					errs <- err
					return
				}
			}
		}()
	}
	wg.Wait()
	close(stopCloser)
	closerWg.Wait()
	close(errs)

	for err := range errs {
		if err != nil {
			t.Fatalf("concurrent Execute failed: %v", err)
		}
	}

	totalRequests := int64(workers * requestsPerWorker)
	if opened := newConns.Load(); opened >= totalRequests {
		t.Fatalf("expected connection pooling to open fewer than %d connections, got %d", totalRequests, opened)
	}
}

func TestRequest_IdempotentOverride(t *testing.T) {
	ctx := context.Background()
	c := NewClient("http://localhost", "testService", "v1.0")

	putReq, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodPut,
		Path:                "/resource",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}
	if putReq.IsIdempotent() {
		t.Fatal("expected default PUT request not to be idempotent")
	}

	idempotentPutReq, err := c.NewRequest(ctx, RequestOptions{
		ContentType:         "application/json",
		ExpectedStatusCodes: []int{http.StatusOK},
		HttpMethod:          http.MethodPut,
		Idempotent:          pointer.To(true),
		Path:                "/resource",
	})
	if err != nil {
		t.Fatalf("NewRequest error: %v", err)
	}
	if !idempotentPutReq.IsIdempotent() {
		t.Fatal("expected PUT request with Idempotent=true to be idempotent")
	}
}
