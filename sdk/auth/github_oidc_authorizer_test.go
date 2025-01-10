// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestGitHubOIDCAuthorizer(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	mockHost := "github-oidc-issuer"
	auth.Client = &oidcMockClient{
		authorization: *env.Authorization,
		platform:      OidcMockClientPlatformGithub,
		mockHost:      mockHost,
	}

	idTokenRequestUrl := fmt.Sprintf("https://%s/vend-id-token", mockHost)
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

	if test.OIDCRequestURL == "" {
		t.Skip("test.OIDCRequestURL was empty")
	}
	if test.OIDCRequestToken == "" {
		t.Skip("test.OIDCRequestToken was empty")
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
		IdTokenRequestUrl:   test.OIDCRequestURL,
		IdTokenRequestToken: test.OIDCRequestToken,
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
