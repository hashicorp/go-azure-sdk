package auth

import (
	"context"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type ManagedIdentityAuthorizerOptions struct {
	// TODO: document me

	Api                           environments.Api
	ClientId                      string
	CustomManagedIdentityEndpoint string
}

// NewManagedIdentityAuthorizer returns an authorizer using a Managed Identity for authentication.
func NewManagedIdentityAuthorizer(ctx context.Context, options ManagedIdentityAuthorizerOptions) (Authorizer, error) {
	conf, err := NewManagedIdentityConfig(options.Api.ResourceUrl(), options.ClientId, options.CustomManagedIdentityEndpoint)
	if err != nil {
		return nil, err
	}
	return conf.TokenSource(ctx), nil
}
