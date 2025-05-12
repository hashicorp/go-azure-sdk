package v2025_04_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/cosmosdb"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/datatransfer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/graphapicompute"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/managedcassandras"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/materializedviewsbuilder"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/mongorbacs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/notebookworkspacesresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/rbacs"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/restorables"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/services"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2025-04-15/sqldedicatedgateway"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	cosmosDBClient, err := cosmosdb.NewCosmosDBClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CosmosDB client: %+v", err)
	}
	configureFunc(cosmosDBClient.Client)

	dataTransferClient, err := datatransfer.NewDataTransferClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataTransfer client: %+v", err)
	}
	configureFunc(dataTransferClient.Client)

	graphAPIComputeClient, err := graphapicompute.NewGraphAPIComputeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GraphAPICompute client: %+v", err)
	}
	configureFunc(graphAPIComputeClient.Client)

	managedCassandrasClient, err := managedcassandras.NewManagedCassandrasClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ManagedCassandras client: %+v", err)
	}
	configureFunc(managedCassandrasClient.Client)

	materializedViewsBuilderClient, err := materializedviewsbuilder.NewMaterializedViewsBuilderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaterializedViewsBuilder client: %+v", err)
	}
	configureFunc(materializedViewsBuilderClient.Client)

	mongorbacsClient, err := mongorbacs.NewMongorbacsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Mongorbacs client: %+v", err)
	}
	configureFunc(mongorbacsClient.Client)

	notebookWorkspacesResourceClient, err := notebookworkspacesresource.NewNotebookWorkspacesResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotebookWorkspacesResource client: %+v", err)
	}
	configureFunc(notebookWorkspacesResourceClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	rbacsClient, err := rbacs.NewRbacsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Rbacs client: %+v", err)
	}
	configureFunc(rbacsClient.Client)

	restorablesClient, err := restorables.NewRestorablesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Restorables client: %+v", err)
	}
	configureFunc(restorablesClient.Client)

	servicesClient, err := services.NewServicesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Services client: %+v", err)
	}
	configureFunc(servicesClient.Client)

	sqlDedicatedGatewayClient, err := sqldedicatedgateway.NewSqlDedicatedGatewayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlDedicatedGateway client: %+v", err)
	}
	configureFunc(sqlDedicatedGatewayClient.Client)

	return &Client{
		CosmosDB:                   cosmosDBClient,
		DataTransfer:               dataTransferClient,
		GraphAPICompute:            graphAPIComputeClient,
		ManagedCassandras:          managedCassandrasClient,
		MaterializedViewsBuilder:   materializedViewsBuilderClient,
		Mongorbacs:                 mongorbacsClient,
		NotebookWorkspacesResource: notebookWorkspacesResourceClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		Rbacs:                      rbacsClient,
		Restorables:                restorablesClient,
		Services:                   servicesClient,
		SqlDedicatedGateway:        sqlDedicatedGatewayClient,
	}, nil
}
