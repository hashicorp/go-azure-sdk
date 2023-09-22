package v2021_06_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2021-06-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2021-06-01/deletedworkspaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/operationalinsights/2021-06-01/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Clusters          *clusters.ClustersClient
	DeletedWorkspaces *deletedworkspaces.DeletedWorkspacesClient
	Workspaces        *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	deletedWorkspacesClient, err := deletedworkspaces.NewDeletedWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DeletedWorkspaces client: %+v", err)
	}
	configureFunc(deletedWorkspacesClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		Clusters:          clustersClient,
		DeletedWorkspaces: deletedWorkspacesClient,
		Workspaces:        workspacesClient,
	}, nil
}
