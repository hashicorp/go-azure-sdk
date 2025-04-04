package managedinstanceoperations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagementOperationState string

const (
	ManagementOperationStateCancelInProgress ManagementOperationState = "CancelInProgress"
	ManagementOperationStateCancelled        ManagementOperationState = "Cancelled"
	ManagementOperationStateFailed           ManagementOperationState = "Failed"
	ManagementOperationStateInProgress       ManagementOperationState = "InProgress"
	ManagementOperationStatePending          ManagementOperationState = "Pending"
	ManagementOperationStateSucceeded        ManagementOperationState = "Succeeded"
)

func PossibleValuesForManagementOperationState() []string {
	return []string{
		string(ManagementOperationStateCancelInProgress),
		string(ManagementOperationStateCancelled),
		string(ManagementOperationStateFailed),
		string(ManagementOperationStateInProgress),
		string(ManagementOperationStatePending),
		string(ManagementOperationStateSucceeded),
	}
}

func (s *ManagementOperationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagementOperationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagementOperationState(input string) (*ManagementOperationState, error) {
	vals := map[string]ManagementOperationState{
		"cancelinprogress": ManagementOperationStateCancelInProgress,
		"cancelled":        ManagementOperationStateCancelled,
		"failed":           ManagementOperationStateFailed,
		"inprogress":       ManagementOperationStateInProgress,
		"pending":          ManagementOperationStatePending,
		"succeeded":        ManagementOperationStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementOperationState(input)
	return &out, nil
}

type ManagementOperationStepState string

const (
	ManagementOperationStepStateCanceled   ManagementOperationStepState = "Canceled"
	ManagementOperationStepStateCompleted  ManagementOperationStepState = "Completed"
	ManagementOperationStepStateFailed     ManagementOperationStepState = "Failed"
	ManagementOperationStepStateInProgress ManagementOperationStepState = "InProgress"
	ManagementOperationStepStateNotStarted ManagementOperationStepState = "NotStarted"
	ManagementOperationStepStateSlowedDown ManagementOperationStepState = "SlowedDown"
)

func PossibleValuesForManagementOperationStepState() []string {
	return []string{
		string(ManagementOperationStepStateCanceled),
		string(ManagementOperationStepStateCompleted),
		string(ManagementOperationStepStateFailed),
		string(ManagementOperationStepStateInProgress),
		string(ManagementOperationStepStateNotStarted),
		string(ManagementOperationStepStateSlowedDown),
	}
}

func (s *ManagementOperationStepState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagementOperationStepState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagementOperationStepState(input string) (*ManagementOperationStepState, error) {
	vals := map[string]ManagementOperationStepState{
		"canceled":   ManagementOperationStepStateCanceled,
		"completed":  ManagementOperationStepStateCompleted,
		"failed":     ManagementOperationStepStateFailed,
		"inprogress": ManagementOperationStepStateInProgress,
		"notstarted": ManagementOperationStepStateNotStarted,
		"sloweddown": ManagementOperationStepStateSlowedDown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagementOperationStepState(input)
	return &out, nil
}
