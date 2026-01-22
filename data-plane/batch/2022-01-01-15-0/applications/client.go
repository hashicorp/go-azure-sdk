package applications

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationsClient struct {
	Client *dataplane.Client
}

func NewApplicationsClientUnconfigured() (*ApplicationsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "applications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationsClient: %+v", err)
	}

	return &ApplicationsClient{
		Client: client,
	}, nil
}

func (c *ApplicationsClient) ApplicationsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewApplicationsClientWithBaseURI(endpoint string) (*ApplicationsClient, error) {
	client, err := dataplane.NewClient(endpoint, "applications", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ApplicationsClient: %+v", err)
	}

	return &ApplicationsClient{
		Client: client,
	}, nil
}
