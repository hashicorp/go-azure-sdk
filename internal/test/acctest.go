package test

import (
	"context"
	"log"
	"os"
	"testing"

	"github.com/hashicorp/go-azure-sdk/internal/utils"
	"github.com/hashicorp/go-azure-sdk/sdk/auth"
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

func envDefault(envVarName, defaultValue string) string {
	if v := os.Getenv(envVarName); v != "" {
		return v
	}
	return defaultValue
}

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
)

type Connection struct {
	AuthConfig *auth.Config
	Authorizer auth.Authorizer
	Claims     *auth.Claims
}

// NewConnection configures and returns a Connection for use in tests.
func NewConnection(tokenVersion auth.TokenVersion) *Connection {
	env, err := environments.FromNamed(environment)
	if err != nil {
		log.Fatal(err)
	}

	t := Connection{
		AuthConfig: &auth.Config{
			Environment:            env,
			Version:                tokenVersion,
			TenantID:               tenantId,
			ClientID:               clientId,
			ClientCertData:         utils.Base64DecodeCertificate(clientCertificate),
			ClientCertPath:         clientCertificatePath,
			ClientCertPassword:     clientCertPassword,
			ClientSecret:           clientSecret,
			IDTokenRequestURL:      gitHubTokenURL,
			IDTokenRequestToken:    gitHubToken,
			IDToken:                idToken,
			MsiEndpoint:            msiEndpoint,
			EnableClientCertAuth:   true,
			EnableClientSecretAuth: true,
			EnableGitHubOIDCAuth:   true,
		},
	}

	return &t
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
	c.Claims, err = auth.ParseClaims(token)
	if err != nil {
		log.Fatalf("parsing token claims: %v", err)
	}
}
