package virtualmachine

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineProperties struct {
	ClaimedByUserId        *string                          `json:"claimedByUserId,omitempty"`
	ConnectionProfile      *VirtualMachineConnectionProfile `json:"connectionProfile,omitempty"`
	ProvisioningState      *ProvisioningState               `json:"provisioningState,omitempty"`
	ResourceOperationError *ResourceOperationError          `json:"resourceOperationError,omitempty"`
	State                  *VirtualMachineState             `json:"state,omitempty"`
	VMType                 *VirtualMachineType              `json:"vmType,omitempty"`
}
