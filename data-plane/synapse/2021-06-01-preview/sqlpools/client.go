package sqlpools

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolsClient struct {
	Client *dataplane.Client
}

func NewSqlPoolsClientUnconfigured() (*SqlPoolsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "sqlpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsClient: %+v", err)
	}

	return &SqlPoolsClient{
		Client: client,
	}, nil
}

func (c *SqlPoolsClient) SqlPoolsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewSqlPoolsClientWithBaseURI(endpoint string) (*SqlPoolsClient, error) {
	client, err := dataplane.NewClient(endpoint, "sqlpools", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlPoolsClient: %+v", err)
	}

	return &SqlPoolsClient{
		Client: client,
	}, nil
}

func (c *SqlPoolsClient) Clone(endpoint string) *SqlPoolsClient {
	return &SqlPoolsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
