package jitrequests

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JitRequestState string

const (
	JitRequestStateApproved     JitRequestState = "Approved"
	JitRequestStateCanceled     JitRequestState = "Canceled"
	JitRequestStateDenied       JitRequestState = "Denied"
	JitRequestStateExpired      JitRequestState = "Expired"
	JitRequestStateFailed       JitRequestState = "Failed"
	JitRequestStateNotSpecified JitRequestState = "NotSpecified"
	JitRequestStatePending      JitRequestState = "Pending"
	JitRequestStateTimeout      JitRequestState = "Timeout"
)

func PossibleValuesForJitRequestState() []string {
	return []string{
		string(JitRequestStateApproved),
		string(JitRequestStateCanceled),
		string(JitRequestStateDenied),
		string(JitRequestStateExpired),
		string(JitRequestStateFailed),
		string(JitRequestStateNotSpecified),
		string(JitRequestStatePending),
		string(JitRequestStateTimeout),
	}
}

func (s *JitRequestState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJitRequestState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJitRequestState(input string) (*JitRequestState, error) {
	vals := map[string]JitRequestState{
		"approved":     JitRequestStateApproved,
		"canceled":     JitRequestStateCanceled,
		"denied":       JitRequestStateDenied,
		"expired":      JitRequestStateExpired,
		"failed":       JitRequestStateFailed,
		"notspecified": JitRequestStateNotSpecified,
		"pending":      JitRequestStatePending,
		"timeout":      JitRequestStateTimeout,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JitRequestState(input)
	return &out, nil
}

type JitSchedulingType string

const (
	JitSchedulingTypeNotSpecified JitSchedulingType = "NotSpecified"
	JitSchedulingTypeOnce         JitSchedulingType = "Once"
	JitSchedulingTypeRecurring    JitSchedulingType = "Recurring"
)

func PossibleValuesForJitSchedulingType() []string {
	return []string{
		string(JitSchedulingTypeNotSpecified),
		string(JitSchedulingTypeOnce),
		string(JitSchedulingTypeRecurring),
	}
}

func (s *JitSchedulingType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseJitSchedulingType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseJitSchedulingType(input string) (*JitSchedulingType, error) {
	vals := map[string]JitSchedulingType{
		"notspecified": JitSchedulingTypeNotSpecified,
		"once":         JitSchedulingTypeOnce,
		"recurring":    JitSchedulingTypeRecurring,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JitSchedulingType(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateAccepted     ProvisioningState = "Accepted"
	ProvisioningStateCanceled     ProvisioningState = "Canceled"
	ProvisioningStateCreated      ProvisioningState = "Created"
	ProvisioningStateCreating     ProvisioningState = "Creating"
	ProvisioningStateDeleted      ProvisioningState = "Deleted"
	ProvisioningStateDeleting     ProvisioningState = "Deleting"
	ProvisioningStateFailed       ProvisioningState = "Failed"
	ProvisioningStateNotSpecified ProvisioningState = "NotSpecified"
	ProvisioningStateReady        ProvisioningState = "Ready"
	ProvisioningStateRunning      ProvisioningState = "Running"
	ProvisioningStateSucceeded    ProvisioningState = "Succeeded"
	ProvisioningStateUpdating     ProvisioningState = "Updating"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateAccepted),
		string(ProvisioningStateCanceled),
		string(ProvisioningStateCreated),
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleted),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateNotSpecified),
		string(ProvisioningStateReady),
		string(ProvisioningStateRunning),
		string(ProvisioningStateSucceeded),
		string(ProvisioningStateUpdating),
	}
}

func (s *ProvisioningState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseProvisioningState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"accepted":     ProvisioningStateAccepted,
		"canceled":     ProvisioningStateCanceled,
		"created":      ProvisioningStateCreated,
		"creating":     ProvisioningStateCreating,
		"deleted":      ProvisioningStateDeleted,
		"deleting":     ProvisioningStateDeleting,
		"failed":       ProvisioningStateFailed,
		"notspecified": ProvisioningStateNotSpecified,
		"ready":        ProvisioningStateReady,
		"running":      ProvisioningStateRunning,
		"succeeded":    ProvisioningStateSucceeded,
		"updating":     ProvisioningStateUpdating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
