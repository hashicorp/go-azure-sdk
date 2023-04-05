package singlesignon

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

type SingleSignOnStates string

const (
	SingleSignOnStatesDisable  SingleSignOnStates = "Disable"
	SingleSignOnStatesEnable   SingleSignOnStates = "Enable"
	SingleSignOnStatesExisting SingleSignOnStates = "Existing"
	SingleSignOnStatesInitial  SingleSignOnStates = "Initial"
)

func PossibleValuesForSingleSignOnStates() []string {
	return []string{
		string(SingleSignOnStatesDisable),
		string(SingleSignOnStatesEnable),
		string(SingleSignOnStatesExisting),
		string(SingleSignOnStatesInitial),
	}
}
