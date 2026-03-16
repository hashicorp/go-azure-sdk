package administrations

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdministrationsClient struct {
	Client *dataplane.Client
}

func NewAdministrationsClientUnconfigured() (*AdministrationsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "administrations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrationsClient: %+v", err)
	}

	return &AdministrationsClient{
		Client: client,
	}, nil
}

func (c *AdministrationsClient) AdministrationsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewAdministrationsClientWithBaseURI(endpoint string) (*AdministrationsClient, error) {
	client, err := dataplane.NewClient(endpoint, "administrations", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating AdministrationsClient: %+v", err)
	}

	return &AdministrationsClient{
		Client: client,
	}, nil
}

func (c *AdministrationsClient) Clone(endpoint string) *AdministrationsClient {
	return &AdministrationsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
