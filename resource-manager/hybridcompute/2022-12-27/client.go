package v2022_12_27

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machineextensionsupgrade"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machinenetworkprofile"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/privatelinkscopes"
)

type Client struct {
	Extensions                 *extensions.ExtensionsClient
	MachineExtensions          *machineextensions.MachineExtensionsClient
	MachineExtensionsUpgrade   *machineextensionsupgrade.MachineExtensionsUpgradeClient
	MachineNetworkProfile      *machinenetworkprofile.MachineNetworkProfileClient
	Machines                   *machines.MachinesClient
	PrivateEndpointConnections *privateendpointconnections.PrivateEndpointConnectionsClient
	PrivateLinkResources       *privatelinkresources.PrivateLinkResourcesClient
	PrivateLinkScopes          *privatelinkscopes.PrivateLinkScopesClient
}

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	extensionsClient := extensions.NewExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&extensionsClient.Client)

	machineExtensionsClient := machineextensions.NewMachineExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&machineExtensionsClient.Client)

	machineExtensionsUpgradeClient := machineextensionsupgrade.NewMachineExtensionsUpgradeClientWithBaseURI(endpoint)
	configureAuthFunc(&machineExtensionsUpgradeClient.Client)

	machineNetworkProfileClient := machinenetworkprofile.NewMachineNetworkProfileClientWithBaseURI(endpoint)
	configureAuthFunc(&machineNetworkProfileClient.Client)

	machinesClient := machines.NewMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&machinesClient.Client)

	privateEndpointConnectionsClient := privateendpointconnections.NewPrivateEndpointConnectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&privateEndpointConnectionsClient.Client)

	privateLinkResourcesClient := privatelinkresources.NewPrivateLinkResourcesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkResourcesClient.Client)

	privateLinkScopesClient := privatelinkscopes.NewPrivateLinkScopesClientWithBaseURI(endpoint)
	configureAuthFunc(&privateLinkScopesClient.Client)

	return Client{
		Extensions:                 &extensionsClient,
		MachineExtensions:          &machineExtensionsClient,
		MachineExtensionsUpgrade:   &machineExtensionsUpgradeClient,
		MachineNetworkProfile:      &machineNetworkProfileClient,
		Machines:                   &machinesClient,
		PrivateEndpointConnections: &privateEndpointConnectionsClient,
		PrivateLinkResources:       &privateLinkResourcesClient,
		PrivateLinkScopes:          &privateLinkScopesClient,
	}
}
