package v2021_04_01_preview

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/outboundnetworkdependenciesendpoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/vnetpeering"
	"github.com/hashicorp/go-azure-sdk/resource-manager/databricks/2021-04-01-preview/workspaces"
)

type Client struct {
	OutboundNetworkDependenciesEndpoints *outboundnetworkdependenciesendpoints.OutboundNetworkDependenciesEndpointsClient
	PrivateEndpointConnections           *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources                 *privatelinkresources.PrivateLinkResourcesClient
	VNetPeering                          *vnetpeering.VNetPeeringClient
	Workspaces                           *workspaces.WorkspacesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	outboundNetworkDependenciesEndpointsClient := outboundnetworkdependenciesendpoints.NewOutboundNetworkDependenciesEndpointsClientWithBaseURI(endpoint)
	configureAuthFunc(&outboundNetworkDependenciesEndpointsClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	vNetPeeringClient := vnetpeering.NewVNetPeeringClientWithBaseURI(endpoint)
	configureAuthFunc(&vNetPeeringClient.Client)

	workspacesClient := workspaces.NewWorkspacesClientWithBaseURI(endpoint)
	configureAuthFunc(&workspacesClient.Client)

	return Client{
		OutboundNetworkDependenciesEndpoints: &outboundNetworkDependenciesEndpointsClient,
		PrivateEndpointConnections:           &privateEndpointConnectionsClient,
		PrivateLinkResources:                 &privateLinkResourcesClient,
		VNetPeering:                          &vNetPeeringClient,
		Workspaces:                           &workspacesClient,
	}
}
