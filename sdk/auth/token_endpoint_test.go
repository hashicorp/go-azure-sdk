// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package auth_test

import (
	"context"
	"strings"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestClientCertificateAuthorizer_TrailingSlashLoginEndpoint(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	// Simulate a custom cloud that returns a login endpoint with a trailing slash
	authorization := *env.Authorization
	authorization.LoginEndpoint = strings.TrimRight(authorization.LoginEndpoint, "/") + "/"
	env.Authorization = &authorization

	auth.Client = &test.AzureADAccessTokenMockClient{
		Authorization: authorization,
	}

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MicrosoftGraph,
		TenantId:     "00000000-1111-0000-0000-000000000000",
		AuxTenantIds: test.AuxiliaryTenantIds,
		ClientId:     "11111111-0000-0000-0000-000000000000",
		Pkcs12Data:   test.Base64DecodeCertificate(t, dummyClientCertificate),
		Pkcs12Pass:   "certpassword",
	}

	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}

	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	if _, err = testObtainAccessToken(ctx, authorizer); err != nil {
		t.Fatal(err)
	}
}

func TestClientSecretAuthorizer_TrailingSlashLoginEndpoint(t *testing.T) {
	ctx := context.Background()
	env := environments.AzurePublic()

	// Simulate a custom cloud that returns a login endpoint with a trailing slash
	authorization := *env.Authorization
	authorization.LoginEndpoint = strings.TrimRight(authorization.LoginEndpoint, "/") + "/"
	env.Authorization = &authorization

	auth.Client = &test.AzureADAccessTokenMockClient{
		Authorization: authorization,
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
