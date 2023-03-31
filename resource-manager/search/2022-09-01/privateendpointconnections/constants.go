package privateendpointconnections

import "strings"

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

func parsePrivateLinkServiceConnectionProvisioningState(input string) (*PrivateLinkServiceConnectionProvisioningState, error) {
	vals := map[string]PrivateLinkServiceConnectionProvisioningState{
		"canceled":   PrivateLinkServiceConnectionProvisioningStateCanceled,
		"deleting":   PrivateLinkServiceConnectionProvisioningStateDeleting,
		"failed":     PrivateLinkServiceConnectionProvisioningStateFailed,
		"incomplete": PrivateLinkServiceConnectionProvisioningStateIncomplete,
		"succeeded":  PrivateLinkServiceConnectionProvisioningStateSucceeded,
		"updating":   PrivateLinkServiceConnectionProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkServiceConnectionProvisioningState(input)
	return &out, nil
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

func parsePrivateLinkServiceConnectionStatus(input string) (*PrivateLinkServiceConnectionStatus, error) {
	vals := map[string]PrivateLinkServiceConnectionStatus{
		"approved":     PrivateLinkServiceConnectionStatusApproved,
		"disconnected": PrivateLinkServiceConnectionStatusDisconnected,
		"pending":      PrivateLinkServiceConnectionStatusPending,
		"rejected":     PrivateLinkServiceConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkServiceConnectionStatus(input)
	return &out, nil
}
