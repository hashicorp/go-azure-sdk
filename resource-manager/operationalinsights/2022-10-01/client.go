package v2022_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2022-10-01/deletedworkspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2022-10-01/tables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2022-10-01/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DeletedWorkspaces *deletedworkspaces.DeletedWorkspacesClient
	Tables            *tables.TablesClient
	Workspaces        *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	deletedWorkspacesClient, err := deletedworkspaces.NewDeletedWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedWorkspaces client: %+v", err)
	}
	configureFunc(deletedWorkspacesClient.Client)

	tablesClient, err := tables.NewTablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Tables client: %+v", err)
	}
	configureFunc(tablesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		DeletedWorkspaces: deletedWorkspacesClient,
		Tables:            tablesClient,
		Workspaces:        workspacesClient,
	}, nil
}
