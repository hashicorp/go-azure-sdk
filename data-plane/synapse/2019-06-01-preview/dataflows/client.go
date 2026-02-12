package dataflows

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowsClient struct {
	Client *dataplane.Client
}

func NewDataFlowsClientUnconfigured() (*DataFlowsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "dataflows", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataFlowsClient: %+v", err)
	}

	return &DataFlowsClient{
		Client: client,
	}, nil
}

func (c *DataFlowsClient) DataFlowsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewDataFlowsClientWithBaseURI(endpoint string) (*DataFlowsClient, error) {
	client, err := dataplane.NewClient(endpoint, "dataflows", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataFlowsClient: %+v", err)
	}

	return &DataFlowsClient{
		Client: client,
	}, nil
}

func (c *DataFlowsClient) Clone(endpoint string) *DataFlowsClient {
	return &DataFlowsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
