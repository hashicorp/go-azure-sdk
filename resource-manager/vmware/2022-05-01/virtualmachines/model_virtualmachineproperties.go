package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineProperties struct {
	DisplayName      *string                              `json:"displayName,omitempty"`
	FolderPath       *string                              `json:"folderPath,omitempty"`
	MoRefId          *string                              `json:"moRefId,omitempty"`
	RestrictMovement *VirtualMachineRestrictMovementState `json:"restrictMovement,omitempty"`
}
