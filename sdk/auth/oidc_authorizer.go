package auth

import (
	"context"
	"fmt"

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
	scope, err := environments.Scope(options.Api)
	if err != nil {
		return nil, fmt.Errorf("determining scope for %q: %+v", options.Api.Name(), err)
	}

	conf := clientCredentialsConfig{
		Environment:        options.Environment,
		TenantID:           options.TenantId,
		AuxiliaryTenantIDs: options.AuxiliaryTenantIds,
		ClientID:           options.ClientId,
		FederatedAssertion: options.FederatedAssertion,
		Scopes: []string{
			*scope,
		},
	}

	return conf.TokenSource(ctx, clientCredentialsAssertionType)
}
