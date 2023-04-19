package v2023_04_03

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/insights/2023-04-03/azuremonitorworkspaces"
)

type Client struct {
	AzureMonitorWorkspaces *azuremonitorworkspaces.AzureMonitorWorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	azureMonitorWorkspacesClient := azuremonitorworkspaces.NewAzureMonitorWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&azureMonitorWorkspacesClient.Client)

	return Client{
		AzureMonitorWorkspaces: &azureMonitorWorkspacesClient,
	}
}
