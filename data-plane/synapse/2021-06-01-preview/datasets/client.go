package datasets

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatasetsClient struct {
	Client *dataplane.Client
}

func NewDatasetsClientUnconfigured() (*DatasetsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "datasets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatasetsClient: %+v", err)
	}

	return &DatasetsClient{
		Client: client,
	}, nil
}

func (c *DatasetsClient) DatasetsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewDatasetsClientWithBaseURI(endpoint string) (*DatasetsClient, error) {
	client, err := dataplane.NewClient(endpoint, "datasets", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DatasetsClient: %+v", err)
	}

	return &DatasetsClient{
		Client: client,
	}, nil
}

func (c *DatasetsClient) Clone(endpoint string) *DatasetsClient {
	return &DatasetsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
