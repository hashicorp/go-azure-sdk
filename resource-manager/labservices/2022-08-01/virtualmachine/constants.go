package virtualmachine

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateLocked    ProvisioningState = "Locked"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateLocked),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type VirtualMachineState string

const (
	VirtualMachineStateRedeploying       VirtualMachineState = "Redeploying"
	VirtualMachineStateReimaging         VirtualMachineState = "Reimaging"
	VirtualMachineStateResettingPassword VirtualMachineState = "ResettingPassword"
	VirtualMachineStateRunning           VirtualMachineState = "Running"
	VirtualMachineStateStarting          VirtualMachineState = "Starting"
	VirtualMachineStateStopped           VirtualMachineState = "Stopped"
	VirtualMachineStateStopping          VirtualMachineState = "Stopping"
)

func PossibleValuesForVirtualMachineState() []string {
	return []string{
		string(VirtualMachineStateRedeploying),
		string(VirtualMachineStateReimaging),
		string(VirtualMachineStateResettingPassword),
		string(VirtualMachineStateRunning),
		string(VirtualMachineStateStarting),
		string(VirtualMachineStateStopped),
		string(VirtualMachineStateStopping),
	}
}

type VirtualMachineType string

const (
	VirtualMachineTypeTemplate VirtualMachineType = "Template"
	VirtualMachineTypeUser     VirtualMachineType = "User"
)

func PossibleValuesForVirtualMachineType() []string {
	return []string{
		string(VirtualMachineTypeTemplate),
		string(VirtualMachineTypeUser),
	}
}
