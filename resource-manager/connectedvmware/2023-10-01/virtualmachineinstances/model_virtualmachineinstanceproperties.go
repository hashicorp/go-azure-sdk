package virtualmachineinstances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineInstanceProperties struct {
	HardwareProfile       *HardwareProfile        `json:"hardwareProfile,omitempty"`
	InfrastructureProfile *InfrastructureProfile  `json:"infrastructureProfile,omitempty"`
	NetworkProfile        *NetworkProfile         `json:"networkProfile,omitempty"`
	OsProfile             *OsProfileForVMInstance `json:"osProfile,omitempty"`
	PlacementProfile      *PlacementProfile       `json:"placementProfile,omitempty"`
	PowerState            *string                 `json:"powerState,omitempty"`
	ProvisioningState     *ProvisioningState      `json:"provisioningState,omitempty"`
	ResourceUid           *string                 `json:"resourceUid,omitempty"`
	SecurityProfile       *SecurityProfile        `json:"securityProfile,omitempty"`
	Statuses              *[]ResourceStatus       `json:"statuses,omitempty"`
	StorageProfile        *StorageProfile         `json:"storageProfile,omitempty"`
}
