package privateendpointconnections

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointProvisioningState string

const (
	PrivateEndpointProvisioningStateApproving PrivateEndpointProvisioningState = "Approving"
	PrivateEndpointProvisioningStateDropping  PrivateEndpointProvisioningState = "Dropping"
	PrivateEndpointProvisioningStateFailed    PrivateEndpointProvisioningState = "Failed"
	PrivateEndpointProvisioningStateReady     PrivateEndpointProvisioningState = "Ready"
	PrivateEndpointProvisioningStateRejecting PrivateEndpointProvisioningState = "Rejecting"
)

func PossibleValuesForPrivateEndpointProvisioningState() []string {
	return []string{
		string(PrivateEndpointProvisioningStateApproving),
		string(PrivateEndpointProvisioningStateDropping),
		string(PrivateEndpointProvisioningStateFailed),
		string(PrivateEndpointProvisioningStateReady),
		string(PrivateEndpointProvisioningStateRejecting),
	}
}

func (s *PrivateEndpointProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateEndpointProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateEndpointProvisioningState(input string) (*PrivateEndpointProvisioningState, error) {
	vals := map[string]PrivateEndpointProvisioningState{
		"approving": PrivateEndpointProvisioningStateApproving,
		"dropping":  PrivateEndpointProvisioningStateDropping,
		"failed":    PrivateEndpointProvisioningStateFailed,
		"ready":     PrivateEndpointProvisioningStateReady,
		"rejecting": PrivateEndpointProvisioningStateRejecting,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateEndpointProvisioningState(input)
	return &out, nil
}

type PrivateLinkServiceConnectionStateActionsRequire string

const (
	PrivateLinkServiceConnectionStateActionsRequireNone PrivateLinkServiceConnectionStateActionsRequire = "None"
)

func PossibleValuesForPrivateLinkServiceConnectionStateActionsRequire() []string {
	return []string{
		string(PrivateLinkServiceConnectionStateActionsRequireNone),
	}
}

func (s *PrivateLinkServiceConnectionStateActionsRequire) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateLinkServiceConnectionStateActionsRequire(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateLinkServiceConnectionStateActionsRequire(input string) (*PrivateLinkServiceConnectionStateActionsRequire, error) {
	vals := map[string]PrivateLinkServiceConnectionStateActionsRequire{
		"none": PrivateLinkServiceConnectionStateActionsRequireNone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkServiceConnectionStateActionsRequire(input)
	return &out, nil
}

type PrivateLinkServiceConnectionStateStatus string

const (
	PrivateLinkServiceConnectionStateStatusApproved     PrivateLinkServiceConnectionStateStatus = "Approved"
	PrivateLinkServiceConnectionStateStatusDisconnected PrivateLinkServiceConnectionStateStatus = "Disconnected"
	PrivateLinkServiceConnectionStateStatusPending      PrivateLinkServiceConnectionStateStatus = "Pending"
	PrivateLinkServiceConnectionStateStatusRejected     PrivateLinkServiceConnectionStateStatus = "Rejected"
)

func PossibleValuesForPrivateLinkServiceConnectionStateStatus() []string {
	return []string{
		string(PrivateLinkServiceConnectionStateStatusApproved),
		string(PrivateLinkServiceConnectionStateStatusDisconnected),
		string(PrivateLinkServiceConnectionStateStatusPending),
		string(PrivateLinkServiceConnectionStateStatusRejected),
	}
}

func (s *PrivateLinkServiceConnectionStateStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrivateLinkServiceConnectionStateStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrivateLinkServiceConnectionStateStatus(input string) (*PrivateLinkServiceConnectionStateStatus, error) {
	vals := map[string]PrivateLinkServiceConnectionStateStatus{
		"approved":     PrivateLinkServiceConnectionStateStatusApproved,
		"disconnected": PrivateLinkServiceConnectionStateStatusDisconnected,
		"pending":      PrivateLinkServiceConnectionStateStatusPending,
		"rejected":     PrivateLinkServiceConnectionStateStatusRejected,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrivateLinkServiceConnectionStateStatus(input)
	return &out, nil
}
