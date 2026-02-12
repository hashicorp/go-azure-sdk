package keyvalues

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KeyValuesClient struct {
	Client *dataplane.Client
}

func NewKeyValuesClientUnconfigured() (*KeyValuesClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "keyvalues", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KeyValuesClient: %+v", err)
	}

	return &KeyValuesClient{
		Client: client,
	}, nil
}

func (c *KeyValuesClient) KeyValuesClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewKeyValuesClientWithBaseURI(endpoint string) (*KeyValuesClient, error) {
	client, err := dataplane.NewClient(endpoint, "keyvalues", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating KeyValuesClient: %+v", err)
	}

	return &KeyValuesClient{
		Client: client,
	}, nil
}

func (c *KeyValuesClient) Clone(endpoint string) *KeyValuesClient {
	return &KeyValuesClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
