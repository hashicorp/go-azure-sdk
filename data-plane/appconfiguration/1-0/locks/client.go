package locks

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocksClient struct {
	Client *dataplane.Client
}

func NewLocksClientUnconfigured() (*LocksClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "locks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocksClient: %+v", err)
	}

	return &LocksClient{
		Client: client,
	}, nil
}

func (c *LocksClient) LocksClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewLocksClientWithBaseURI(endpoint string) (*LocksClient, error) {
	client, err := dataplane.NewClient(endpoint, "locks", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LocksClient: %+v", err)
	}

	return &LocksClient{
		Client: client,
	}, nil
}

func (c *LocksClient) Clone(endpoint string) *LocksClient {
	return &LocksClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
