package auth

import (
	"context"
	"fmt"
	"os"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
	"golang.org/x/crypto/pkcs12"
)

type ClientCertificateAuthorizerOptions struct {
	// TODO: document these

	// Environment is a reference to the Azure Environment/Cloud being targeted.
	Environment environments.Environment

	// Api describes the Azure API being used
	Api          environments.Api
	TenantId     string
	AuxTenantIds []string
	ClientId     string
	Pkcs12Data   []byte
	Pkcs12Path   string
	Pkcs12Pass   string
}

// NewClientCertificateAuthorizer returns an authorizer which uses client certificate authentication.
func NewClientCertificateAuthorizer(ctx context.Context, options ClientCertificateAuthorizerOptions) (Authorizer, error) {
	if len(options.Pkcs12Data) == 0 {
		var err error
		options.Pkcs12Data, err = os.ReadFile(options.Pkcs12Path)
		if err != nil {
			return nil, fmt.Errorf("could not read PKCS#12 archive at %q: %s", options.Pkcs12Path, err)
		}
	}

	key, cert, err := pkcs12.Decode(options.Pkcs12Data, options.Pkcs12Pass)
	if err != nil {
		return nil, fmt.Errorf("could not decode PKCS#12 archive: %s", err)
	}

	conf := clientCredentialsConfig{
		Environment:        options.Environment,
		TenantID:           options.TenantId,
		AuxiliaryTenantIDs: options.AuxTenantIds,
		ClientID:           options.ClientId,
		PrivateKey:         key,
		Certificate:        cert,
		ResourceUrl:        options.Api.ResourceUrl(),
		Scopes:             []string{options.Api.DefaultScope()},
	}
	return conf.TokenSource(ctx, clientCredentialsAssertionType)
}
