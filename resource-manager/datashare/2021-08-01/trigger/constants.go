package trigger

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisioningState string

const (
	ProvisioningStateCreating  ProvisioningState = "Creating"
	ProvisioningStateDeleting  ProvisioningState = "Deleting"
	ProvisioningStateFailed    ProvisioningState = "Failed"
	ProvisioningStateMoving    ProvisioningState = "Moving"
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateCreating),
		string(ProvisioningStateDeleting),
		string(ProvisioningStateFailed),
		string(ProvisioningStateMoving),
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"creating":  ProvisioningStateCreating,
		"deleting":  ProvisioningStateDeleting,
		"failed":    ProvisioningStateFailed,
		"moving":    ProvisioningStateMoving,
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}

type RecurrenceInterval string

const (
	RecurrenceIntervalDay  RecurrenceInterval = "Day"
	RecurrenceIntervalHour RecurrenceInterval = "Hour"
)

func PossibleValuesForRecurrenceInterval() []string {
	return []string{
		string(RecurrenceIntervalDay),
		string(RecurrenceIntervalHour),
	}
}

func parseRecurrenceInterval(input string) (*RecurrenceInterval, error) {
	vals := map[string]RecurrenceInterval{
		"day":  RecurrenceIntervalDay,
		"hour": RecurrenceIntervalHour,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecurrenceInterval(input)
	return &out, nil
}

type Status string

const (
	StatusAccepted         Status = "Accepted"
	StatusCanceled         Status = "Canceled"
	StatusFailed           Status = "Failed"
	StatusInProgress       Status = "InProgress"
	StatusSucceeded        Status = "Succeeded"
	StatusTransientFailure Status = "TransientFailure"
)

func PossibleValuesForStatus() []string {
	return []string{
		string(StatusAccepted),
		string(StatusCanceled),
		string(StatusFailed),
		string(StatusInProgress),
		string(StatusSucceeded),
		string(StatusTransientFailure),
	}
}

func parseStatus(input string) (*Status, error) {
	vals := map[string]Status{
		"accepted":         StatusAccepted,
		"canceled":         StatusCanceled,
		"failed":           StatusFailed,
		"inprogress":       StatusInProgress,
		"succeeded":        StatusSucceeded,
		"transientfailure": StatusTransientFailure,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Status(input)
	return &out, nil
}

type SynchronizationMode string

const (
	SynchronizationModeFullSync    SynchronizationMode = "FullSync"
	SynchronizationModeIncremental SynchronizationMode = "Incremental"
)

func PossibleValuesForSynchronizationMode() []string {
	return []string{
		string(SynchronizationModeFullSync),
		string(SynchronizationModeIncremental),
	}
}

func parseSynchronizationMode(input string) (*SynchronizationMode, error) {
	vals := map[string]SynchronizationMode{
		"fullsync":    SynchronizationModeFullSync,
		"incremental": SynchronizationModeIncremental,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SynchronizationMode(input)
	return &out, nil
}

type TriggerKind string

const (
	TriggerKindScheduleBased TriggerKind = "ScheduleBased"
)

func PossibleValuesForTriggerKind() []string {
	return []string{
		string(TriggerKindScheduleBased),
	}
}

func parseTriggerKind(input string) (*TriggerKind, error) {
	vals := map[string]TriggerKind{
		"schedulebased": TriggerKindScheduleBased,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerKind(input)
	return &out, nil
}

type TriggerStatus string

const (
	TriggerStatusActive                              TriggerStatus = "Active"
	TriggerStatusInactive                            TriggerStatus = "Inactive"
	TriggerStatusSourceSynchronizationSettingDeleted TriggerStatus = "SourceSynchronizationSettingDeleted"
)

func PossibleValuesForTriggerStatus() []string {
	return []string{
		string(TriggerStatusActive),
		string(TriggerStatusInactive),
		string(TriggerStatusSourceSynchronizationSettingDeleted),
	}
}

func parseTriggerStatus(input string) (*TriggerStatus, error) {
	vals := map[string]TriggerStatus{
		"active":                              TriggerStatusActive,
		"inactive":                            TriggerStatusInactive,
		"sourcesynchronizationsettingdeleted": TriggerStatusSourceSynchronizationSettingDeleted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TriggerStatus(input)
	return &out, nil
}
