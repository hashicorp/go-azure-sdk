package v2021_10_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/cosmosdb"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/managedcassandras"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/notebookworkspacesresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/rbacs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2021-10-15/restorables"
)

type Client struct {
	CosmosDB                   *cosmosdb.CosmosDBClient
	ManagedCassandras          *managedcassandras.ManagedCassandrasClient
	NotebookWorkspacesResource *notebookworkspacesresource.NotebookWorkspacesResourceClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Rbacs                      *rbacs.RbacsClient
	Restorables                *restorables.RestorablesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	cosmosDBClient := cosmosdb.NewCosmosDBClientWithBaseURI(endpoint)
	configureAuthFunc(&cosmosDBClient.Client)

	managedCassandrasClient := managedcassandras.NewManagedCassandrasClientWithBaseURI(endpoint)
	configureAuthFunc(&managedCassandrasClient.Client)

	notebookWorkspacesResourceClient := notebookworkspacesresource.NewNotebookWorkspacesResourceClientWithBaseURI(endpoint)
	configureAuthFunc(&notebookWorkspacesResourceClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	rbacsClient := rbacs.NewRbacsClientWithBaseURI(endpoint)
	configureAuthFunc(&rbacsClient.Client)

	restorablesClient := restorables.NewRestorablesClientWithBaseURI(endpoint)
	configureAuthFunc(&restorablesClient.Client)

	return Client{
		CosmosDB:                   &cosmosDBClient,
		ManagedCassandras:          &managedCassandrasClient,
		NotebookWorkspacesResource: &notebookWorkspacesResourceClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		Rbacs:                      &rbacsClient,
		Restorables:                &restorablesClient,
	}
}
