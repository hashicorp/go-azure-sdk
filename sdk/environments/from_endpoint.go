package environments

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/internal/metadata"
)

// FromEndpoint attempts to load an environment from the given Endpoint.
func FromEndpoint(ctx context.Context, endpoint string) (*Environment, error) {
	env := baseEnvironmentWithName("FromEnvironment")

	client := metadata.NewClientWithEndpoint(endpoint)
	config, err := client.GetMetaData(ctx)
	if err != nil {
		return nil, fmt.Errorf("retrieving metadata from endpoint %q: %+v", endpoint, err)
	}

	if config.Name == "" {
		return nil, fmt.Errorf("retrieving metadata from endpoint: `name` was nil")
	}
	env.Name = config.Name

	// TODO: support mapping the Microsoft Graph endpoint once this is returned (~Feb 2023)
	if config.ResourceManagerEndpoint == "" {
		return nil, fmt.Errorf("retrieving metadata from endpoint: no `resourceManagerEndpoint` was returned")
	}
	env.ResourceManager = ResourceManagerAPI(config.ResourceManagerEndpoint)

	if err := env.updateFromMetaData(config); err != nil {
		return nil, fmt.Errorf("updating Environment from MetaData: %+v", err)
	}

	return &env, nil
}
