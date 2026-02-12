package checkprincipalaccess

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CheckPrincipalAccessClient struct {
	Client *dataplane.Client
}

func NewCheckPrincipalAccessClientUnconfigured() (*CheckPrincipalAccessClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "checkprincipalaccess", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CheckPrincipalAccessClient: %+v", err)
	}

	return &CheckPrincipalAccessClient{
		Client: client,
	}, nil
}

func (c *CheckPrincipalAccessClient) CheckPrincipalAccessClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewCheckPrincipalAccessClientWithBaseURI(endpoint string) (*CheckPrincipalAccessClient, error) {
	client, err := dataplane.NewClient(endpoint, "checkprincipalaccess", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating CheckPrincipalAccessClient: %+v", err)
	}

	return &CheckPrincipalAccessClient{
		Client: client,
	}, nil
}

func (c *CheckPrincipalAccessClient) Clone(endpoint string) *CheckPrincipalAccessClient {
	return &CheckPrincipalAccessClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
