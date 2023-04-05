package account

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Name string

const (
	NameStandard Name = "Standard"
)

func PossibleValuesForName() []string {
	return []string{
		string(NameStandard),
	}
}

type ProvisioningState string

const (
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateMoving       ProvisioningState = "Moving"
	ProvisioningStateSoftDeleted  ProvisioningState = "SoftDeleted"
	ProvisioningStateSoftDeleting ProvisioningState = "SoftDeleting"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUnknown      ProvisioningState = "Unknown"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSoftDeleted),
		string(ProvisioningStateSoftDeleting),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUnknown),
	}
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessDisabled     PublicNetworkAccess = "Disabled"
	PublicNetworkAccessEnabled      PublicNetworkAccess = "Enabled"
	PublicNetworkAccessNotSpecified PublicNetworkAccess = "NotSpecified"
)

func PossibleValuesForPublicNetworkAccess() []string {
	return []string{
		string(PublicNetworkAccessDisabled),
		string(PublicNetworkAccessEnabled),
		string(PublicNetworkAccessNotSpecified),
	}
}

type Status string

const (
	StatusApproved     Status = "Approved"
	StatusDisconnected Status = "Disconnected"
	StatusPending      Status = "Pending"
	StatusRejected     Status = "Rejected"
	StatusUnknown      Status = "Unknown"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusApproved),
		string(StatusDisconnected),
		string(StatusPending),
		string(StatusRejected),
		string(StatusUnknown),
	}
}
