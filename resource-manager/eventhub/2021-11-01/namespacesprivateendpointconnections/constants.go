package namespacesprivateendpointconnections

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndPointProvisioningState string

const (
	EndPointProvisioningStateCanceled  EndPointProvisioningState = "Canceled"
	EndPointProvisioningStateCreating  EndPointProvisioningState = "Creating"
	EndPointProvisioningStateDeleting  EndPointProvisioningState = "Deleting"
	EndPointProvisioningStateFailed    EndPointProvisioningState = "Failed"
	EndPointProvisioningStateSucceeded EndPointProvisioningState = "Succeeded"
	EndPointProvisioningStateUpdating  EndPointProvisioningState = "Updating"
)

func PossibleValuesForEndPointProvisioningState() []string {
	return []string{
		string(EndPointProvisioningStateCanceled),
		string(EndPointProvisioningStateCreating),
		string(EndPointProvisioningStateDeleting),
		string(EndPointProvisioningStateFailed),
		string(EndPointProvisioningStateSucceeded),
		string(EndPointProvisioningStateUpdating),
	}
}

func parseEndPointProvisioningState(input string) (*EndPointProvisioningState, error) {
	vals := map[string]EndPointProvisioningState{
		"canceled":  EndPointProvisioningStateCanceled,
		"creating":  EndPointProvisioningStateCreating,
		"deleting":  EndPointProvisioningStateDeleting,
		"failed":    EndPointProvisioningStateFailed,
		"succeeded": EndPointProvisioningStateSucceeded,
		"updating":  EndPointProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndPointProvisioningState(input)
	return &out, nil
}

type PrivateLinkConnectionStatus string

const (
	PrivateLinkConnectionStatusApproved     PrivateLinkConnectionStatus = "Approved"
	PrivateLinkConnectionStatusDisconnected PrivateLinkConnectionStatus = "Disconnected"
	PrivateLinkConnectionStatusPending      PrivateLinkConnectionStatus = "Pending"
	PrivateLinkConnectionStatusRejected     PrivateLinkConnectionStatus = "Rejected"
)

func PossibleValuesForPrivateLinkConnectionStatus() []string {
	return []string{
		string(PrivateLinkConnectionStatusApproved),
		string(PrivateLinkConnectionStatusDisconnected),
		string(PrivateLinkConnectionStatusPending),
		string(PrivateLinkConnectionStatusRejected),
	}
}

func parsePrivateLinkConnectionStatus(input string) (*PrivateLinkConnectionStatus, error) {
	vals := map[string]PrivateLinkConnectionStatus{
		"approved":     PrivateLinkConnectionStatusApproved,
		"disconnected": PrivateLinkConnectionStatusDisconnected,
		"pending":      PrivateLinkConnectionStatusPending,
		"rejected":     PrivateLinkConnectionStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkConnectionStatus(input)
	return &out, nil
}
