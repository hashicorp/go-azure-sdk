package v2023_04_03

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-04-03/azuremonitorworkspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AzureMonitorWorkspaces *azuremonitorworkspaces.AzureMonitorWorkspacesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	azureMonitorWorkspacesClient, err := azuremonitorworkspaces.NewAzureMonitorWorkspacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AzureMonitorWorkspaces client: %+v", err)
	}
	configureFunc(azureMonitorWorkspacesClient.Client)

	return &Client{
		AzureMonitorWorkspaces: azureMonitorWorkspacesClient,
	}, nil
}
