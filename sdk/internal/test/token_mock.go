// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package test

import (
	"bytes"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"io"
	"net/http"
	"net/url"
	"regexp"
)

const (
	uuidRegex = "[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}"
)

type AzureADAccessTokenMockClient struct {
	Authorization environments.Authorization
}

func (c *AzureADAccessTokenMockClient) Do(r *http.Request) (resp *http.Response, err error) {
	if r == nil {
		return nil, fmt.Errorf("request was nil")
	}

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

	if err = resp.Request.ParseForm(); err != nil {
		return nil, err
	}

	u, err := url.Parse(c.Authorization.LoginEndpoint)
	if err != nil {
		return nil, err
	}

	loginHost := u.Host
	if p := u.Port(); p != "" {
		loginHost = fmt.Sprintf("%s:%s", loginHost, p)
	}

	metadataHost := "169.254.169.254"

	switch r.Host {
	case loginHost:
		return c.tokenFromLoginEndpoint(resp)
	case metadataHost:
		return c.tokenFromInstanceMetadataService(resp)
	}

	return nil, fmt.Errorf("unexpected Host header, expected %q, received %q", u.Host, loginHost)

}

func (c *AzureADAccessTokenMockClient) tokenFromLoginEndpoint(resp *http.Response) (*http.Response, error) {
	if resp.Request.Method != http.MethodPost {
		return nil, fmt.Errorf("unexpected request method, should be POST, received %s", resp.Request.Method)
	}

	pathRegex := fmt.Sprintf(`^\/(%s)(\/oauth2\/v2.0\/token)$`, uuidRegex)
	if m := regexp.MustCompile(pathRegex).FindStringSubmatch(resp.Request.URL.Path); len(m) != 3 {
		return nil, fmt.Errorf("URL path did not match pattern %q: %q", pathRegex, resp.Request.URL.Path)
	}

	clientId := resp.Request.PostForm.Get("client_id")
	if clientId == "" {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing 'client_id' field"}`))
		return resp, nil
	}

	scope := resp.Request.PostForm.Get("scope")
	if scope == "" {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing 'scope' field"}`))
		return resp, nil
	}

	grantType := resp.Request.PostForm.Get("grant_type")
	switch grantType {
	case "client_credentials":
		if clientSecret := resp.Request.PostForm.Get("client_secret"); clientSecret != "" {
			return c.token(resp)
		}

		if clientAssertion := resp.Request.PostForm.Get("client_assertion"); clientAssertion != "" {
			switch clientAssertionType := resp.Request.PostForm.Get("client_assertion_type"); clientAssertionType {
			case "urn:ietf:params:oauth:client-assertion-type:jwt-bearer":
				return c.token(resp)
			}

			resp.StatusCode = http.StatusBadRequest
			resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"unrecognised value for 'client_assertion_type' field"}`))

			return resp, nil
		}
	default:
		return nil, fmt.Errorf("grant type %q is not supported", grantType)
	}

	resp.StatusCode = http.StatusBadRequest
	resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"unsupported or unrecognised authentication parameters"}`))

	return resp, nil
}

func (c *AzureADAccessTokenMockClient) tokenFromInstanceMetadataService(resp *http.Response) (*http.Response, error) {
	if resp.Request.Method != http.MethodGet {
		return nil, fmt.Errorf("unexpected request method, should be GET, received %s", resp.Request.Method)
	}

	if expected := "/metadata/identity/oauth2/token"; resp.Request.URL.Path != expected {
		return nil, fmt.Errorf("unexpected URL path, expected %q, received %q", expected, resp.Request.URL.Path)
	}

	if resp.Request.Header.Get("Metadata") != "true" {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing or invalid 'Metadata' header"}`))
		return resp, nil
	}

	apiVersion := resp.Request.Form.Get("api-version")
	if apiVersion == "" {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing 'api-version' field"}`))
		return resp, nil
	}

	clientId := resp.Request.Form.Get("client_id")
	if clientId == "" {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing 'client_id' field"}`))
		return resp, nil
	}

	resource := resp.Request.Form.Get("resource")
	if resource == "" {
		resp.StatusCode = http.StatusBadRequest
		resp.Body = io.NopCloser(bytes.NewBufferString(`{"error":"missing 'resource' field"}`))
		return resp, nil
	}

	body := fmt.Sprintf(`{"access_token":"%s","client_id":"%s","expires_in":"86391","expires_on":"1611701390","ext_expires_in":"86399","not_before":"1611614690","resource":"%s","token_type":"Bearer"}`,
		DummyAccessToken, clientId, resource)

	resp.Body = io.NopCloser(bytes.NewBufferString(body))

	return resp, nil
}

func (c *AzureADAccessTokenMockClient) token(resp *http.Response) (*http.Response, error) {
	resp.Body = io.NopCloser(bytes.NewBufferString(fmt.Sprintf(`{"token_type":"Bearer","expires_in":3599,"ext_expires_in":3599,"access_token":"%s"}`, DummyAccessToken)))
	return resp, nil
}
