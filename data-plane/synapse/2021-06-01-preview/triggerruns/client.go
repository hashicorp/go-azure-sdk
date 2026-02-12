package triggerruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TriggerrunsClient struct {
	Client *dataplane.Client
}

func NewTriggerrunsClientUnconfigured() (*TriggerrunsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "triggerruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggerrunsClient: %+v", err)
	}

	return &TriggerrunsClient{
		Client: client,
	}, nil
}

func (c *TriggerrunsClient) TriggerrunsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewTriggerrunsClientWithBaseURI(endpoint string) (*TriggerrunsClient, error) {
	client, err := dataplane.NewClient(endpoint, "triggerruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating TriggerrunsClient: %+v", err)
	}

	return &TriggerrunsClient{
		Client: client,
	}, nil
}

func (c *TriggerrunsClient) Clone(endpoint string) *TriggerrunsClient {
	return &TriggerrunsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
