package v2022_03_10

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/machineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/machineextensionsupgrade"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-03-10/privatelinkscopes"
)

type Client struct {
	MachineExtensions          *machineextensions.MachineExtensionsClient
	MachineExtensionsUpgrade   *machineextensionsupgrade.MachineExtensionsUpgradeClient
	Machines                   *machines.MachinesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	PrivateLinkScopes          *privatelinkscopes.PrivateLinkScopesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	machineExtensionsClient := machineextensions.NewMachineExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&machineExtensionsClient.Client)

	machineExtensionsUpgradeClient := machineextensionsupgrade.NewMachineExtensionsUpgradeClientWithBaseURI(endpoint)
	configureAuthFunc(&machineExtensionsUpgradeClient.Client)

	machinesClient := machines.NewMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&machinesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	privateLinkScopesClient := privatelinkscopes.NewPrivateLinkScopesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkScopesClient.Client)

	return Client{
		MachineExtensions:          &machineExtensionsClient,
		MachineExtensionsUpgrade:   &machineExtensionsUpgradeClient,
		Machines:                   &machinesClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		PrivateLinkScopes:          &privateLinkScopesClient,
	}
}
