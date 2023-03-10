package v2022_04_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/accessconnector"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/put"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/vnetpeering"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/workspaces"
)

type Client struct {
	AccessConnector                      *accessconnector.AccessConnectorClient
	DELETE                               *delete.DELETEClient
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PUT                                  *put.PUTClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	VNetPeering                          *vnetpeering.VNetPeeringClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	accessConnectorClient := accessconnector.NewAccessConnectorClientWithBaseURI(endpoint)
	configureAuthFunc(&accessConnectorClient.Client)

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
		AccessConnector:                      &accessConnectorClient,
		DELETE:                               &dELETEClient,
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PUT:                                  &pUTClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		PrivateLinkResources:                 &privateLinkResourcesClient,
		VNetPeering:                          &vNetPeeringClient,
		Workspaces:                           &workspacesClient,
	}
}
