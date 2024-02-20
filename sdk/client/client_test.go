// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package client

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"io"
	"log"
	"net/http"
	"net/url"
	"reflect"
	"testing"

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

func unmarshalResponse(body io.ReadCloser, unmarshal func(in []byte) error) error {
	respBody, err := io.ReadAll(body)
	if err != nil {
		return fmt.Errorf("parsing response body: %+v", err)
	}
	body.Close()

	return unmarshal(respBody)
}
