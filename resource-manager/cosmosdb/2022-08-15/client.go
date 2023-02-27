package v2022_08_15

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/cosmosdb"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/datatransfer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/graphapicompute"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/managedcassandras"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/materializedviewsbuilder"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/mongorbacs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/notebookworkspacesresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/rbacs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/restorables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2022-08-15/sqldedicatedgateway"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Client struct {
	CosmosDB                   *cosmosdb.CosmosDBClient
	DataTransfer               *datatransfer.DataTransferClient
	GraphAPICompute            *graphapicompute.GraphAPIComputeClient
	ManagedCassandras          *managedcassandras.ManagedCassandrasClient
	MaterializedViewsBuilder   *materializedviewsbuilder.MaterializedViewsBuilderClient
	Mongorbacs                 *mongorbacs.MongorbacsClient
	NotebookWorkspacesResource *notebookworkspacesresource.NotebookWorkspacesResourceClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	Rbacs                      *rbacs.RbacsClient
	Restorables                *restorables.RestorablesClient
	Services                   *services.ServicesClient
	SqlDedicatedGateway        *sqldedicatedgateway.SqlDedicatedGatewayClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	cosmosDBClient := cosmosdb.NewCosmosDBClientWithBaseURI(endpoint)
	configureAuthFunc(&cosmosDBClient.Client)

	dataTransferClient := datatransfer.NewDataTransferClientWithBaseURI(endpoint)
	configureAuthFunc(&dataTransferClient.Client)

	graphAPIComputeClient := graphapicompute.NewGraphAPIComputeClientWithBaseURI(endpoint)
	configureAuthFunc(&graphAPIComputeClient.Client)

	managedCassandrasClient := managedcassandras.NewManagedCassandrasClientWithBaseURI(endpoint)
	configureAuthFunc(&managedCassandrasClient.Client)

	materializedViewsBuilderClient := materializedviewsbuilder.NewMaterializedViewsBuilderClientWithBaseURI(endpoint)
	configureAuthFunc(&materializedViewsBuilderClient.Client)

	mongorbacsClient := mongorbacs.NewMongorbacsClientWithBaseURI(endpoint)
	configureAuthFunc(&mongorbacsClient.Client)

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

	servicesClient := services.NewServicesClientWithBaseURI(endpoint)
	configureAuthFunc(&servicesClient.Client)

	sqlDedicatedGatewayClient := sqldedicatedgateway.NewSqlDedicatedGatewayClientWithBaseURI(endpoint)
	configureAuthFunc(&sqlDedicatedGatewayClient.Client)

	return Client{
		CosmosDB:                   &cosmosDBClient,
		DataTransfer:               &dataTransferClient,
		GraphAPICompute:            &graphAPIComputeClient,
		ManagedCassandras:          &managedCassandrasClient,
		MaterializedViewsBuilder:   &materializedViewsBuilderClient,
		Mongorbacs:                 &mongorbacsClient,
		NotebookWorkspacesResource: &notebookWorkspacesResourceClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		Rbacs:                      &rbacsClient,
		Restorables:                &restorablesClient,
		Services:                   &servicesClient,
		SqlDedicatedGateway:        &sqlDedicatedGatewayClient,
	}
}
