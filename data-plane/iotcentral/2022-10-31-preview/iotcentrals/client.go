package iotcentrals

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IotcentralsClient struct {
	Client *dataplane.Client
}

func NewIotcentralsClientUnconfigured() (*IotcentralsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "iotcentrals", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IotcentralsClient: %+v", err)
	}

	return &IotcentralsClient{
		Client: client,
	}, nil
}

func (c *IotcentralsClient) IotcentralsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewIotcentralsClientWithBaseURI(endpoint string) (*IotcentralsClient, error) {
	client, err := dataplane.NewClient(endpoint, "iotcentrals", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating IotcentralsClient: %+v", err)
	}

	return &IotcentralsClient{
		Client: client,
	}, nil
}

func (c *IotcentralsClient) Clone(endpoint string) *IotcentralsClient {
	return &IotcentralsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
