package virtualmachinescalesets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineScaleSetProperties struct {
	AdditionalCapabilities                 *AdditionalCapabilities          `json:"additionalCapabilities,omitempty"`
	AutomaticRepairsPolicy                 *AutomaticRepairsPolicy          `json:"automaticRepairsPolicy,omitempty"`
	ConstrainedMaximumCapacity             *bool                            `json:"constrainedMaximumCapacity,omitempty"`
	DoNotRunExtensionsOnOverprovisionedVMs *bool                            `json:"doNotRunExtensionsOnOverprovisionedVMs,omitempty"`
	HostGroup                              *SubResource                     `json:"hostGroup,omitempty"`
	OrchestrationMode                      *OrchestrationMode               `json:"orchestrationMode,omitempty"`
	Overprovision                          *bool                            `json:"overprovision,omitempty"`
	PlatformFaultDomainCount               *int64                           `json:"platformFaultDomainCount,omitempty"`
	PriorityMixPolicy                      *PriorityMixPolicy               `json:"priorityMixPolicy,omitempty"`
	ProvisioningState                      *string                          `json:"provisioningState,omitempty"`
	ProximityPlacementGroup                *SubResource                     `json:"proximityPlacementGroup,omitempty"`
	ScaleInPolicy                          *ScaleInPolicy                   `json:"scaleInPolicy,omitempty"`
	SinglePlacementGroup                   *bool                            `json:"singlePlacementGroup,omitempty"`
	SpotRestorePolicy                      *SpotRestorePolicy               `json:"spotRestorePolicy,omitempty"`
	TimeCreated                            *string                          `json:"timeCreated,omitempty"`
	UniqueId                               *string                          `json:"uniqueId,omitempty"`
	UpgradePolicy                          *UpgradePolicy                   `json:"upgradePolicy,omitempty"`
	VirtualMachineProfile                  *VirtualMachineScaleSetVMProfile `json:"virtualMachineProfile,omitempty"`
	ZoneBalance                            *bool                            `json:"zoneBalance,omitempty"`
}
