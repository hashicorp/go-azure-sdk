package bigdatapools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BigDataPoolsClient struct {
	Client *dataplane.Client
}

func NewBigDataPoolsClientUnconfigured() (*BigDataPoolsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "bigdatapools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BigDataPoolsClient: %+v", err)
	}

	return &BigDataPoolsClient{
		Client: client,
	}, nil
}

func (c *BigDataPoolsClient) BigDataPoolsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewBigDataPoolsClientWithBaseURI(endpoint string) (*BigDataPoolsClient, error) {
	client, err := dataplane.NewClient(endpoint, "bigdatapools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating BigDataPoolsClient: %+v", err)
	}

	return &BigDataPoolsClient{
		Client: client,
	}, nil
}

func (c *BigDataPoolsClient) Clone(endpoint string) *BigDataPoolsClient {
	return &BigDataPoolsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
