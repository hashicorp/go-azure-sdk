package auth

import (
	"context"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type OIDCAuthorizerOptions struct {
	// TODO: document me

	Environment        environments.Environment
	Api                environments.Api
	TenantId           string
	AuxiliaryTenantIds []string
	ClientId           string
	FederatedAssertion string
}

// NewOIDCAuthorizer returns an authorizer which uses OIDC authentication (federated client credentials)
func NewOIDCAuthorizer(ctx context.Context, options OIDCAuthorizerOptions) (Authorizer, error) {
	conf := clientCredentialsConfig{
		Environment:        options.Environment,
		TenantID:           options.TenantId,
		AuxiliaryTenantIDs: options.AuxiliaryTenantIds,
		ClientID:           options.ClientId,
		FederatedAssertion: options.FederatedAssertion,
		ResourceUrl:        options.Api.ResourceUrl(),
		Scopes:             []string{options.Api.DefaultScope()},
	}

	return conf.TokenSource(ctx, clientCredentialsAssertionType), nil
}
