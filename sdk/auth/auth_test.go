package auth_test

import (
	"context"
	"fmt"
	"math/rand"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/internal/test"
	"github.com/hashicorp/go-azure-sdk/internal/utils"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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
	environment           = envDefault("ARM_ENVIRONMENT", "global")
	gitHubTokenURL        = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	gitHubToken           = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	idToken               = os.Getenv("ARM_OIDC_TOKEN")
	msiEndpoint           = os.Getenv("ARM_MSI_ENDPOINT")
	msiToken              = os.Getenv("ARM_MSI_TOKEN")
)

func envDefault(key, def string) (ret string) {
	ret = os.Getenv(key)
	if ret == "" {
		ret = def
	}
	return
}

func TestClientCertificateAuthorizerV1(t *testing.T) {
	ctx := context.Background()
	testClientCertificateAuthorizer(ctx, t, auth.TokenVersion1)
}

func TestClientCertificateAuthorizerV2(t *testing.T) {
	ctx := context.Background()
	testClientCertificateAuthorizer(ctx, t, auth.TokenVersion2)
}

func testClientCertificateAuthorizer(ctx context.Context, t *testing.T, tokenVersion auth.TokenVersion) (token *oauth2.Token) {
	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	pfx := utils.Base64DecodeCertificate(clientCertificate)

	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, env, env.MSGraph, tokenVersion, tenantId, []string{}, clientId, pfx, clientCertificatePath, clientCertPassword)
	if err != nil {
		t.Fatalf("NewClientCertificateAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err = authorizer.Token()
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
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
	testClientSecretAuthorizer(ctx, t, auth.TokenVersion1)
}

func TestClientSecretAuthorizerV2(t *testing.T) {
	ctx := context.Background()
	testClientSecretAuthorizer(ctx, t, auth.TokenVersion2)
}

func testClientSecretAuthorizer(ctx context.Context, t *testing.T, tokenVersion auth.TokenVersion) (token *oauth2.Token) {
	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	authorizer, err := auth.NewClientSecretAuthorizer(ctx, env, env.MSGraph, tokenVersion, tenantId, []string{}, clientId, clientSecret)
	if err != nil {
		t.Fatalf("NewClientSecretAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err = authorizer.Token()
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
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
	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	authorizer, err := auth.NewAzureCliAuthorizer(ctx, env.MSGraph, tenantId)
	if err != nil {
		t.Fatalf("NewAzureCliAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err = authorizer.Token()
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
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
	port := 8000 + rand.Intn(999)
	if msiToken != "" {
		msiEndpoint = fmt.Sprintf("http://localhost:%d/metadata/identity/oauth2/token", port)
		done := test.MsiStubServer(ctx, port, msiToken)
		defer func() {
			done <- true
		}()
	}

	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	auth, err := auth.NewMsiAuthorizer(ctx, env.MSGraph, msiEndpoint, clientId)
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

func TestOIDCAuthorizer(t *testing.T) {
	if idToken == "" {
		t.Skip("idToken was empty")
	}

	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	authorizer, err := auth.NewOIDCAuthorizer(context.Background(), env, env.MSGraph, tenantId, []string{}, clientId, idToken)
	if err != nil {
		t.Fatalf("NewOIDCAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token()
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
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

	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	authorizer, err := auth.NewGitHubOIDCAuthorizer(context.Background(), env, env.MSGraph, tenantId, []string{}, clientId, gitHubTokenURL, gitHubToken)
	if err != nil {
		t.Fatalf("NewGitHubOIDCAuthorizer(): %v", err)
	}
	if authorizer == nil {
		t.Fatal("authorizer is nil, expected Authorizer")
	}

	token, err := authorizer.Token()
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
	}
	if token == nil {
		t.Fatal("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatal("token.AccessToken was empty")
	}
}
