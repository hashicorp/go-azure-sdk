package v2021_07_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/availabilitysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/capacityreservation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/capacityreservationgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/capacityreservations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/communitygalleryimageversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/compute"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/dedicatedhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/dedicatedhostgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/dedicatedhosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/galleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/galleryapplications"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/galleryapplicationversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/galleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/galleryimageversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/gallerysharingupdate"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/images"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/loganalytics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/proximityplacementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/restorepointcollections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleries"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleryimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sharedgalleryimageversions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/skus"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/sshpublickeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachineextensionimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachineimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachineruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinescalesetextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinescalesetrollingupgrades"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinescalesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinescalesetvmextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinescalesetvmruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinescalesetvms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-07-01/virtualmachinesizes"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailabilitySets                      *availabilitysets.AvailabilitySetsClient
	CapacityReservation                   *capacityreservation.CapacityReservationClient
	CapacityReservationGroups             *capacityreservationgroups.CapacityReservationGroupsClient
	CapacityReservations                  *capacityreservations.CapacityReservationsClient
	CommunityGalleries                    *communitygalleries.CommunityGalleriesClient
	CommunityGalleryImageVersions         *communitygalleryimageversions.CommunityGalleryImageVersionsClient
	CommunityGalleryImages                *communitygalleryimages.CommunityGalleryImagesClient
	Compute                               *compute.ComputeClient
	DedicatedHost                         *dedicatedhost.DedicatedHostClient
	DedicatedHostGroups                   *dedicatedhostgroups.DedicatedHostGroupsClient
	DedicatedHosts                        *dedicatedhosts.DedicatedHostsClient
	Galleries                             *galleries.GalleriesClient
	GalleryApplicationVersions            *galleryapplicationversions.GalleryApplicationVersionsClient
	GalleryApplications                   *galleryapplications.GalleryApplicationsClient
	GalleryImageVersions                  *galleryimageversions.GalleryImageVersionsClient
	GalleryImages                         *galleryimages.GalleryImagesClient
	GallerySharingUpdate                  *gallerysharingupdate.GallerySharingUpdateClient
	Images                                *images.ImagesClient
	LogAnalytics                          *loganalytics.LogAnalyticsClient
	ProximityPlacementGroups              *proximityplacementgroups.ProximityPlacementGroupsClient
	RestorePointCollections               *restorepointcollections.RestorePointCollectionsClient
	SharedGalleries                       *sharedgalleries.SharedGalleriesClient
	SharedGalleryImageVersions            *sharedgalleryimageversions.SharedGalleryImageVersionsClient
	SharedGalleryImages                   *sharedgalleryimages.SharedGalleryImagesClient
	Skus                                  *skus.SkusClient
	SshPublicKeys                         *sshpublickeys.SshPublicKeysClient
	VirtualMachineExtensionImages         *virtualmachineextensionimages.VirtualMachineExtensionImagesClient
	VirtualMachineExtensions              *virtualmachineextensions.VirtualMachineExtensionsClient
	VirtualMachineImages                  *virtualmachineimages.VirtualMachineImagesClient
	VirtualMachineRunCommands             *virtualmachineruncommands.VirtualMachineRunCommandsClient
	VirtualMachineScaleSetExtensions      *virtualmachinescalesetextensions.VirtualMachineScaleSetExtensionsClient
	VirtualMachineScaleSetRollingUpgrades *virtualmachinescalesetrollingupgrades.VirtualMachineScaleSetRollingUpgradesClient
	VirtualMachineScaleSetVMExtensions    *virtualmachinescalesetvmextensions.VirtualMachineScaleSetVMExtensionsClient
	VirtualMachineScaleSetVMRunCommands   *virtualmachinescalesetvmruncommands.VirtualMachineScaleSetVMRunCommandsClient
	VirtualMachineScaleSetVMs             *virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient
	VirtualMachineScaleSets               *virtualmachinescalesets.VirtualMachineScaleSetsClient
	VirtualMachineSizes                   *virtualmachinesizes.VirtualMachineSizesClient
	VirtualMachines                       *virtualmachines.VirtualMachinesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	availabilitySetsClient, err := availabilitysets.NewAvailabilitySetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailabilitySets client: %+v", err)
	}
	configureFunc(availabilitySetsClient.Client)

	capacityReservationClient, err := capacityreservation.NewCapacityReservationClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityReservation client: %+v", err)
	}
	configureFunc(capacityReservationClient.Client)

	capacityReservationGroupsClient, err := capacityreservationgroups.NewCapacityReservationGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityReservationGroups client: %+v", err)
	}
	configureFunc(capacityReservationGroupsClient.Client)

	capacityReservationsClient, err := capacityreservations.NewCapacityReservationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CapacityReservations client: %+v", err)
	}
	configureFunc(capacityReservationsClient.Client)

	communityGalleriesClient, err := communitygalleries.NewCommunityGalleriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunityGalleries client: %+v", err)
	}
	configureFunc(communityGalleriesClient.Client)

	communityGalleryImageVersionsClient, err := communitygalleryimageversions.NewCommunityGalleryImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunityGalleryImageVersions client: %+v", err)
	}
	configureFunc(communityGalleryImageVersionsClient.Client)

	communityGalleryImagesClient, err := communitygalleryimages.NewCommunityGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building CommunityGalleryImages client: %+v", err)
	}
	configureFunc(communityGalleryImagesClient.Client)

	computeClient, err := compute.NewComputeClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Compute client: %+v", err)
	}
	configureFunc(computeClient.Client)

	dedicatedHostClient, err := dedicatedhost.NewDedicatedHostClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHost client: %+v", err)
	}
	configureFunc(dedicatedHostClient.Client)

	dedicatedHostGroupsClient, err := dedicatedhostgroups.NewDedicatedHostGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHostGroups client: %+v", err)
	}
	configureFunc(dedicatedHostGroupsClient.Client)

	dedicatedHostsClient, err := dedicatedhosts.NewDedicatedHostsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building DedicatedHosts client: %+v", err)
	}
	configureFunc(dedicatedHostsClient.Client)

	galleriesClient, err := galleries.NewGalleriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Galleries client: %+v", err)
	}
	configureFunc(galleriesClient.Client)

	galleryApplicationVersionsClient, err := galleryapplicationversions.NewGalleryApplicationVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryApplicationVersions client: %+v", err)
	}
	configureFunc(galleryApplicationVersionsClient.Client)

	galleryApplicationsClient, err := galleryapplications.NewGalleryApplicationsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryApplications client: %+v", err)
	}
	configureFunc(galleryApplicationsClient.Client)

	galleryImageVersionsClient, err := galleryimageversions.NewGalleryImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImageVersions client: %+v", err)
	}
	configureFunc(galleryImageVersionsClient.Client)

	galleryImagesClient, err := galleryimages.NewGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GalleryImages client: %+v", err)
	}
	configureFunc(galleryImagesClient.Client)

	gallerySharingUpdateClient, err := gallerysharingupdate.NewGallerySharingUpdateClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building GallerySharingUpdate client: %+v", err)
	}
	configureFunc(gallerySharingUpdateClient.Client)

	imagesClient, err := images.NewImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Images client: %+v", err)
	}
	configureFunc(imagesClient.Client)

	logAnalyticsClient, err := loganalytics.NewLogAnalyticsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building LogAnalytics client: %+v", err)
	}
	configureFunc(logAnalyticsClient.Client)

	proximityPlacementGroupsClient, err := proximityplacementgroups.NewProximityPlacementGroupsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ProximityPlacementGroups client: %+v", err)
	}
	configureFunc(proximityPlacementGroupsClient.Client)

	restorePointCollectionsClient, err := restorepointcollections.NewRestorePointCollectionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorePointCollections client: %+v", err)
	}
	configureFunc(restorePointCollectionsClient.Client)

	sharedGalleriesClient, err := sharedgalleries.NewSharedGalleriesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedGalleries client: %+v", err)
	}
	configureFunc(sharedGalleriesClient.Client)

	sharedGalleryImageVersionsClient, err := sharedgalleryimageversions.NewSharedGalleryImageVersionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedGalleryImageVersions client: %+v", err)
	}
	configureFunc(sharedGalleryImageVersionsClient.Client)

	sharedGalleryImagesClient, err := sharedgalleryimages.NewSharedGalleryImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SharedGalleryImages client: %+v", err)
	}
	configureFunc(sharedGalleryImagesClient.Client)

	skusClient, err := skus.NewSkusClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Skus client: %+v", err)
	}
	configureFunc(skusClient.Client)

	sshPublicKeysClient, err := sshpublickeys.NewSshPublicKeysClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SshPublicKeys client: %+v", err)
	}
	configureFunc(sshPublicKeysClient.Client)

	virtualMachineExtensionImagesClient, err := virtualmachineextensionimages.NewVirtualMachineExtensionImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineExtensionImages client: %+v", err)
	}
	configureFunc(virtualMachineExtensionImagesClient.Client)

	virtualMachineExtensionsClient, err := virtualmachineextensions.NewVirtualMachineExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineExtensions client: %+v", err)
	}
	configureFunc(virtualMachineExtensionsClient.Client)

	virtualMachineImagesClient, err := virtualmachineimages.NewVirtualMachineImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineImages client: %+v", err)
	}
	configureFunc(virtualMachineImagesClient.Client)

	virtualMachineRunCommandsClient, err := virtualmachineruncommands.NewVirtualMachineRunCommandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineRunCommands client: %+v", err)
	}
	configureFunc(virtualMachineRunCommandsClient.Client)

	virtualMachineScaleSetExtensionsClient, err := virtualmachinescalesetextensions.NewVirtualMachineScaleSetExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetExtensions client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetExtensionsClient.Client)

	virtualMachineScaleSetRollingUpgradesClient, err := virtualmachinescalesetrollingupgrades.NewVirtualMachineScaleSetRollingUpgradesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetRollingUpgrades client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetRollingUpgradesClient.Client)

	virtualMachineScaleSetVMExtensionsClient, err := virtualmachinescalesetvmextensions.NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetVMExtensions client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetVMExtensionsClient.Client)

	virtualMachineScaleSetVMRunCommandsClient, err := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetVMRunCommands client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetVMRunCommandsClient.Client)

	virtualMachineScaleSetVMsClient, err := virtualmachinescalesetvms.NewVirtualMachineScaleSetVMsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSetVMs client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetVMsClient.Client)

	virtualMachineScaleSetsClient, err := virtualmachinescalesets.NewVirtualMachineScaleSetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineScaleSets client: %+v", err)
	}
	configureFunc(virtualMachineScaleSetsClient.Client)

	virtualMachineSizesClient, err := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineSizes client: %+v", err)
	}
	configureFunc(virtualMachineSizesClient.Client)

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	return &Client{
		AvailabilitySets:                      availabilitySetsClient,
		CapacityReservation:                   capacityReservationClient,
		CapacityReservationGroups:             capacityReservationGroupsClient,
		CapacityReservations:                  capacityReservationsClient,
		CommunityGalleries:                    communityGalleriesClient,
		CommunityGalleryImageVersions:         communityGalleryImageVersionsClient,
		CommunityGalleryImages:                communityGalleryImagesClient,
		Compute:                               computeClient,
		DedicatedHost:                         dedicatedHostClient,
		DedicatedHostGroups:                   dedicatedHostGroupsClient,
		DedicatedHosts:                        dedicatedHostsClient,
		Galleries:                             galleriesClient,
		GalleryApplicationVersions:            galleryApplicationVersionsClient,
		GalleryApplications:                   galleryApplicationsClient,
		GalleryImageVersions:                  galleryImageVersionsClient,
		GalleryImages:                         galleryImagesClient,
		GallerySharingUpdate:                  gallerySharingUpdateClient,
		Images:                                imagesClient,
		LogAnalytics:                          logAnalyticsClient,
		ProximityPlacementGroups:              proximityPlacementGroupsClient,
		RestorePointCollections:               restorePointCollectionsClient,
		SharedGalleries:                       sharedGalleriesClient,
		SharedGalleryImageVersions:            sharedGalleryImageVersionsClient,
		SharedGalleryImages:                   sharedGalleryImagesClient,
		Skus:                                  skusClient,
		SshPublicKeys:                         sshPublicKeysClient,
		VirtualMachineExtensionImages:         virtualMachineExtensionImagesClient,
		VirtualMachineExtensions:              virtualMachineExtensionsClient,
		VirtualMachineImages:                  virtualMachineImagesClient,
		VirtualMachineRunCommands:             virtualMachineRunCommandsClient,
		VirtualMachineScaleSetExtensions:      virtualMachineScaleSetExtensionsClient,
		VirtualMachineScaleSetRollingUpgrades: virtualMachineScaleSetRollingUpgradesClient,
		VirtualMachineScaleSetVMExtensions:    virtualMachineScaleSetVMExtensionsClient,
		VirtualMachineScaleSetVMRunCommands:   virtualMachineScaleSetVMRunCommandsClient,
		VirtualMachineScaleSetVMs:             virtualMachineScaleSetVMsClient,
		VirtualMachineScaleSets:               virtualMachineScaleSetsClient,
		VirtualMachineSizes:                   virtualMachineSizesClient,
		VirtualMachines:                       virtualMachinesClient,
	}, nil
}
