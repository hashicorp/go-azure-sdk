package workspace

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceClient struct {
	Client *dataplane.Client
}

func NewWorkspaceClientUnconfigured() (*WorkspaceClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "workspace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceClient: %+v", err)
	}

	return &WorkspaceClient{
		Client: client,
	}, nil
}

func (c *WorkspaceClient) WorkspaceClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewWorkspaceClientWithBaseURI(endpoint string) (*WorkspaceClient, error) {
	client, err := dataplane.NewClient(endpoint, "workspace", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceClient: %+v", err)
	}

	return &WorkspaceClient{
		Client: client,
	}, nil
}

func (c *WorkspaceClient) Clone(endpoint string) *WorkspaceClient {
	return &WorkspaceClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
