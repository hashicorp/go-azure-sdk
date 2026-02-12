package trigger

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerClient struct {
	Client *dataplane.Client
}

func NewTriggerClientUnconfigured() (*TriggerClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "trigger", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggerClient: %+v", err)
	}

	return &TriggerClient{
		Client: client,
	}, nil
}

func (c *TriggerClient) TriggerClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewTriggerClientWithBaseURI(endpoint string) (*TriggerClient, error) {
	client, err := dataplane.NewClient(endpoint, "trigger", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggerClient: %+v", err)
	}

	return &TriggerClient{
		Client: client,
	}, nil
}

func (c *TriggerClient) Clone(endpoint string) *TriggerClient {
	return &TriggerClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
