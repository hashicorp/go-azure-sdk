package v2026_03_15

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/datatransfer"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/fleets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/graphapicompute"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/materializedviewsbuilder"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/notebookworkspacesresource"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/openapis"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/cosmosdb/2026-03-15/sqldedicatedgateway"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	DataTransfer               *datatransfer.DataTransferClient
	Fleets                     *fleets.FleetsClient
	GraphAPICompute            *graphapicompute.GraphAPIComputeClient
	MaterializedViewsBuilder   *materializedviewsbuilder.MaterializedViewsBuilderClient
	NotebookWorkspacesResource *notebookworkspacesresource.NotebookWorkspacesResourceClient
	Openapis                   *openapis.OpenapisClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	SqlDedicatedGateway        *sqldedicatedgateway.SqlDedicatedGatewayClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	dataTransferClient, err := datatransfer.NewDataTransferClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataTransfer client: %+v", err)
	}
	configureFunc(dataTransferClient.Client)

	fleetsClient, err := fleets.NewFleetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Fleets client: %+v", err)
	}
	configureFunc(fleetsClient.Client)

	graphAPIComputeClient, err := graphapicompute.NewGraphAPIComputeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GraphAPICompute client: %+v", err)
	}
	configureFunc(graphAPIComputeClient.Client)

	materializedViewsBuilderClient, err := materializedviewsbuilder.NewMaterializedViewsBuilderClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MaterializedViewsBuilder client: %+v", err)
	}
	configureFunc(materializedViewsBuilderClient.Client)

	notebookWorkspacesResourceClient, err := notebookworkspacesresource.NewNotebookWorkspacesResourceClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NotebookWorkspacesResource client: %+v", err)
	}
	configureFunc(notebookWorkspacesResourceClient.Client)

	openapisClient, err := openapis.NewOpenapisClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Openapis client: %+v", err)
	}
	configureFunc(openapisClient.Client)

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

	sqlDedicatedGatewayClient, err := sqldedicatedgateway.NewSqlDedicatedGatewayClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SqlDedicatedGateway client: %+v", err)
	}
	configureFunc(sqlDedicatedGatewayClient.Client)

	return &Client{
		DataTransfer:               dataTransferClient,
		Fleets:                     fleetsClient,
		GraphAPICompute:            graphAPIComputeClient,
		MaterializedViewsBuilder:   materializedViewsBuilderClient,
		NotebookWorkspacesResource: notebookWorkspacesResourceClient,
		Openapis:                   openapisClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		SqlDedicatedGateway:        sqlDedicatedGatewayClient,
	}, nil
}
