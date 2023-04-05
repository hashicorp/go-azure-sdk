package privateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateLinkServiceConnectionProvisioningState string

const (
	PrivateLinkServiceConnectionProvisioningStateCanceled   PrivateLinkServiceConnectionProvisioningState = "Canceled"
	PrivateLinkServiceConnectionProvisioningStateDeleting   PrivateLinkServiceConnectionProvisioningState = "Deleting"
	PrivateLinkServiceConnectionProvisioningStateFailed     PrivateLinkServiceConnectionProvisioningState = "Failed"
	PrivateLinkServiceConnectionProvisioningStateIncomplete PrivateLinkServiceConnectionProvisioningState = "Incomplete"
	PrivateLinkServiceConnectionProvisioningStateSucceeded  PrivateLinkServiceConnectionProvisioningState = "Succeeded"
	PrivateLinkServiceConnectionProvisioningStateUpdating   PrivateLinkServiceConnectionProvisioningState = "Updating"
)

func PossibleValuesForPrivateLinkServiceConnectionProvisioningState() []string {
	return []string{
		string(PrivateLinkServiceConnectionProvisioningStateCanceled),
		string(PrivateLinkServiceConnectionProvisioningStateDeleting),
		string(PrivateLinkServiceConnectionProvisioningStateFailed),
		string(PrivateLinkServiceConnectionProvisioningStateIncomplete),
		string(PrivateLinkServiceConnectionProvisioningStateSucceeded),
		string(PrivateLinkServiceConnectionProvisioningStateUpdating),
	}
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusApproved     PrivateLinkServiceConnectionStatus = "Approved"
	PrivateLinkServiceConnectionStatusDisconnected PrivateLinkServiceConnectionStatus = "Disconnected"
	PrivateLinkServiceConnectionStatusPending      PrivateLinkServiceConnectionStatus = "Pending"
	PrivateLinkServiceConnectionStatusRejected     PrivateLinkServiceConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStatusApproved),
		string(PrivateLinkServiceConnectionStatusDisconnected),
		string(PrivateLinkServiceConnectionStatusPending),
		string(PrivateLinkServiceConnectionStatusRejected),
	}
}
