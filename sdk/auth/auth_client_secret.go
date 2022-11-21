package auth

import (
	"context"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type ClientSecretAuthorizerOptions struct {
	// TODO: document these

	Environment  environments.Environment
	Api          environments.Api
	TokenVersion TokenVersion
	TenantId     string
	AuxTenantIds []string
	ClientId     string
	ClientSecret string
}

// NewClientSecretAuthorizer returns an authorizer which uses client secret authentication.
func NewClientSecretAuthorizer(ctx context.Context, options ClientSecretAuthorizerOptions) (Authorizer, error) {
	conf := clientCredentialsConfig{
		Environment:        options.Environment,
		TenantID:           options.TenantId,
		AuxiliaryTenantIDs: options.AuxTenantIds,
		ClientID:           options.ClientId,
		ClientSecret:       options.ClientSecret,
		ResourceUrl:        options.Api.ResourceUrl(),
		Scopes:             []string{options.Api.DefaultScope()},
		TokenVersion:       options.TokenVersion,
	}

	return conf.TokenSource(ctx, ClientCredentialsSecretType), nil
}
