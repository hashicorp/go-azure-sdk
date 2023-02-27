// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package v2021_11_01

import (
	"github.com/Azure/go-autorest/autorest"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/availabilitysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/capacityreservation"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/capacityreservationgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/capacityreservations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/compute"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhost"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhostgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/dedicatedhosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/images"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/loganalytics"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/proximityplacementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/restorepointcollections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/sshpublickeys"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachineextensionimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachineimages"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachineruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinescalesetextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinescalesetrollingupgrades"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinescalesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinescalesetvmextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinescalesetvmruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinescalesetvms"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2021-11-01/virtualmachinesizes"
)

type Client struct {
	AvailabilitySets                      *availabilitysets.AvailabilitySetsClient
	CapacityReservation                   *capacityreservation.CapacityReservationClient
	CapacityReservationGroups             *capacityreservationgroups.CapacityReservationGroupsClient
	CapacityReservations                  *capacityreservations.CapacityReservationsClient
	Compute                               *compute.ComputeClient
	DedicatedHost                         *dedicatedhost.DedicatedHostClient
	DedicatedHostGroups                   *dedicatedhostgroups.DedicatedHostGroupsClient
	DedicatedHosts                        *dedicatedhosts.DedicatedHostsClient
	Images                                *images.ImagesClient
	LogAnalytics                          *loganalytics.LogAnalyticsClient
	ProximityPlacementGroups              *proximityplacementgroups.ProximityPlacementGroupsClient
	RestorePointCollections               *restorepointcollections.RestorePointCollectionsClient
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

	computeClient := compute.NewComputeClientWithBaseURI(endpoint)
	configureAuthFunc(&computeClient.Client)

	dedicatedHostClient := dedicatedhost.NewDedicatedHostClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHostClient.Client)

	dedicatedHostGroupsClient := dedicatedhostgroups.NewDedicatedHostGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHostGroupsClient.Client)

	dedicatedHostsClient := dedicatedhosts.NewDedicatedHostsClientWithBaseURI(endpoint)
	configureAuthFunc(&dedicatedHostsClient.Client)

	imagesClient := images.NewImagesClientWithBaseURI(endpoint)
	configureAuthFunc(&imagesClient.Client)

	logAnalyticsClient := loganalytics.NewLogAnalyticsClientWithBaseURI(endpoint)
	configureAuthFunc(&logAnalyticsClient.Client)

	proximityPlacementGroupsClient := proximityplacementgroups.NewProximityPlacementGroupsClientWithBaseURI(endpoint)
	configureAuthFunc(&proximityPlacementGroupsClient.Client)

	restorePointCollectionsClient := restorepointcollections.NewRestorePointCollectionsClientWithBaseURI(endpoint)
	configureAuthFunc(&restorePointCollectionsClient.Client)

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
		Compute:                               &computeClient,
		DedicatedHost:                         &dedicatedHostClient,
		DedicatedHostGroups:                   &dedicatedHostGroupsClient,
		DedicatedHosts:                        &dedicatedHostsClient,
		Images:                                &imagesClient,
		LogAnalytics:                          &logAnalyticsClient,
		ProximityPlacementGroups:              &proximityPlacementGroupsClient,
		RestorePointCollections:               &restorePointCollectionsClient,
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
