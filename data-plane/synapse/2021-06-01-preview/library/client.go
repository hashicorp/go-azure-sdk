package library

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LibraryClient struct {
	Client *dataplane.Client
}

func NewLibraryClientUnconfigured() (*LibraryClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "library", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LibraryClient: %+v", err)
	}

	return &LibraryClient{
		Client: client,
	}, nil
}

func (c *LibraryClient) LibraryClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewLibraryClientWithBaseURI(endpoint string) (*LibraryClient, error) {
	client, err := dataplane.NewClient(endpoint, "library", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating LibraryClient: %+v", err)
	}

	return &LibraryClient{
		Client: client,
	}, nil
}

func (c *LibraryClient) Clone(endpoint string) *LibraryClient {
	return &LibraryClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
