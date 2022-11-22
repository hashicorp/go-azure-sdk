package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccOIDCAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.IdToken == "" {
		t.Skip("test.IdToken was empty")
	}

	env, err := environments.FromNamed(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.OIDCAuthorizerOptions{
		Environment:        *env,
		Api:                env.MSGraph,
		TenantId:           test.TenantId,
		AuxiliaryTenantIds: []string{},
		ClientId:           test.ClientId,
		FederatedAssertion: test.IdToken,
	}
	authorizer, err := auth.NewOIDCAuthorizer(context.Background(), opts)
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