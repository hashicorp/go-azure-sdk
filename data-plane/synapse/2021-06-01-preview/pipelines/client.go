package pipelines

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelinesClient struct {
	Client *dataplane.Client
}

func NewPipelinesClientUnconfigured() (*PipelinesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "pipelines", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelinesClient: %+v", err)
	}

	return &PipelinesClient{
		Client: client,
	}, nil
}

func (c *PipelinesClient) PipelinesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewPipelinesClientWithBaseURI(endpoint string) (*PipelinesClient, error) {
	client, err := dataplane.NewClient(endpoint, "pipelines", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating PipelinesClient: %+v", err)
	}

	return &PipelinesClient{
		Client: client,
	}, nil
}

func (c *PipelinesClient) Clone(endpoint string) *PipelinesClient {
	return &PipelinesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
