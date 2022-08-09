package client

import (
	"context"
	"fmt"
	"github.com/hashicorp/go-azure-sdk/auth"
	"github.com/hashicorp/go-azure-sdk/environments"
	"github.com/hashicorp/go-azure-sdk/internal/utils"
	"github.com/hashicorp/go-azure-sdk/odata"
	"log"
	"os"
	"testing"
)

func envDefault(envVarName, defaultValue string) string {
	if v := os.Getenv(envVarName); v != "" {
		return v
	}
	return defaultValue
}

var (
	tenantId              = envDefault("TENANT_ID", "6df54acb-f3cd-4734-85e3-7511ade57a02")
	tenantDomain          = envDefault("TENANT_DOMAIN", "hamiltontesting2.onmicrosoft.com")
	clientId              = envDefault("CLIENT_ID", "c072182f-ead2-4e94-aa7f-a5975558c945")
	clientCertificate     = os.Getenv("CLIENT_CERTIFICATE")
	clientCertificatePath = os.Getenv("CLIENT_CERTIFICATE_PATH")
	clientCertPassword    = os.Getenv("CLIENT_CERTIFICATE_PASSWORD")
	clientSecret          = os.Getenv("CLIENT_SECRET")
	environment           = os.Getenv("AZURE_ENVIRONMENT")
	idTokenRequestUrl     = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_URL")
	idTokenRequestToken   = os.Getenv("ACTIONS_ID_TOKEN_REQUEST_TOKEN")
	retryMax              = envDefault("RETRY_MAX", "14")
)

type Connection struct {
	AuthConfig *auth.Config
	Authorizer auth.Authorizer
}

// NewConnection configures and returns a Connection for use in tests.
func NewConnection(tokenVersion auth.TokenVersion) *Connection {
	env, err := environments.EnvironmentFromString(environment)
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
			IDTokenRequestURL:      idTokenRequestUrl,
			IDTokenRequestToken:    idTokenRequestToken,
			EnableClientCertAuth:   true,
			EnableClientSecretAuth: true,
			EnableAzureCliToken:    true,
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
}

func TestClient(t *testing.T) {
	conn := NewConnection(auth.TokenVersion2)
	conn.Authorize(context.TODO(), conn.AuthConfig.Environment.MsGraph)

	c := NewClient(environments.MsGraphGlobalEndpoint)
	c.Authorizer = conn.Authorizer

	req, err := c.NewGetRequest(context.TODO(), "/v1.0/servicePrincipals", odata.Query{})
	if err != nil {
		t.Fatal(err)
	}

	resp, o, status, err := req.ExecutePaged()
	if err != nil {
		t.Fatal(err)
	}

	fmt.Printf("%#v\n\n", resp)
	fmt.Printf("%#v\n\n", o)
	fmt.Printf("%d\n\n", status)
}
