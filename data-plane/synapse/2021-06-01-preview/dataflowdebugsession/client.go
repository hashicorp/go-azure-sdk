package dataflowdebugsession

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlowDebugSessionClient struct {
	Client *dataplane.Client
}

func NewDataFlowDebugSessionClientUnconfigured() (*DataFlowDebugSessionClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "dataflowdebugsession", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataFlowDebugSessionClient: %+v", err)
	}

	return &DataFlowDebugSessionClient{
		Client: client,
	}, nil
}

func (c *DataFlowDebugSessionClient) DataFlowDebugSessionClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewDataFlowDebugSessionClientWithBaseURI(endpoint string) (*DataFlowDebugSessionClient, error) {
	client, err := dataplane.NewClient(endpoint, "dataflowdebugsession", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataFlowDebugSessionClient: %+v", err)
	}

	return &DataFlowDebugSessionClient{
		Client: client,
	}, nil
}

func (c *DataFlowDebugSessionClient) Clone(endpoint string) *DataFlowDebugSessionClient {
	return &DataFlowDebugSessionClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
