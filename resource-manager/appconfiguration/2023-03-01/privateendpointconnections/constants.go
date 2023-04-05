package privateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActionsRequired string

const (
	ActionsRequiredNone     ActionsRequired = "None"
	ActionsRequiredRecreate ActionsRequired = "Recreate"
)

func PossibleValuesForActionsRequired() []string {
	return []string{
		string(ActionsRequiredNone),
		string(ActionsRequiredRecreate),
	}
}

type ConnectionStatus string

const (
	ConnectionStatusApproved     ConnectionStatus = "Approved"
	ConnectionStatusDisconnected ConnectionStatus = "Disconnected"
	ConnectionStatusPending      ConnectionStatus = "Pending"
	ConnectionStatusRejected     ConnectionStatus = "Rejected"
)

func PossibleValuesForConnectionStatus() []string {
	return []string{
		string(ConnectionStatusApproved),
		string(ConnectionStatusDisconnected),
		string(ConnectionStatusPending),
		string(ConnectionStatusRejected),
	}
}

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
