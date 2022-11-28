package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetUpdateProperties struct {
	AdditionalCapabilities                 *AdditionalCapabilities                `json:"additionalCapabilities"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy                `json:"automaticRepairsPolicy"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                                  `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
	Overprovision                          *bool                                  `json:"overprovision,omitempty"`
	ProximityPlacementGroup                *SubResource                           `json:"proximityPlacementGroup"`
	ScaleInPolicy                          *ScaleInPolicy                         `json:"scaleInPolicy"`
	SinglePlacementGroup                   *bool                                  `json:"singlePlacementGroup,omitempty"`
	UpgradePolicy                          *UpgradePolicy                         `json:"upgradePolicy"`
	VirtualMachineProfile                  *VirtualMachineScaleSetUpdateVMProfile `json:"virtualMachineProfile"`
}
