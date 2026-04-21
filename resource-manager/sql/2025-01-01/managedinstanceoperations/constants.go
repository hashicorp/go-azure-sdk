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

type UpsertManagedServerOperationStepWithEstimatesAndDurationStatus string

const (
	UpsertManagedServerOperationStepWithEstimatesAndDurationStatusCanceled   UpsertManagedServerOperationStepWithEstimatesAndDurationStatus = "Canceled"
	UpsertManagedServerOperationStepWithEstimatesAndDurationStatusCompleted  UpsertManagedServerOperationStepWithEstimatesAndDurationStatus = "Completed"
	UpsertManagedServerOperationStepWithEstimatesAndDurationStatusFailed     UpsertManagedServerOperationStepWithEstimatesAndDurationStatus = "Failed"
	UpsertManagedServerOperationStepWithEstimatesAndDurationStatusInProgress UpsertManagedServerOperationStepWithEstimatesAndDurationStatus = "InProgress"
	UpsertManagedServerOperationStepWithEstimatesAndDurationStatusNotStarted UpsertManagedServerOperationStepWithEstimatesAndDurationStatus = "NotStarted"
	UpsertManagedServerOperationStepWithEstimatesAndDurationStatusSlowedDown UpsertManagedServerOperationStepWithEstimatesAndDurationStatus = "SlowedDown"
)

func PossibleValuesForUpsertManagedServerOperationStepWithEstimatesAndDurationStatus() []string {
	return []string{
		string(UpsertManagedServerOperationStepWithEstimatesAndDurationStatusCanceled),
		string(UpsertManagedServerOperationStepWithEstimatesAndDurationStatusCompleted),
		string(UpsertManagedServerOperationStepWithEstimatesAndDurationStatusFailed),
		string(UpsertManagedServerOperationStepWithEstimatesAndDurationStatusInProgress),
		string(UpsertManagedServerOperationStepWithEstimatesAndDurationStatusNotStarted),
		string(UpsertManagedServerOperationStepWithEstimatesAndDurationStatusSlowedDown),
	}
}

func (s *UpsertManagedServerOperationStepWithEstimatesAndDurationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUpsertManagedServerOperationStepWithEstimatesAndDurationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUpsertManagedServerOperationStepWithEstimatesAndDurationStatus(input string) (*UpsertManagedServerOperationStepWithEstimatesAndDurationStatus, error) {
	vals := map[string]UpsertManagedServerOperationStepWithEstimatesAndDurationStatus{
		"canceled":   UpsertManagedServerOperationStepWithEstimatesAndDurationStatusCanceled,
		"completed":  UpsertManagedServerOperationStepWithEstimatesAndDurationStatusCompleted,
		"failed":     UpsertManagedServerOperationStepWithEstimatesAndDurationStatusFailed,
		"inprogress": UpsertManagedServerOperationStepWithEstimatesAndDurationStatusInProgress,
		"notstarted": UpsertManagedServerOperationStepWithEstimatesAndDurationStatusNotStarted,
		"sloweddown": UpsertManagedServerOperationStepWithEstimatesAndDurationStatusSlowedDown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpsertManagedServerOperationStepWithEstimatesAndDurationStatus(input)
	return &out, nil
}
