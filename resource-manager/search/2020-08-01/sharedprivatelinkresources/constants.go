package sharedprivatelinkresources

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SharedPrivateLinkResourceProvisioningState string

const (
	SharedPrivateLinkResourceProvisioningStateDeleting   SharedPrivateLinkResourceProvisioningState = "Deleting"
	SharedPrivateLinkResourceProvisioningStateFailed     SharedPrivateLinkResourceProvisioningState = "Failed"
	SharedPrivateLinkResourceProvisioningStateIncomplete SharedPrivateLinkResourceProvisioningState = "Incomplete"
	SharedPrivateLinkResourceProvisioningStateSucceeded  SharedPrivateLinkResourceProvisioningState = "Succeeded"
	SharedPrivateLinkResourceProvisioningStateUpdating   SharedPrivateLinkResourceProvisioningState = "Updating"
)

func PossibleValuesForSharedPrivateLinkResourceProvisioningState() []string {
	return []string{
		string(SharedPrivateLinkResourceProvisioningStateDeleting),
		string(SharedPrivateLinkResourceProvisioningStateFailed),
		string(SharedPrivateLinkResourceProvisioningStateIncomplete),
		string(SharedPrivateLinkResourceProvisioningStateSucceeded),
		string(SharedPrivateLinkResourceProvisioningStateUpdating),
	}
}

type SharedPrivateLinkResourceStatus string

const (
	SharedPrivateLinkResourceStatusApproved     SharedPrivateLinkResourceStatus = "Approved"
	SharedPrivateLinkResourceStatusDisconnected SharedPrivateLinkResourceStatus = "Disconnected"
	SharedPrivateLinkResourceStatusPending      SharedPrivateLinkResourceStatus = "Pending"
	SharedPrivateLinkResourceStatusRejected     SharedPrivateLinkResourceStatus = "Rejected"
)

func PossibleValuesForSharedPrivateLinkResourceStatus() []string {
	return []string{
		string(SharedPrivateLinkResourceStatusApproved),
		string(SharedPrivateLinkResourceStatusDisconnected),
		string(SharedPrivateLinkResourceStatusPending),
		string(SharedPrivateLinkResourceStatusRejected),
	}
}
