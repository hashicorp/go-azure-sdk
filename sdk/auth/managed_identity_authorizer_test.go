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

func TestManagedIdentityAuthorizer(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	auth.MetadataClient = &test.AzureADAccessTokenMockClient{
		Authorization: *env.Authorization,
	}

	opts := auth.ManagedIdentityAuthorizerOptions{
		Api:      env.MicrosoftGraph,
		ClientId: "11111111-0000-0000-0000-0000000000000000",
	}

	authorizer, err := auth.NewManagedIdentityAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewManagedIdentityAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccManagedIdentityAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.CustomManagedIdentityEndpoint == "" {
		t.Skip("test.CustomManagedIdentityEndpoint was empty")
	}

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ManagedIdentityAuthorizerOptions{
		Api:                           env.MicrosoftGraph,
		ClientId:                      test.ClientId,
		CustomManagedIdentityEndpoint: test.CustomManagedIdentityEndpoint,
	}

	authorizer, err := auth.NewManagedIdentityAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewManagedIdentityAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}
