package test

import (
	"context"
	"net/http"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/claims"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

func AccTest(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping acceptance test when `-short` is set")
	}

	if os.Getenv("ACCTEST") == "" {
		t.Skip("skipping acceptance test, ACCTEST is not set")
	}
}

type Connection struct {
	AuthConfig auth.Credentials
	Authorizer auth.Authorizer
	Claims     *claims.Claims
}

// NewConnection configures and returns a Connection for use in tests.
func NewConnection(t *testing.T) *Connection {
	env, err := environments.FromName(Environment)
	if err != nil {
		t.Fatal(err)
	}

	return &Connection{
		AuthConfig: auth.Credentials{
			Environment:                                *env,
			TenantID:                                   TenantId,
			ClientID:                                   ClientId,
			ClientCertificateData:                      Base64DecodeCertificate(t, ClientCertificate),
			ClientCertificatePath:                      ClientCertificatePath,
			ClientCertificatePassword:                  ClientCertPassword,
			ClientSecret:                               ClientSecret,
			GitHubOIDCTokenRequestURL:                  GitHubTokenURL,
			GitHubOIDCTokenRequestToken:                GitHubToken,
			OIDCAssertionToken:                         IdToken,
			CustomManagedIdentityEndpoint:              CustomManagedIdentityEndpoint,
			EnableAuthenticatingUsingClientCertificate: true,
			EnableAuthenticatingUsingClientSecret:      true,
			EnableAuthenticationUsingGitHubOIDC:        true,
		},
	}
}

// Authorize configures an Authorizer for the Connection
func (c *Connection) Authorize(ctx context.Context, t *testing.T, api environments.Api) {
	var err error
	c.Authorizer, err = auth.NewAuthorizerFromCredentials(ctx, c.AuthConfig, api)
	if err != nil {
		t.Fatalf("building authorizer from credentials: %+v", err)
	}
	token, err := c.Authorizer.Token(ctx, &http.Request{})
	if err != nil {
		t.Fatalf("acquiring access token: %v", err)
	}
	c.Claims, err = claims.ParseClaims(token)
	if err != nil {
		t.Fatalf("parsing token claims: %v", err)
	}
}
