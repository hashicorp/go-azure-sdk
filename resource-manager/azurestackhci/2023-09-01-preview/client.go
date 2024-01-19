package v2023_09_01_preview

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/galleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/guestagents"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/hybrididentitymetadata"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/logicalnetworks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/marketplacegalleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/networkinterfaces"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/storagecontainers"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/virtualharddisks"
	"github.com/hashicorp/go-azure-sdk/resource-manager/azurestackhci/2023-09-01-preview/virtualmachineinstances"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	GalleryImages            *galleryimages.GalleryImagesClient
	GuestAgents              *guestagents.GuestAgentsClient
	HybridIdentityMetadata   *hybrididentitymetadata.HybridIdentityMetadataClient
	LogicalNetworks          *logicalnetworks.LogicalNetworksClient
	MarketplaceGalleryImages *marketplacegalleryimages.MarketplaceGalleryImagesClient
	NetworkInterfaces        *networkinterfaces.NetworkInterfacesClient
	StorageContainers        *storagecontainers.StorageContainersClient
	VirtualHardDisks         *virtualharddisks.VirtualHardDisksClient
	VirtualMachineInstances  *virtualmachineinstances.VirtualMachineInstancesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	galleryImagesClient, err := galleryimages.NewGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImages client: %+v", err)
	}
	configureFunc(galleryImagesClient.Client)

	guestAgentsClient, err := guestagents.NewGuestAgentsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GuestAgents client: %+v", err)
	}
	configureFunc(guestAgentsClient.Client)

	hybridIdentityMetadataClient, err := hybrididentitymetadata.NewHybridIdentityMetadataClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building HybridIdentityMetadata client: %+v", err)
	}
	configureFunc(hybridIdentityMetadataClient.Client)

	logicalNetworksClient, err := logicalnetworks.NewLogicalNetworksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogicalNetworks client: %+v", err)
	}
	configureFunc(logicalNetworksClient.Client)

	marketplaceGalleryImagesClient, err := marketplacegalleryimages.NewMarketplaceGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building MarketplaceGalleryImages client: %+v", err)
	}
	configureFunc(marketplaceGalleryImagesClient.Client)

	networkInterfacesClient, err := networkinterfaces.NewNetworkInterfacesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building NetworkInterfaces client: %+v", err)
	}
	configureFunc(networkInterfacesClient.Client)

	storageContainersClient, err := storagecontainers.NewStorageContainersClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building StorageContainers client: %+v", err)
	}
	configureFunc(storageContainersClient.Client)

	virtualHardDisksClient, err := virtualharddisks.NewVirtualHardDisksClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualHardDisks client: %+v", err)
	}
	configureFunc(virtualHardDisksClient.Client)

	virtualMachineInstancesClient, err := virtualmachineinstances.NewVirtualMachineInstancesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineInstances client: %+v", err)
	}
	configureFunc(virtualMachineInstancesClient.Client)

	return &Client{
		GalleryImages:            galleryImagesClient,
		GuestAgents:              guestAgentsClient,
		HybridIdentityMetadata:   hybridIdentityMetadataClient,
		LogicalNetworks:          logicalNetworksClient,
		MarketplaceGalleryImages: marketplaceGalleryImagesClient,
		NetworkInterfaces:        networkInterfacesClient,
		StorageContainers:        storageContainersClient,
		VirtualHardDisks:         virtualHardDisksClient,
		VirtualMachineInstances:  virtualMachineInstancesClient,
	}, nil
}
