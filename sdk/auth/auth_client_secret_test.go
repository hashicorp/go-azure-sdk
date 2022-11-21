package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
	"golang.org/x/oauth2"
)

func TestAccClientSecretAuthorizerV1(t *testing.T) {
	ctx := context.Background()
	testClientSecretAuthorizer(ctx, t, auth.TokenVersion1)
}

func TestAccClientSecretAuthorizerV2(t *testing.T) {
	ctx := context.Background()
	testClientSecretAuthorizer(ctx, t, auth.TokenVersion2)
}

func testClientSecretAuthorizer(ctx context.Context, t *testing.T, tokenVersion auth.TokenVersion) (token *oauth2.Token) {
	test.AccTest(t)

	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ClientSecretAuthorizerOptions{
		Environment:  *env,
		Api:          env.MSGraph,
		TokenVersion: tokenVersion,
		TenantId:     tenantId,
		AuxTenantIds: []string{},
		ClientId:     clientId,
		ClientSecret: clientSecret,
	}
	authorizer, err := auth.NewClientSecretAuthorizer(ctx, opts)
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
