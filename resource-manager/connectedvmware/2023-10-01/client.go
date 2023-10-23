package v2023_10_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/clusters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/datastores"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/hosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/inventoryitems"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/resourcepools"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vcenters"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualmachineinstances"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualmachinetemplates"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/virtualnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vminstanceguestagents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/connectedvmware/2023-10-01/vminstancehybrididentitymetadata"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	Clusters                         *clusters.ClustersClient
	DataStores                       *datastores.DataStoresClient
	Hosts                            *hosts.HostsClient
	InventoryItems                   *inventoryitems.InventoryItemsClient
	ResourcePools                    *resourcepools.ResourcePoolsClient
	VCenters                         *vcenters.VCentersClient
	VMInstanceGuestAgents            *vminstanceguestagents.VMInstanceGuestAgentsClient
	VMInstanceHybridIdentityMetadata *vminstancehybrididentitymetadata.VMInstanceHybridIdentityMetadataClient
	VirtualMachineInstances          *virtualmachineinstances.VirtualMachineInstancesClient
	VirtualMachineTemplates          *virtualmachinetemplates.VirtualMachineTemplatesClient
	VirtualNetworks                  *virtualnetworks.VirtualNetworksClient
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

	hostsClient, err := hosts.NewHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Hosts client: %+v", err)
	}
	configureFunc(hostsClient.Client)

	inventoryItemsClient, err := inventoryitems.NewInventoryItemsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building InventoryItems client: %+v", err)
	}
	configureFunc(inventoryItemsClient.Client)

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

	vMInstanceGuestAgentsClient, err := vminstanceguestagents.NewVMInstanceGuestAgentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMInstanceGuestAgents client: %+v", err)
	}
	configureFunc(vMInstanceGuestAgentsClient.Client)

	vMInstanceHybridIdentityMetadataClient, err := vminstancehybrididentitymetadata.NewVMInstanceHybridIdentityMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VMInstanceHybridIdentityMetadata client: %+v", err)
	}
	configureFunc(vMInstanceHybridIdentityMetadataClient.Client)

	virtualMachineInstancesClient, err := virtualmachineinstances.NewVirtualMachineInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineInstances client: %+v", err)
	}
	configureFunc(virtualMachineInstancesClient.Client)

	virtualMachineTemplatesClient, err := virtualmachinetemplates.NewVirtualMachineTemplatesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineTemplates client: %+v", err)
	}
	configureFunc(virtualMachineTemplatesClient.Client)

	virtualNetworksClient, err := virtualnetworks.NewVirtualNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualNetworks client: %+v", err)
	}
	configureFunc(virtualNetworksClient.Client)

	return &Client{
		Clusters:                         clustersClient,
		DataStores:                       dataStoresClient,
		Hosts:                            hostsClient,
		InventoryItems:                   inventoryItemsClient,
		ResourcePools:                    resourcePoolsClient,
		VCenters:                         vCentersClient,
		VMInstanceGuestAgents:            vMInstanceGuestAgentsClient,
		VMInstanceHybridIdentityMetadata: vMInstanceHybridIdentityMetadataClient,
		VirtualMachineInstances:          virtualMachineInstancesClient,
		VirtualMachineTemplates:          virtualMachineTemplatesClient,
		VirtualNetworks:                  virtualNetworksClient,
	}, nil
}
