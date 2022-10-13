package v2022_01_10_preview

import (
	"github.com/Azure/go-autorest/autorest"
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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	clustersClient := clusters.NewClustersClientWithBaseURI(endpoint)
	configureAuthFunc(&clustersClient.Client)

	dataStoresClient := datastores.NewDataStoresClientWithBaseURI(endpoint)
	configureAuthFunc(&dataStoresClient.Client)

	guestAgentsClient := guestagents.NewGuestAgentsClientWithBaseURI(endpoint)
	configureAuthFunc(&guestAgentsClient.Client)

	hostsClient := hosts.NewHostsClientWithBaseURI(endpoint)
	configureAuthFunc(&hostsClient.Client)

	hybridIdentityMetadataClient := hybrididentitymetadata.NewHybridIdentityMetadataClientWithBaseURI(endpoint)
	configureAuthFunc(&hybridIdentityMetadataClient.Client)

	inventoryItemsClient := inventoryitems.NewInventoryItemsClientWithBaseURI(endpoint)
	configureAuthFunc(&inventoryItemsClient.Client)

	machineExtensionsClient := machineextensions.NewMachineExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&machineExtensionsClient.Client)

	resourcePoolsClient := resourcepools.NewResourcePoolsClientWithBaseURI(endpoint)
	configureAuthFunc(&resourcePoolsClient.Client)

	vCentersClient := vcenters.NewVCentersClientWithBaseURI(endpoint)
	configureAuthFunc(&vCentersClient.Client)

	virtualMachineTemplatesClient := virtualmachinetemplates.NewVirtualMachineTemplatesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineTemplatesClient.Client)

	virtualMachinesClient := virtualmachines.NewVirtualMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachinesClient.Client)

	virtualNetworksClient := virtualnetworks.NewVirtualNetworksClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualNetworksClient.Client)

	return Client{
		Clusters:                &clustersClient,
		DataStores:              &dataStoresClient,
		GuestAgents:             &guestAgentsClient,
		Hosts:                   &hostsClient,
		HybridIdentityMetadata:  &hybridIdentityMetadataClient,
		InventoryItems:          &inventoryItemsClient,
		MachineExtensions:       &machineExtensionsClient,
		ResourcePools:           &resourcePoolsClient,
		VCenters:                &vCentersClient,
		VirtualMachineTemplates: &virtualMachineTemplatesClient,
		VirtualMachines:         &virtualMachinesClient,
		VirtualNetworks:         &virtualNetworksClient,
	}
}
