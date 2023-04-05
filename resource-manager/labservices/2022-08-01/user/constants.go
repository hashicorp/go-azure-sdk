package user

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InvitationState string

const (
	InvitationStateFailed  InvitationState = "Failed"
	InvitationStateNotSent InvitationState = "NotSent"
	InvitationStateSending InvitationState = "Sending"
	InvitationStateSent    InvitationState = "Sent"
)

func PossibleValuesForInvitationState() []string {
	return []string{
		string(InvitationStateFailed),
		string(InvitationStateNotSent),
		string(InvitationStateSending),
		string(InvitationStateSent),
	}
}

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

type RegistrationState string

const (
	RegistrationStateNotRegistered RegistrationState = "NotRegistered"
	RegistrationStateRegistered    RegistrationState = "Registered"
)

func PossibleValuesForRegistrationState() []string {
	return []string{
		string(RegistrationStateNotRegistered),
		string(RegistrationStateRegistered),
	}
}
