// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestOIDCAuthorizer(t *testing.T) {
	if test.DummyIDToken == "" {
		t.Skip("test.DummyIDToken was empty")
	}

	ctx := context.Background()
	env := environments.AzurePublic()

	auth.Client = &test.AzureADAccessTokenMockClient{
		Authorization: *env.Authorization,
	}

	opts := auth.OIDCAuthorizerOptions{
		Environment:        *env,
		Api:                env.MicrosoftGraph,
		TenantId:           "00000000-1111-0000-0000-000000000000",
		AuxiliaryTenantIds: test.AuxiliaryTenantIds,
		ClientId:           "11111111-0000-0000-0000-000000000000",
		FederatedAssertion: test.DummyIDToken,
	}

	authorizer, err := auth.NewOIDCAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewOIDCAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccOIDCAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.IdToken == "" {
		t.Skip("test.IdToken was empty")
	}

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.OIDCAuthorizerOptions{
		Environment:        *env,
		Api:                env.MicrosoftGraph,
		TenantId:           test.TenantId,
		AuxiliaryTenantIds: test.AuxiliaryTenantIds,
		ClientId:           test.ClientId,
		FederatedAssertion: test.IdToken,
	}

	authorizer, err := auth.NewOIDCAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewOIDCAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}
