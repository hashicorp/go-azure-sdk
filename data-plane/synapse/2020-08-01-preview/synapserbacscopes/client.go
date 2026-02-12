package synapserbacscopes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynapseRbacScopesClient struct {
	Client *dataplane.Client
}

func NewSynapseRbacScopesClientUnconfigured() (*SynapseRbacScopesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "synapserbacscopes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynapseRbacScopesClient: %+v", err)
	}

	return &SynapseRbacScopesClient{
		Client: client,
	}, nil
}

func (c *SynapseRbacScopesClient) SynapseRbacScopesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewSynapseRbacScopesClientWithBaseURI(endpoint string) (*SynapseRbacScopesClient, error) {
	client, err := dataplane.NewClient(endpoint, "synapserbacscopes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SynapseRbacScopesClient: %+v", err)
	}

	return &SynapseRbacScopesClient{
		Client: client,
	}, nil
}

func (c *SynapseRbacScopesClient) Clone(endpoint string) *SynapseRbacScopesClient {
	return &SynapseRbacScopesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
