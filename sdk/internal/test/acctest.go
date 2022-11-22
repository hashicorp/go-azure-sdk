package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/sdk/auth"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"github.com/hashicorp/go-azure-sdk/sdk/internal/claims"
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
	AuthConfig *auth.Config
	Authorizer auth.Authorizer
	Claims     *claims.Claims
}

// NewConnection configures and returns a Connection for use in tests.
func NewConnection(t *testing.T, tokenVersion auth.TokenVersion) *Connection {
	env, err := environments.FromNamed(Environment)
	if err != nil {
		t.Fatal(err)
	}

	connection := Connection{
		AuthConfig: &auth.Config{
			Environment:                   *env,
			Version:                       tokenVersion,
			TenantID:                      TenantId,
			ClientID:                      ClientId,
			ClientCertData:                Base64DecodeCertificate(t, ClientCertificate),
			ClientCertPath:                ClientCertificatePath,
			ClientCertPassword:            ClientCertPassword,
			ClientSecret:                  ClientSecret,
			IDTokenRequestURL:             GitHubTokenURL,
			IDTokenRequestToken:           GitHubToken,
			IDToken:                       IdToken,
			CustomManagedIdentityEndpoint: CustomManagedIdentityEndpoint,
			EnableClientCertAuth:          true,
			EnableClientSecretAuth:        true,
			EnableGitHubOIDCAuth:          true,
		},
	}

	return &connection
}

// Authorize configures an Authorizer for the Connection
func (c *Connection) Authorize(ctx context.Context, api environments.Api) {
	var err error
	c.Authorizer, err = c.AuthConfig.NewAuthorizer(ctx, api)
	if err != nil {
		log.Fatal(err)
	}
	token, err := c.Authorizer.Token()
	if err != nil {
		log.Fatalf("acquiring access token: %v", err)
	}
	c.Claims, err = claims.ParseClaims(token)
	if err != nil {
		log.Fatalf("parsing token claims: %v", err)
	}
}
