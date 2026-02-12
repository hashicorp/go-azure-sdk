package integrationruntimes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegrationRuntimesClient struct {
	Client *dataplane.Client
}

func NewIntegrationRuntimesClientUnconfigured() (*IntegrationRuntimesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "integrationruntimes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationRuntimesClient: %+v", err)
	}

	return &IntegrationRuntimesClient{
		Client: client,
	}, nil
}

func (c *IntegrationRuntimesClient) IntegrationRuntimesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewIntegrationRuntimesClientWithBaseURI(endpoint string) (*IntegrationRuntimesClient, error) {
	client, err := dataplane.NewClient(endpoint, "integrationruntimes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IntegrationRuntimesClient: %+v", err)
	}

	return &IntegrationRuntimesClient{
		Client: client,
	}, nil
}

func (c *IntegrationRuntimesClient) Clone(endpoint string) *IntegrationRuntimesClient {
	return &IntegrationRuntimesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
