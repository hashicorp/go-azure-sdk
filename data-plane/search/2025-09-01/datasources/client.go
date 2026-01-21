package datasources

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataSourcesClient struct {
	Client *dataplane.Client
}

func NewDataSourcesClientUnconfigured() (*DataSourcesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "", "datasources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSourcesClient: %+v", err)
	}

	return &DataSourcesClient{
		Client: client,
	}, nil
}

func (c *DataSourcesClient) DataSourcesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func (c *DataSourcesClient) DataSourcesClientSetAdditionalEndpoint(endpoint string) {
	c.Client.AdditionalEndpoint = endpoint
}

func NewDataSourcesClientWithBaseURI(endpoint string, additionalEndpoint string) (*DataSourcesClient, error) {
	client, err := dataplane.NewClient(endpoint, additionalEndpoint, "datasources", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating DataSourcesClient: %+v", err)
	}

	return &DataSourcesClient{
		Client: client,
	}, nil
}
