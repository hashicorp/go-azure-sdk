package triggers

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggersClient struct {
	Client *dataplane.Client
}

func NewTriggersClientUnconfigured() (*TriggersClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "triggers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggersClient: %+v", err)
	}

	return &TriggersClient{
		Client: client,
	}, nil
}

func (c *TriggersClient) TriggersClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewTriggersClientWithBaseURI(endpoint string) (*TriggersClient, error) {
	client, err := dataplane.NewClient(endpoint, "triggers", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggersClient: %+v", err)
	}

	return &TriggersClient{
		Client: client,
	}, nil
}

func (c *TriggersClient) Clone(endpoint string) *TriggersClient {
	return &TriggersClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
