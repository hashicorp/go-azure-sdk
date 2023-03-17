package v2022_04_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/accessconnector"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/delete"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/put"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/vnetpeering"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2022-04-01-preview/workspaces"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	"github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(api environments.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	accessConnectorClient, err := accessconnector.NewAccessConnectorClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building AccessConnector client: %+v", err)
	}
	configureFunc(accessConnectorClient.Client)

	dELETEClient, err := delete.NewDELETEClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building DELETE client: %+v", err)
	}
	configureFunc(dELETEClient.Client)

	outboundNetworkDependenciesEndpointsClient, err := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building OutboundNetworkDependenciesEndpoints client: %+v", err)
	}
	configureFunc(outboundNetworkDependenciesEndpointsClient.Client)

	pUTClient, err := put.NewPUTClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PUT client: %+v", err)
	}
	configureFunc(pUTClient.Client)

	privateEndpointConnectionsClient, err := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateEndpointConnections client: %+v", err)
	}
	configureFunc(privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient, err := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkResources client: %+v", err)
	}
	configureFunc(privateLinkResourcesClient.Client)

	vNetPeeringClient, err := vnetpeering.NewVNetPeeringClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building VNetPeering client: %+v", err)
	}
	configureFunc(vNetPeeringClient.Client)

	workspacesClient, err := workspaces.NewWorkspacesClientWithBaseURI(api)
	if err != nil {
		return nil, fmt.Errorf("building Workspaces client: %+v", err)
	}
	configureFunc(workspacesClient.Client)

	return &Client{
		AccessConnector:                      accessConnectorClient,
		DELETE:                               dELETEClient,
		OutboundNetworkDependenciesEndpoints: outboundNetworkDependenciesEndpointsClient,
		PUT:                                  pUTClient,
		PrivateEndpointConnections:           privateEndpointConnectionsClient,
		PrivateLinkResources:                 privateLinkResourcesClient,
		VNetPeering:                          vNetPeeringClient,
		Workspaces:                           workspacesClient,
	}, nil
}
