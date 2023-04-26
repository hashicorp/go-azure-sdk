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

func TestClientSecretAuthorizer(t *testing.T) {
	if test.DummyAccessToken == "" {
		t.Skip("test.DummyAccessToken was empty")
	}

	ctx := context.Background()
	env := environments.AzurePublic()

	auth.Client = &test.AzureADAccessTokenMockClient{
		Authorization: *env.Authorization,
	}

	opts := auth.ClientSecretAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     "00000000-1111-0000-0000-000000000000",
		AuxTenantIds: test.AuxiliaryTenantIds,
		ClientId:     "11111111-0000-0000-0000-000000000000",
		ClientSecret: "supersecret",
	}

	authorizer, err := auth.NewClientSecretAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientSecretAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccClientSecretAuthorizer(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ClientSecretAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: test.AuxiliaryTenantIds,
		ClientId:     test.ClientId,
		ClientSecret: test.ClientSecret,
	}

	authorizer, err := auth.NewClientSecretAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientSecretAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}
