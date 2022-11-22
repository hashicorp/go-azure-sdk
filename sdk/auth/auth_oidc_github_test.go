package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccGitHubOIDCAuthorizer(t *testing.T) {
	test.AccTest(t)

	if test.GitHubTokenURL == "" {
		t.Skip("test.GitHubTokenURL was empty")
	}
	if test.GitHubToken == "" {
		t.Skip("test.GitHubToken was empty")
	}

	env, err := environments.FromNamed(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.GitHubOIDCAuthorizerOptions{
		Api:                 env.MSGraph,
		AuxiliaryTenantIds:  nil,
		ClientId:            test.ClientId,
		Environment:         *env,
		TenantId:            test.TenantId,
		IdTokenRequestUrl:   test.GitHubTokenURL,
		IdTokenRequestToken: test.GitHubToken,
	}
	authorizer, err := auth.NewGitHubOIDCAuthorizer(context.Background(), opts)
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