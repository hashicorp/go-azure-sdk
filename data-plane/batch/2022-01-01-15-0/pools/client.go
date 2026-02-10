package pools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolsClient struct {
	Client *dataplane.Client
}

func NewPoolsClientUnconfigured() (*PoolsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "pools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PoolsClient: %+v", err)
	}

	return &PoolsClient{
		Client: client,
	}, nil
}

func (c *PoolsClient) PoolsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewPoolsClientWithBaseURI(endpoint string) (*PoolsClient, error) {
	client, err := dataplane.NewClient(endpoint, "pools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PoolsClient: %+v", err)
	}

	return &PoolsClient{
		Client: client,
	}, nil
}

func (c *PoolsClient) Clone(endpoint string) *PoolsClient {
	return &PoolsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
