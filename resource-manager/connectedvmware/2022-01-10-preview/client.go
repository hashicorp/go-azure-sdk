package v2022_01_10_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/guestagents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/hosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/hybrididentitymetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/inventoryitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/machineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/resourcepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/vcenters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/virtualmachinetemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2022-01-10-preview/virtualnetworks"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Clusters                *clusters.ClustersClient
	DataStores              *datastores.DataStoresClient
	GuestAgents             *guestagents.GuestAgentsClient
	Hosts                   *hosts.HostsClient
	HybridIdentityMetadata  *hybrididentitymetadata.HybridIdentityMetadataClient
	InventoryItems          *inventoryitems.InventoryItemsClient
	MachineExtensions       *machineextensions.MachineExtensionsClient
	ResourcePools           *resourcepools.ResourcePoolsClient
	VCenters                *vcenters.VCentersClient
	VirtualMachineTemplates *virtualmachinetemplates.VirtualMachineTemplatesClient
	VirtualMachines         *virtualmachines.VirtualMachinesClient
	VirtualNetworks         *virtualnetworks.VirtualNetworksClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	clustersClient, err := clusters.NewClustersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Clusters client: %+v", err)
	}
	configureFunc(clustersClient.Client)

	dataStoresClient, err := datastores.NewDataStoresClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DataStores client: %+v", err)
	}
	configureFunc(dataStoresClient.Client)

	guestAgentsClient, err := guestagents.NewGuestAgentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GuestAgents client: %+v", err)
	}
	configureFunc(guestAgentsClient.Client)

	hostsClient, err := hosts.NewHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Hosts client: %+v", err)
	}
	configureFunc(hostsClient.Client)

	hybridIdentityMetadataClient, err := hybrididentitymetadata.NewHybridIdentityMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridIdentityMetadata client: %+v", err)
	}
	configureFunc(hybridIdentityMetadataClient.Client)

	inventoryItemsClient, err := inventoryitems.NewInventoryItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InventoryItems client: %+v", err)
	}
	configureFunc(inventoryItemsClient.Client)

	machineExtensionsClient, err := machineextensions.NewMachineExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MachineExtensions client: %+v", err)
	}
	configureFunc(machineExtensionsClient.Client)

	resourcePoolsClient, err := resourcepools.NewResourcePoolsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ResourcePools client: %+v", err)
	}
	configureFunc(resourcePoolsClient.Client)

	vCentersClient, err := vcenters.NewVCentersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VCenters client: %+v", err)
	}
	configureFunc(vCentersClient.Client)

	virtualMachineTemplatesClient, err := virtualmachinetemplates.NewVirtualMachineTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineTemplates client: %+v", err)
	}
	configureFunc(virtualMachineTemplatesClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	virtualNetworksClient, err := virtualnetworks.NewVirtualNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworks client: %+v", err)
	}
	configureFunc(virtualNetworksClient.Client)

	return &Client{
		Clusters:                clustersClient,
		DataStores:              dataStoresClient,
		GuestAgents:             guestAgentsClient,
		Hosts:                   hostsClient,
		HybridIdentityMetadata:  hybridIdentityMetadataClient,
		InventoryItems:          inventoryItemsClient,
		MachineExtensions:       machineExtensionsClient,
		ResourcePools:           resourcePoolsClient,
		VCenters:                vCentersClient,
		VirtualMachineTemplates: virtualMachineTemplatesClient,
		VirtualMachines:         virtualMachinesClient,
		VirtualNetworks:         virtualNetworksClient,
	}, nil
}
