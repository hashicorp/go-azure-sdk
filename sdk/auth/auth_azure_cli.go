package auth

import (
	"context"

	"github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type AzureCliAuthorizerOptions struct {
	// TODO: document these

	Api      environments.Api
	TenantId string
}

// NewAzureCliAuthorizer returns an Authorizer which authenticates using the Azure CLI.
func NewAzureCliAuthorizer(ctx context.Context, options AzureCliAuthorizerOptions) (Authorizer, error) {
	conf, err := newAzureCliConfig(options.Api, options.TenantId)
	if err != nil {
		return nil, err
	}
	return conf.TokenSource(ctx), nil
}
