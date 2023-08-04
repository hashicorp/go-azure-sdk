package v2022_12_27

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/extensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machineextensionsupgrade"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machinenetworkprofile"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/machines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/privateendpointconnections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/privatelinkresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/hybridcompute/2022-12-27/privatelinkscopes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
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

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	extensionsClient, err := extensions.NewExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Extensions client: %+v", err)
	}
	configureFunc(extensionsClient.Client)

	machineExtensionsClient, err := machineextensions.NewMachineExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineExtensions client: %+v", err)
	}
	configureFunc(machineExtensionsClient.Client)

	machineExtensionsUpgradeClient, err := machineextensionsupgrade.NewMachineExtensionsUpgradeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineExtensionsUpgrade client: %+v", err)
	}
	configureFunc(machineExtensionsUpgradeClient.Client)

	machineNetworkProfileClient, err := machinenetworkprofile.NewMachineNetworkProfileClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineNetworkProfile client: %+v", err)
	}
	configureFunc(machineNetworkProfileClient.Client)

	machinesClient, err := machines.NewMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Machines client: %+v", err)
	}
	configureFunc(machinesClient.Client)

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

	privateLinkScopesClient, err := privatelinkscopes.NewPrivateLinkScopesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building PrivateLinkScopes client: %+v", err)
	}
	configureFunc(privateLinkScopesClient.Client)

	return &Client{
		Extensions:                 extensionsClient,
		MachineExtensions:          machineExtensionsClient,
		MachineExtensionsUpgrade:   machineExtensionsUpgradeClient,
		MachineNetworkProfile:      machineNetworkProfileClient,
		Machines:                   machinesClient,
		PrivateEndpointConnections: privateEndpointConnectionsClient,
		PrivateLinkResources:       privateLinkResourcesClient,
		PrivateLinkScopes:          privateLinkScopesClient,
	}, nil
}
