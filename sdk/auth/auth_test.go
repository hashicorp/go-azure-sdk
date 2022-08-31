package auth_test

import (
	"context"
	auth2 "github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/internal/test"
	"github.com/hashicorp/go-azure-sdk/internal/utils"
	"golang.org/x/oauth2"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var (
	tenantId              = os.Getenv("ARM_TENANT_ID")
	clientId              = os.Getenv("ARM_CLIENT_ID")
	clientCertificate     = os.Getenv("ARM_CLIENT_CERTIFICATE")
	clientCertificatePath = os.Getenv("ARM_CLIENT_CERTIFICATE_PATH")
	clientCertPassword    = os.Getenv("ARM_CLIENT_CERTIFICATE_PASSWORD")
	clientSecret          = os.Getenv("ARM_CLIENT_SECRET")
	environment           = os.Getenv("ARM_ENVIRONMENT")
	msiEndpoint           = os.Getenv("ARM_MSI_ENDPOINT")
	msiToken              = os.Getenv("ARM_MSI_TOKEN")

	gitHubTokenURL = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	gitHubToken    = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
)

func TestClientCertificateAuthorizerV1(t *testing.T) {
	ctx := context.Background()
	testClientCertificateAuthorizer(ctx, t, auth2.TokenVersion1)
}

func TestClientCertificateAuthorizerV2(t *testing.T) {
	ctx := context.Background()
	testClientCertificateAuthorizer(ctx, t, auth2.TokenVersion2)
}

func testClientCertificateAuthorizer(ctx context.Context, t *testing.T, tokenVersion auth2.TokenVersion) (token *oauth2.Token) {
	env, err := environments.EnvironmentFromString(environment)
	if err != nil {
		t.Fatal(err)
	}

	pfx := utils.Base64DecodeCertificate(clientCertificate)

	auth, err := auth2.NewClientCertificateAuthorizer(ctx, env, env.MsGraph, tokenVersion, tenantId, []string{}, clientId, pfx, clientCertificatePath, clientCertPassword)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}
	if auth == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	token, err = auth.Token()
	if err != nil {
		t.Fatalf("auth.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatal("token.AccessToken was empty")
	}

	return
}

func TestClientSecretAuthorizerV1(t *testing.T) {
	ctx := context.Background()
	testClientSecretAuthorizer(ctx, t, auth2.TokenVersion1)
}

func TestClientSecretAuthorizerV2(t *testing.T) {
	ctx := context.Background()
	testClientSecretAuthorizer(ctx, t, auth2.TokenVersion2)
}

func testClientSecretAuthorizer(ctx context.Context, t *testing.T, tokenVersion auth2.TokenVersion) (token *oauth2.Token) {
	env, err := environments.EnvironmentFromString(environment)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := auth2.NewClientSecretAuthorizer(ctx, env, env.MsGraph, tokenVersion, tenantId, []string{}, clientId, clientSecret)
	if err != nil {
		t.Fatalf("NewClientSecretAuthorizer(): %v", err)
	}
	if auth == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	token, err = auth.Token()
	if err != nil {
		t.Fatalf("auth.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatalf("token.AccessToken was empty")
	}

	return
}

func TestAzureCliAuthorizer(t *testing.T) {
	ctx := context.Background()
	testAzureCliAuthorizer(ctx, t)
}

func testAzureCliAuthorizer(ctx context.Context, t *testing.T) (token *oauth2.Token) {
	env, err := environments.EnvironmentFromString(environment)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := auth2.NewAzureCliAuthorizer(ctx, env.MsGraph, tenantId)
	if err != nil {
		t.Fatalf("NewAzureCliAuthorizer(): %v", err)
	}
	if auth == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	token, err = auth.Token()
	if err != nil {
		t.Fatalf("auth.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatalf("token.AccessToken was empty")
	}

	return
}

func TestMsiAuthorizer(t *testing.T) {
	ctx := context.Background()
	if msiToken != "" {
		msiEndpoint = "http://localhost:8080/metadata/identity/oauth2/token"
		done := test.MsiStubServer(ctx, 8080, msiToken)
		defer func() {
			done <- true
		}()
	}

	env, err := environments.EnvironmentFromString(environment)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := auth2.NewMsiAuthorizer(ctx, env.MsGraph, msiEndpoint, clientId)
	if err != nil {
		t.Fatalf("NewMsiAuthorizer(): %v", err)
	}
	if auth == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	token, err := auth.Token()
	if err != nil {
		t.Fatalf("auth.Token(): %v", err)
	}
	if token == nil {
		t.Fatal("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatal("token.AccessToken was empty")
	}
}

func TestGitHubOIDCAuthorizer(t *testing.T) {
	if gitHubTokenURL == "" {
		t.Skip("gitHubTokenURL was empty")
	}
	if gitHubToken == "" {
		t.Skip("gitHubToken was empty")
	}

	env, err := environments.EnvironmentFromString(environment)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := auth2.NewGitHubOIDCAuthorizer(context.Background(), env, env.MsGraph, tenantId, []string{}, clientId, gitHubTokenURL, gitHubToken)
	if err != nil {
		t.Fatalf("NewGitHubOIDCAuthorizer(): %v", err)
	}
	if auth == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	token, err := auth.Token()
	if err != nil {
		t.Fatalf("auth.Token(): %v", err)
	}
	if token == nil {
		t.Fatal("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatal("token.AccessToken was empty")
	}
}
