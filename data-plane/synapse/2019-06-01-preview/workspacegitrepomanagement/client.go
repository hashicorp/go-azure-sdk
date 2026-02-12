package workspacegitrepomanagement

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/client/dataplane"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkspaceGitRepoManagementClient struct {
	Client *dataplane.Client
}

func NewWorkspaceGitRepoManagementClientUnconfigured() (*WorkspaceGitRepoManagementClient, error) {
	client, err := dataplane.NewClient("please_configure_client_endpoint", "workspacegitrepomanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceGitRepoManagementClient: %+v", err)
	}

	return &WorkspaceGitRepoManagementClient{
		Client: client,
	}, nil
}

func (c *WorkspaceGitRepoManagementClient) WorkspaceGitRepoManagementClientSetEndpoint(endpoint string) {
	c.Client.Client.BaseUri = endpoint
}

func NewWorkspaceGitRepoManagementClientWithBaseURI(endpoint string) (*WorkspaceGitRepoManagementClient, error) {
	client, err := dataplane.NewClient(endpoint, "workspacegitrepomanagement", defaultApiVersion)
	if err != nil {
		return nil, fmt.Errorf("instantiating WorkspaceGitRepoManagementClient: %+v", err)
	}

	return &WorkspaceGitRepoManagementClient{
		Client: client,
	}, nil
}

func (c *WorkspaceGitRepoManagementClient) Clone(endpoint string) *WorkspaceGitRepoManagementClient {
	return &WorkspaceGitRepoManagementClient{
		Client: c.Client.CloneClient(endpoint),
	}
}
