package sqlscripts

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlScriptsClient struct {
	Client *dataplane.Client
}

func NewSqlScriptsClientUnconfigured() (*SqlScriptsClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "sqlscripts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlScriptsClient: %+v", err)
	}

	return &SqlScriptsClient{
		Client: client,
	}, nil
}

func (c *SqlScriptsClient) SqlScriptsClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewSqlScriptsClientWithBaseURI(endpoint string) (*SqlScriptsClient, error) {
	client, err := dataplane.NewClient(endpoint, "sqlscripts", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating SqlScriptsClient: %+v", err)
	}

	return &SqlScriptsClient{
		Client: client,
	}, nil
}

func (c *SqlScriptsClient) Clone(endpoint string) *SqlScriptsClient {
	return &SqlScriptsClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
