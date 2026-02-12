package operationresult

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationResultClient struct {
	Client *dataplane.Client
}

func NewOperationResultClientUnconfigured() (*OperationResultClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "operationresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationResultClient: %+v", err)
	}

	return &OperationResultClient{
		Client: client,
	}, nil
}

func (c *OperationResultClient) OperationResultClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewOperationResultClientWithBaseURI(endpoint string) (*OperationResultClient, error) {
	client, err := dataplane.NewClient(endpoint, "operationresult", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating OperationResultClient: %+v", err)
	}

	return &OperationResultClient{
		Client: client,
	}, nil
}

func (c *OperationResultClient) Clone(endpoint string) *OperationResultClient {
	return &OperationResultClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
