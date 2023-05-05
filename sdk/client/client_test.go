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
	"net/http"
	"reflect"
	"testing"

	"github.com/hashicorp/go-azure-helpers/lang/pointer"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

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

	c := NewClient(*endpoint, "example", "2020-01-01")
	c.Authorizer = conn.Authorizer

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

	_, err = req.ExecutePaged(ctx)
	if err != nil {
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
		var unmarshaled = pointer.To(make([]byte, 0))
		if err := r.Unmarshal(&unmarshaled); err != nil {
			t.Fatalf("unmarshaling: %+v", err)
		}
		if string(*unmarshaled) != "you serve butter" {
			t.Fatalf("unexpected difference in decoded objects. Expected %q\n\nGot: %q", string(expected), string(*unmarshaled))
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
	}{}
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
		if err := r.Unmarshal(&respModel.Model); err != nil {
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
