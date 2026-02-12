package activityruns

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityrunsClient struct {
	Client *dataplane.Client
}

func NewActivityrunsClientUnconfigured() (*ActivityrunsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "activityruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityrunsClient: %+v", err)
	}

	return &ActivityrunsClient{
		Client: client,
	}, nil
}

func (c *ActivityrunsClient) ActivityrunsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewActivityrunsClientWithBaseURI(endpoint string) (*ActivityrunsClient, error) {
	client, err := dataplane.NewClient(endpoint, "activityruns", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating ActivityrunsClient: %+v", err)
	}

	return &ActivityrunsClient{
		Client: client,
	}, nil
}

func (c *ActivityrunsClient) Clone(endpoint string) *ActivityrunsClient {
	return &ActivityrunsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
