package auth_test

import (
	"context"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/test"
	"golang.org/x/oauth2"
)

func TestAccClientCertificateAuthorizerV1(t *testing.T) {
	ctx := context.Background()
	testClientCertificateAuthorizer(ctx, t, auth.TokenVersion1)
}

func TestAccClientCertificateAuthorizerV2(t *testing.T) {
	ctx := context.Background()
	testClientCertificateAuthorizer(ctx, t, auth.TokenVersion2)
}

func testClientCertificateAuthorizer(ctx context.Context, t *testing.T, tokenVersion auth.TokenVersion) (token *oauth2.Token) {
	test.AccTest(t)

	env, err := environments.FromNamed(environment)
	if err != nil {
		t.Fatal(err)
	}

	pfx := test.Base64DecodeCertificate(t, clientCertificate)

	opts := auth.ClientCertificateAuthorizerOptions{
		Environment:  *env,
		Api:          env.MSGraph,
		TokenVersion: tokenVersion,
		TenantId:     tenantId,
		AuxTenantIds: []string{},
		ClientId:     clientId,
		Pkcs12Data:   pfx,
		Pkcs12Path:   clientCertPassword,
		Pkcs12Pass:   clientCertificatePath,
	}
	authorizer, err := auth.NewClientCertificateAuthorizer(ctx, opts)
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
