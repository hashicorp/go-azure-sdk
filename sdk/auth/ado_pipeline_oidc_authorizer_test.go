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

func TestADOPipelineOIDCAuthorizer(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	mockHost := "ado-oidc-issuer"
	auth.Client = &oidcMockClient{
		authorization: *env.Authorization,
		platform:      OidcMockClientPlatformADOPipeline,
		mockHost:      mockHost,
	}

	idTokenRequestUrl := fmt.Sprintf("https://%s/vend-id-token", mockHost)
	idTokenRequestToken := test.DummyAccessToken

	opts := auth.ADOPipelineOIDCAuthorizerOptions{
		Api:                 env.MicrosoftGraph,
		AuxiliaryTenantIds:  test.AuxiliaryTenantIds,
		ClientId:            "11111111-0000-0000-0000-000000000000",
		Environment:         *env,
		IdTokenRequestToken: idTokenRequestToken,
		IdTokenRequestUrl:   idTokenRequestUrl,
		ServiceConnectionId: "test-service-connection",
		TenantId:            "00000000-1111-0000-0000-000000000000",
	}

	authorizer, err := auth.NewADOPipelineOIDCAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewADOPipelineOIDCAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestAccADOPipelineOIDCAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.ADOPipelineToken == "" {
		t.Skip("test.ADOPipelineToken was empty")
	}
	if test.ADOPipelineTokenURL == "" {
		t.Skip("test.ADOPipelineTokenURL was empty")
	}
	if test.ADOServiceConnectionId == "" {
		t.Skip("test.ADOServiceConnectionId was empty")
	}

	ctx := context.Background()

	env, err := environments.FromName(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ADOPipelineOIDCAuthorizerOptions{
		Api:                 env.MicrosoftGraph,
		AuxiliaryTenantIds:  test.AuxiliaryTenantIds,
		ClientId:            test.ClientId,
		Environment:         *env,
		TenantId:            test.TenantId,
		IdTokenRequestUrl:   test.ADOPipelineTokenURL,
		IdTokenRequestToken: test.ADOPipelineToken,
		ServiceConnectionId: test.ADOServiceConnectionId,
	}

	authorizer, err := auth.NewADOPipelineOIDCAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewADOPipelineOIDCAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}
