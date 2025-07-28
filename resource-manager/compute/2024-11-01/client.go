package v2024_11_01

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import (
	"fmt"

	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/availabilitysets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/capacityreservationgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/capacityreservations"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/computerps"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/dedicatedhostgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/dedicatedhosts"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/images"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/proximityplacementgroups"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/restorepointcollections"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/restorepoints"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/rollingupgradestatusinfos"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/sshpublickeyresources"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachineextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachineruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachines"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachinescalesetextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachinescalesets"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachinescalesetvmextensions"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachinescalesetvmruncommands"
	"github.com/hashicorp/go-azure-sdk/resource-manager/compute/2024-11-01/virtualmachinescalesetvms"
	"github.com/hashicorp/go-azure-sdk/sdk/client/resourcemanager"
	sdkEnv "github.com/hashicorp/go-azure-sdk/sdk/environments"
)

type Client struct {
	AvailabilitySets                    *availabilitysets.AvailabilitySetsClient
	CapacityReservationGroups           *capacityreservationgroups.CapacityReservationGroupsClient
	CapacityReservations                *capacityreservations.CapacityReservationsClient
	ComputeRPS                          *computerps.ComputeRPSClient
	DedicatedHostGroups                 *dedicatedhostgroups.DedicatedHostGroupsClient
	DedicatedHosts                      *dedicatedhosts.DedicatedHostsClient
	Images                              *images.ImagesClient
	ProximityPlacementGroups            *proximityplacementgroups.ProximityPlacementGroupsClient
	RestorePointCollections             *restorepointcollections.RestorePointCollectionsClient
	RestorePoints                       *restorepoints.RestorePointsClient
	RollingUpgradeStatusInfos           *rollingupgradestatusinfos.RollingUpgradeStatusInfosClient
	SshPublicKeyResources               *sshpublickeyresources.SshPublicKeyResourcesClient
	VirtualMachineExtensions            *virtualmachineextensions.VirtualMachineExtensionsClient
	VirtualMachineRunCommands           *virtualmachineruncommands.VirtualMachineRunCommandsClient
	VirtualMachineScaleSetExtensions    *virtualmachinescalesetextensions.VirtualMachineScaleSetExtensionsClient
	VirtualMachineScaleSetVMExtensions  *virtualmachinescalesetvmextensions.VirtualMachineScaleSetVMExtensionsClient
	VirtualMachineScaleSetVMRunCommands *virtualmachinescalesetvmruncommands.VirtualMachineScaleSetVMRunCommandsClient
	VirtualMachineScaleSetVMs           *virtualmachinescalesetvms.VirtualMachineScaleSetVMsClient
	VirtualMachineScaleSets             *virtualmachinescalesets.VirtualMachineScaleSetsClient
	VirtualMachines                     *virtualmachines.VirtualMachinesClient
}

func NewClientWithBaseURI(sdkApi sdkEnv.Api, configureFunc func(c *resourcemanager.Client)) (*Client, error) {
	availabilitySetsClient, err := availabilitysets.NewAvailabilitySetsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building AvailabilitySets client: %+v", err)
	}
	configureFunc(availabilitySetsClient.Client)

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

	computeRPSClient, err := computerps.NewComputeRPSClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building ComputeRPS client: %+v", err)
	}
	configureFunc(computeRPSClient.Client)

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

	imagesClient, err := images.NewImagesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building Images client: %+v", err)
	}
	configureFunc(imagesClient.Client)

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

	restorePointsClient, err := restorepoints.NewRestorePointsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RestorePoints client: %+v", err)
	}
	configureFunc(restorePointsClient.Client)

	rollingUpgradeStatusInfosClient, err := rollingupgradestatusinfos.NewRollingUpgradeStatusInfosClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building RollingUpgradeStatusInfos client: %+v", err)
	}
	configureFunc(rollingUpgradeStatusInfosClient.Client)

	sshPublicKeyResourcesClient, err := sshpublickeyresources.NewSshPublicKeyResourcesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building SshPublicKeyResources client: %+v", err)
	}
	configureFunc(sshPublicKeyResourcesClient.Client)

	virtualMachineExtensionsClient, err := virtualmachineextensions.NewVirtualMachineExtensionsClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachineExtensions client: %+v", err)
	}
	configureFunc(virtualMachineExtensionsClient.Client)

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

	virtualMachinesClient, err := virtualmachines.NewVirtualMachinesClientWithBaseURI(sdkApi)
	if err != nil {
		return nil, fmt.Errorf("building VirtualMachines client: %+v", err)
	}
	configureFunc(virtualMachinesClient.Client)

	return &Client{
		AvailabilitySets:                    availabilitySetsClient,
		CapacityReservationGroups:           capacityReservationGroupsClient,
		CapacityReservations:                capacityReservationsClient,
		ComputeRPS:                          computeRPSClient,
		DedicatedHostGroups:                 dedicatedHostGroupsClient,
		DedicatedHosts:                      dedicatedHostsClient,
		Images:                              imagesClient,
		ProximityPlacementGroups:            proximityPlacementGroupsClient,
		RestorePointCollections:             restorePointCollectionsClient,
		RestorePoints:                       restorePointsClient,
		RollingUpgradeStatusInfos:           rollingUpgradeStatusInfosClient,
		SshPublicKeyResources:               sshPublicKeyResourcesClient,
		VirtualMachineExtensions:            virtualMachineExtensionsClient,
		VirtualMachineRunCommands:           virtualMachineRunCommandsClient,
		VirtualMachineScaleSetExtensions:    virtualMachineScaleSetExtensionsClient,
		VirtualMachineScaleSetVMExtensions:  virtualMachineScaleSetVMExtensionsClient,
		VirtualMachineScaleSetVMRunCommands: virtualMachineScaleSetVMRunCommandsClient,
		VirtualMachineScaleSetVMs:           virtualMachineScaleSetVMsClient,
		VirtualMachineScaleSets:             virtualMachineScaleSetsClient,
		VirtualMachines:                     virtualMachinesClient,
	}, nil
}
