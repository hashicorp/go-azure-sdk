package v2023_02_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/put"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/vnetpeering"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2023-02-01/workspaces"
)

type Client struct {
	DELETE                               *delete.DELETEClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PUT                                  *put.PUTClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	VNetPeering                          *vnetpeering.VNetPeeringClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	dELETEClient := delete.NewDELETEClientWithBaseURI(endpoint)
	configureAuthFunc(&dELETEClient.Client)

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	pUTClient := put.NewPUTClientWithBaseURI(endpoint)
	configureAuthFunc(&pUTClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	vNetPeeringClient := vnetpeering.NewVNetPeeringClientWithBaseURI(endpoint)
	configureAuthFunc(&vNetPeeringClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		DELETE:                               &dELETEClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PUT:                                  &pUTClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		PrivateLinkResources:                 &privateLinkResourcesClient,
		VNetPeering:                          &vNetPeeringClient,
		Workspaces:                           &workspacesClient,
	}
}
