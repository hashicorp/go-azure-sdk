package privatelink

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnectionProvisioningState string

const (
	PrivateEndpointConnectionProvisioningStateCreating        PrivateEndpointConnectionProvisioningState = "Creating"
	PrivateEndpointConnectionProvisioningStateDeleted         PrivateEndpointConnectionProvisioningState = "Deleted"
	PrivateEndpointConnectionProvisioningStateDeleting        PrivateEndpointConnectionProvisioningState = "Deleting"
	PrivateEndpointConnectionProvisioningStateDeletingByProxy PrivateEndpointConnectionProvisioningState = "DeletingByProxy"
	PrivateEndpointConnectionProvisioningStateSucceeded       PrivateEndpointConnectionProvisioningState = "Succeeded"
	PrivateEndpointConnectionProvisioningStateUnknown         PrivateEndpointConnectionProvisioningState = "Unknown"
	PrivateEndpointConnectionProvisioningStateUpdating        PrivateEndpointConnectionProvisioningState = "Updating"
	PrivateEndpointConnectionProvisioningStateUpdatingByProxy PrivateEndpointConnectionProvisioningState = "UpdatingByProxy"
)

func PossibleValuesForPrivateEndpointConnectionProvisioningState() []string {
	return []string{
		string(PrivateEndpointConnectionProvisioningStateCreating),
		string(PrivateEndpointConnectionProvisioningStateDeleted),
		string(PrivateEndpointConnectionProvisioningStateDeleting),
		string(PrivateEndpointConnectionProvisioningStateDeletingByProxy),
		string(PrivateEndpointConnectionProvisioningStateSucceeded),
		string(PrivateEndpointConnectionProvisioningStateUnknown),
		string(PrivateEndpointConnectionProvisioningStateUpdating),
		string(PrivateEndpointConnectionProvisioningStateUpdatingByProxy),
	}
}

func (s *PrivateEndpointConnectionProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateEndpointConnectionProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateEndpointConnectionProvisioningState(input string) (*PrivateEndpointConnectionProvisioningState, error) {
	vals := map[string]PrivateEndpointConnectionProvisioningState{
		"creating":        PrivateEndpointConnectionProvisioningStateCreating,
		"deleted":         PrivateEndpointConnectionProvisioningStateDeleted,
		"deleting":        PrivateEndpointConnectionProvisioningStateDeleting,
		"deletingbyproxy": PrivateEndpointConnectionProvisioningStateDeletingByProxy,
		"succeeded":       PrivateEndpointConnectionProvisioningStateSucceeded,
		"unknown":         PrivateEndpointConnectionProvisioningStateUnknown,
		"updating":        PrivateEndpointConnectionProvisioningStateUpdating,
		"updatingbyproxy": PrivateEndpointConnectionProvisioningStateUpdatingByProxy,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointConnectionProvisioningState(input)
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

func (s *PrivateLinkConnectionStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateLinkConnectionStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
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
