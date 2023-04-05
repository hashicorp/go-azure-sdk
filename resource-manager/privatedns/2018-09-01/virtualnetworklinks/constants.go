package virtualnetworklinks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCanceled  ProvisioningState = "Canceled"
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
	ProvisioningStateUpdating  ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type VirtualNetworkLinkState string

const (
	VirtualNetworkLinkStateCompleted  VirtualNetworkLinkState = "Completed"
	VirtualNetworkLinkStateInProgress VirtualNetworkLinkState = "InProgress"
)

func PossibleValuesForVirtualNetworkLinkState() []string {
	return []string{
		string(VirtualNetworkLinkStateCompleted),
		string(VirtualNetworkLinkStateInProgress),
	}
}
