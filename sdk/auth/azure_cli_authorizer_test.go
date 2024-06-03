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

func TestAccAzureCliAuthorizer(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.AzureCliAuthorizerOptions{
		Api: env.MicrosoftGraph,
	}

	authorizer, err := auth.NewAzureCliAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewAzureCliAuthorizer(): %v", err)
	}

	cliAuth, err := testCheckAzureCliAuthorizer(authorizer)
	if err != nil {
		t.Fatal(err)
	}

	if cliAuth.TenantID == "" {
		t.Fatal("cliAuth.TenantID has unexpected empty value (should have been auto-detected)")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccAzureCliAuthorizerWithSubscription(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.AzureCliAuthorizerOptions{
		Api:                env.ResourceManager,
		SubscriptionIdHint: test.SubscriptionId,
	}

	authorizer, err := auth.NewAzureCliAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewAzureCliAuthorizer(): %v", err)
	}

	cliAuth, err := testCheckAzureCliAuthorizer(authorizer)
	if err != nil {
		t.Fatal(err)
	}

	if cliAuth.SubscriptionIDHint != test.SubscriptionId {
		t.Fatalf("cliAuth.SubscriptionIDHint has unexpected value %q", cliAuth.SubscriptionIDHint)
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccAzureCliAuthorizerWithTenant(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.AzureCliAuthorizerOptions{
		Api:      env.MicrosoftGraph,
		TenantId: test.TenantId,
	}

	authorizer, err := auth.NewAzureCliAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewAzureCliAuthorizer(): %v", err)
	}

	cliAuth, err := testCheckAzureCliAuthorizer(authorizer)
	if err != nil {
		t.Fatal(err)
	}

	if cliAuth.TenantID != test.TenantId {
		t.Fatalf("cliAuth.TenantID has unexpected value %q", cliAuth.TenantID)
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func testCheckAzureCliAuthorizer(authorizer auth.Authorizer) (*auth.AzureCliAuthorizer, error) {
	if authorizer == nil {
		return nil, fmt.Errorf("authorizer is nil, expected Authorizer")
	}

	cachedAuth, ok := authorizer.(*auth.CachedAuthorizer)
	if !ok {
		return nil, fmt.Errorf("authorizer is not a *CachedAuthorizer")
	}

	cliAuth, ok := cachedAuth.Source.(*auth.AzureCliAuthorizer)
	if !ok {
		return nil, fmt.Errorf("cachedAuth.Source is not an *AzureCliAuthorizer")
	}

	if cliAuth.DefaultSubscriptionID == "" {
		return cliAuth, fmt.Errorf("cliAuth.DefaultSubscriptionID has unexpected empty value (should have been auto-detected)")
	}

	return cliAuth, nil
}
