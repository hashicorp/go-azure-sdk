package auth

import (
	"context"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type GitHubOIDCAuthorizerOptions struct {
	// TODO: document these

	Api                 environments.Api
	AuxiliaryTenantIds  []string
	ClientId            string
	Environment         environments.Environment
	TenantId            string
	IdTokenRequestUrl   string
	IdTokenRequestToken string
}

// NewGitHubOIDCAuthorizer returns an authorizer which acquires a client assertion from a GitHub endpoint, then uses client assertion authentication to obtain an access token.
func NewGitHubOIDCAuthorizer(ctx context.Context, options GitHubOIDCAuthorizerOptions) (Authorizer, error) {
	conf := GitHubOIDCConfig{
		Environment:         options.Environment,
		TenantID:            options.TenantId,
		AuxiliaryTenantIDs:  options.AuxiliaryTenantIds,
		ClientID:            options.ClientId,
		IDTokenRequestURL:   options.IdTokenRequestUrl,
		IDTokenRequestToken: options.IdTokenRequestToken,
		Scopes:              []string{options.Api.DefaultScope()},
	}

	return conf.TokenSource(ctx), nil
}
