package v2021_07_01

import (
	"github.com/Azure/go-autorest/autorest"
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
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

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

func NewClientWithBaseURI(endpoint string, configureAuthFunc func(c *autorest.Client)) Client {

	availabilitySetsClient := availabilitysets.NewAvailabilitySetsClientWithBaseURI(endpoint)
	configureAuthFunc(&availabilitySetsClient.Client)

	capacityReservationClient := capacityreservation.NewCapacityReservationClientWithBaseURI(endpoint)
	configureAuthFunc(&capacityReservationClient.Client)

	capacityReservationGroupsClient := capacityreservationgroups.NewCapacityReservationGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&capacityReservationGroupsClient.Client)

	capacityReservationsClient := capacityreservations.NewCapacityReservationsClientWithBaseURI(endpoint)
	configureAuthFunc(&capacityReservationsClient.Client)

	communityGalleriesClient := communitygalleries.NewCommunityGalleriesClientWithBaseURI(endpoint)
	configureAuthFunc(&communityGalleriesClient.Client)

	communityGalleryImageVersionsClient := communitygalleryimageversions.NewCommunityGalleryImageVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&communityGalleryImageVersionsClient.Client)

	communityGalleryImagesClient := communitygalleryimages.NewCommunityGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&communityGalleryImagesClient.Client)

	computeClient := compute.NewComputeClientWithBaseURI(endpoint)
	configureAuthFunc(&computeClient.Client)

	dedicatedHostClient := dedicatedhost.NewDedicatedHostClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHostClient.Client)

	dedicatedHostGroupsClient := dedicatedhostgroups.NewDedicatedHostGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHostGroupsClient.Client)

	dedicatedHostsClient := dedicatedhosts.NewDedicatedHostsClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHostsClient.Client)

	galleriesClient := galleries.NewGalleriesClientWithBaseURI(endpoint)
	configureAuthFunc(&galleriesClient.Client)

	galleryApplicationVersionsClient := galleryapplicationversions.NewGalleryApplicationVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryApplicationVersionsClient.Client)

	galleryApplicationsClient := galleryapplications.NewGalleryApplicationsClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryApplicationsClient.Client)

	galleryImageVersionsClient := galleryimageversions.NewGalleryImageVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryImageVersionsClient.Client)

	galleryImagesClient := galleryimages.NewGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&galleryImagesClient.Client)

	gallerySharingUpdateClient := gallerysharingupdate.NewGallerySharingUpdateClientWithBaseURI(endpoint)
	configureAuthFunc(&gallerySharingUpdateClient.Client)

	imagesClient := images.NewImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&imagesClient.Client)

	logAnalyticsClient := loganalytics.NewLogAnalyticsClientWithBaseURI(endpoint)
	configureAuthFunc(&logAnalyticsClient.Client)

	proximityPlacementGroupsClient := proximityplacementgroups.NewProximityPlacementGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&proximityPlacementGroupsClient.Client)

	restorePointCollectionsClient := restorepointcollections.NewRestorePointCollectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&restorePointCollectionsClient.Client)

	sharedGalleriesClient := sharedgalleries.NewSharedGalleriesClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedGalleriesClient.Client)

	sharedGalleryImageVersionsClient := sharedgalleryimageversions.NewSharedGalleryImageVersionsClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedGalleryImageVersionsClient.Client)

	sharedGalleryImagesClient := sharedgalleryimages.NewSharedGalleryImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&sharedGalleryImagesClient.Client)

	skusClient := skus.NewSkusClientWithBaseURI(endpoint)
	configureAuthFunc(&skusClient.Client)

	sshPublicKeysClient := sshpublickeys.NewSshPublicKeysClientWithBaseURI(endpoint)
	configureAuthFunc(&sshPublicKeysClient.Client)

	virtualMachineExtensionImagesClient := virtualmachineextensionimages.NewVirtualMachineExtensionImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineExtensionImagesClient.Client)

	virtualMachineExtensionsClient := virtualmachineextensions.NewVirtualMachineExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineExtensionsClient.Client)

	virtualMachineImagesClient := virtualmachineimages.NewVirtualMachineImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineImagesClient.Client)

	virtualMachineRunCommandsClient := virtualmachineruncommands.NewVirtualMachineRunCommandsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineRunCommandsClient.Client)

	virtualMachineScaleSetExtensionsClient := virtualmachinescalesetextensions.NewVirtualMachineScaleSetExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineScaleSetExtensionsClient.Client)

	virtualMachineScaleSetRollingUpgradesClient := virtualmachinescalesetrollingupgrades.NewVirtualMachineScaleSetRollingUpgradesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineScaleSetRollingUpgradesClient.Client)

	virtualMachineScaleSetVMExtensionsClient := virtualmachinescalesetvmextensions.NewVirtualMachineScaleSetVMExtensionsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineScaleSetVMExtensionsClient.Client)

	virtualMachineScaleSetVMRunCommandsClient := virtualmachinescalesetvmruncommands.NewVirtualMachineScaleSetVMRunCommandsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineScaleSetVMRunCommandsClient.Client)

	virtualMachineScaleSetVMsClient := virtualmachinescalesetvms.NewVirtualMachineScaleSetVMsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineScaleSetVMsClient.Client)

	virtualMachineScaleSetsClient := virtualmachinescalesets.NewVirtualMachineScaleSetsClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineScaleSetsClient.Client)

	virtualMachineSizesClient := virtualmachinesizes.NewVirtualMachineSizesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachineSizesClient.Client)

	virtualMachinesClient := virtualmachines.NewVirtualMachinesClientWithBaseURI(endpoint)
	configureAuthFunc(&virtualMachinesClient.Client)

	return Client{
		AvailabilitySets:                      &availabilitySetsClient,
		CapacityReservation:                   &capacityReservationClient,
		CapacityReservationGroups:             &capacityReservationGroupsClient,
		CapacityReservations:                  &capacityReservationsClient,
		CommunityGalleries:                    &communityGalleriesClient,
		CommunityGalleryImageVersions:         &communityGalleryImageVersionsClient,
		CommunityGalleryImages:                &communityGalleryImagesClient,
		Compute:                               &computeClient,
		DedicatedHost:                         &dedicatedHostClient,
		DedicatedHostGroups:                   &dedicatedHostGroupsClient,
		DedicatedHosts:                        &dedicatedHostsClient,
		Galleries:                             &galleriesClient,
		GalleryApplicationVersions:            &galleryApplicationVersionsClient,
		GalleryApplications:                   &galleryApplicationsClient,
		GalleryImageVersions:                  &galleryImageVersionsClient,
		GalleryImages:                         &galleryImagesClient,
		GallerySharingUpdate:                  &gallerySharingUpdateClient,
		Images:                                &imagesClient,
		LogAnalytics:                          &logAnalyticsClient,
		ProximityPlacementGroups:              &proximityPlacementGroupsClient,
		RestorePointCollections:               &restorePointCollectionsClient,
		SharedGalleries:                       &sharedGalleriesClient,
		SharedGalleryImageVersions:            &sharedGalleryImageVersionsClient,
		SharedGalleryImages:                   &sharedGalleryImagesClient,
		Skus:                                  &skusClient,
		SshPublicKeys:                         &sshPublicKeysClient,
		VirtualMachineExtensionImages:         &virtualMachineExtensionImagesClient,
		VirtualMachineExtensions:              &virtualMachineExtensionsClient,
		VirtualMachineImages:                  &virtualMachineImagesClient,
		VirtualMachineRunCommands:             &virtualMachineRunCommandsClient,
		VirtualMachineScaleSetExtensions:      &virtualMachineScaleSetExtensionsClient,
		VirtualMachineScaleSetRollingUpgrades: &virtualMachineScaleSetRollingUpgradesClient,
		VirtualMachineScaleSetVMExtensions:    &virtualMachineScaleSetVMExtensionsClient,
		VirtualMachineScaleSetVMRunCommands:   &virtualMachineScaleSetVMRunCommandsClient,
		VirtualMachineScaleSetVMs:             &virtualMachineScaleSetVMsClient,
		VirtualMachineScaleSets:               &virtualMachineScaleSetsClient,
		VirtualMachineSizes:                   &virtualMachineSizesClient,
		VirtualMachines:                       &virtualMachinesClient,
	}
}
