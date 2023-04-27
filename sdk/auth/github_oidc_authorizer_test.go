// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"net/http"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestGitHubOIDCAuthorizer(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	auth.Client = &oidcMockClient{
		authorization: *env.Authorization,
	}

	idTokenRequestUrl := fmt.Sprintf("https://%s:%d%s", oidcIssuerHost, oidcIssuerPort, oidcIssuerPath)
	idTokenRequestToken := test.DummyAccessToken

	opts := auth.GitHubOIDCAuthorizerOptions{
		Api:                 env.MicrosoftGraph,
		AuxiliaryTenantIds:  test.AuxiliaryTenantIds,
		ClientId:            "11111111-0000-0000-0000-000000000000",
		Environment:         *env,
		IdTokenRequestToken: idTokenRequestToken,
		IdTokenRequestUrl:   idTokenRequestUrl,
		TenantId:            "00000000-1111-0000-0000-000000000000",
	}

	authorizer, err := auth.NewGitHubOIDCAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewGitHubOIDCAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccGitHubOIDCAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.GitHubTokenURL == "" {
		t.Skip("test.GitHubTokenURL was empty")
	}
	if test.GitHubToken == "" {
		t.Skip("test.GitHubToken was empty")
	}

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.GitHubOIDCAuthorizerOptions{
		Api:                 env.MicrosoftGraph,
		AuxiliaryTenantIds:  test.AuxiliaryTenantIds,
		ClientId:            test.ClientId,
		Environment:         *env,
		TenantId:            test.TenantId,
		IdTokenRequestUrl:   test.GitHubTokenURL,
		IdTokenRequestToken: test.GitHubToken,
	}

	authorizer, err := auth.NewGitHubOIDCAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewGitHubOIDCAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

const (
	oidcIssuerHost = "github-oidc-issuer"
	oidcIssuerPort = 1234
	oidcIssuerPath = "/vend-id-token"
)

type oidcMockClient struct {
	authorization environments.Authorization
}

func (c *oidcMockClient) Do(r *http.Request) (resp *http.Response, err error) {
	if r == nil {
		return nil, fmt.Errorf("request was nil")
	}

	switch r.Host {
	case fmt.Sprintf("%s:%d", oidcIssuerHost, oidcIssuerPort):
		if r.URL.Path != oidcIssuerPath {
			return nil, fmt.Errorf("unexpected URL path, expected %q, received %q", oidcIssuerPath, r.URL.Path)
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

		resp.Body = io.NopCloser(bytes.NewBufferString(fmt.Sprintf(`{"count":1,"value":"%s"}`, test.DummyIDToken)))

		return
	}

	client := &test.AzureADAccessTokenMockClient{
		Authorization: c.authorization,
	}

	return client.Do(r)
}
