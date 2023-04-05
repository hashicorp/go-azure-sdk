package virtualmachines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualMachineRestrictMovementState string

const (
	VirtualMachineRestrictMovementStateDisabled VirtualMachineRestrictMovementState = "Disabled"
	VirtualMachineRestrictMovementStateEnabled  VirtualMachineRestrictMovementState = "Enabled"
)

func PossibleValuesForVirtualMachineRestrictMovementState() []string {
	return []string{
		string(VirtualMachineRestrictMovementStateDisabled),
		string(VirtualMachineRestrictMovementStateEnabled),
	}
}
