package auth_test

import (
	"context"
	"fmt"
	"math/rand"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
)

func TestAccManagedIdentityAuthorizer(t *testing.T) {
	test.AccTest(t)

	ctx := context.Background()
	port := 8000 + rand.Intn(999)
	managedIdentityEndpoint := test.CustomManagedIdentityEndpoint
	if test.ManagedIdentityToken != "" {
		managedIdentityEndpoint = fmt.Sprintf("http://localhost:%d/metadata/identity/oauth2/token", port)
		done := test.ManagedIdentityStubServer(ctx, port, test.ManagedIdentityToken)
		defer func() {
			done <- true
		}()
	}

	env, err := environments.FromNamed(test.Environment)
	if err != nil {
		t.Fatal(err)
	}

	opts := auth.ManagedIdentityAuthorizerOptions{
		Api:                           env.MSGraph,
		ClientId:                      test.ClientId,
		CustomManagedIdentityEndpoint: managedIdentityEndpoint,
	}
	auth, err := auth.NewManagedIdentityAuthorizer(ctx, opts)
	if err != nil {
		t.Fatalf("NewManagedIdentityAuthorizer(): %v", err)
	}
	if auth == nil {
		t.Fatal("auth is nil, expected Authorizer")
	}

	token, err := auth.Token(ctx)
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
