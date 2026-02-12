package sparkjobdefinitions

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobDefinitionsClient struct {
	Client *dataplane.Client
}

func NewSparkJobDefinitionsClientUnconfigured() (*SparkJobDefinitionsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "sparkjobdefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SparkJobDefinitionsClient: %+v", err)
	}

	return &SparkJobDefinitionsClient{
		Client: client,
	}, nil
}

func (c *SparkJobDefinitionsClient) SparkJobDefinitionsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewSparkJobDefinitionsClientWithBaseURI(endpoint string) (*SparkJobDefinitionsClient, error) {
	client, err := dataplane.NewClient(endpoint, "sparkjobdefinitions", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SparkJobDefinitionsClient: %+v", err)
	}

	return &SparkJobDefinitionsClient{
		Client: client,
	}, nil
}

func (c *SparkJobDefinitionsClient) Clone(endpoint string) *SparkJobDefinitionsClient {
	return &SparkJobDefinitionsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
