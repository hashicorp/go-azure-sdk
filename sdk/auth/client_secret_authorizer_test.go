package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccClientSecretAuthorizerV1(t *testing.T) {
	test.AccTest(t)

	env, err := environments.FromNamed(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	opts := auth.ClientSecretAuthorizerOptions{
		Environment:  *env,
		Api:          env.MSGraph,
		TenantId:     test.TenantId,
		AuxTenantIds: []string{},
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

	token, err := authorizer.Token(ctx, nil)
	if err != nil {
		t.Fatalf("authorizer.Token(): %v", err)
	}
	if token == nil {
		t.Fatalf("token was nil")
	}
	if token.AccessToken == "" {
		t.Fatalf("token.AccessToken was empty")
	}
}
