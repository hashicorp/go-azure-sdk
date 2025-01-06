// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

type OidcMockClientPlatform int

const (
	OidcMockClientPlatformGithub OidcMockClientPlatform = iota
	OidcMockClientPlatformADOPipeline
)

type oidcMockClient struct {
	authorization environments.Authorization
	platform      OidcMockClientPlatform
	mockHost      string
}

func (c *oidcMockClient) Do(r *http.Request) (resp *http.Response, err error) {
	if r == nil {
		return nil, fmt.Errorf("request was nil")
	}

	if r.Host == c.mockHost {
		resp = &http.Response{
			Proto:      "HTTP/1.1",
			ProtoMajor: 1,
			ProtoMinor: 1,
			StatusCode: http.StatusOK,
			Header: http.Header{
				"Content-Type": []string{"application/json; charset=utf-8"},
			},
			Request: r,
		}

		auth := strings.Split(r.Header.Get("Authorization"), " ")
		if len(auth) != 2 {
			resp.StatusCode = http.StatusUnauthorized
			resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing or malformed Authorization header"}`))
			return
		}

		if !strings.EqualFold(auth[0], "Bearer") {
			resp.StatusCode = http.StatusUnauthorized
			resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"unsupported Authorization header"}`))
			return
		}

		if auth[1] != test.DummyAccessToken {
			resp.StatusCode = http.StatusUnauthorized
			resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"request access token is invalid"}`))
			return
		}

		q := r.URL.Query()
		if q.Get("audience") != "api://AzureADTokenExchange" {
			resp.StatusCode = http.StatusBadRequest
			resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"invalid audience"}`))
			return
		}

		var body string
		switch c.platform {
		case OidcMockClientPlatformGithub:
			body = fmt.Sprintf(`{"count":1,"value":"%s"}`, test.DummyIDToken)
		case OidcMockClientPlatformADOPipeline:
			body = fmt.Sprintf(`{"oidcToken":"%s"}`, test.DummyIDToken)
		default:
			return nil, fmt.Errorf("unsupported platform: %v", c.platform)
		}

		resp.Body = io.NopCloser(bytes.NewBufferString(body))

		return
	}
	client := &test.AzureADAccessTokenMockClient{
		Authorization: c.authorization,
	}
	return client.Do(r)
}
