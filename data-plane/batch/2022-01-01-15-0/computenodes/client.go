package computenodes

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodesClient struct {
	Client *dataplane.Client
}

func NewComputeNodesClientUnconfigured() (*ComputeNodesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "computenodes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComputeNodesClient: %+v", err)
	}

	return &ComputeNodesClient{
		Client: client,
	}, nil
}

func (c *ComputeNodesClient) ComputeNodesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewComputeNodesClientWithBaseURI(endpoint string) (*ComputeNodesClient, error) {
	client, err := dataplane.NewClient(endpoint, "computenodes", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ComputeNodesClient: %+v", err)
	}

	return &ComputeNodesClient{
		Client: client,
	}, nil
}

func (c *ComputeNodesClient) Clone(endpoint string) *ComputeNodesClient {
	return &ComputeNodesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
