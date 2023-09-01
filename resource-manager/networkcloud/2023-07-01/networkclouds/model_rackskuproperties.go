package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RackSkuProperties struct {
	ComputeMachines     *[]MachineSkuSlot          `json:"computeMachines,omitempty"`
	ControllerMachines  *[]MachineSkuSlot          `json:"controllerMachines,omitempty"`
	Description         *string                    `json:"description,omitempty"`
	MaxClusterSlots     *int64                     `json:"maxClusterSlots,omitempty"`
	ProvisioningState   *RackSkuProvisioningState  `json:"provisioningState,omitempty"`
	RackType            *RackSkuType               `json:"rackType,omitempty"`
	StorageAppliances   *[]StorageApplianceSkuSlot `json:"storageAppliances,omitempty"`
	SupportedRackSkuIds *[]string                  `json:"supportedRackSkuIds,omitempty"`
}
