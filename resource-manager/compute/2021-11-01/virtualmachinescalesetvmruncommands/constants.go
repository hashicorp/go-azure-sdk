package virtualmachinescalesetvmruncommands

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExecutionState string

const (
	ExecutionStateCanceled  ExecutionState = "Canceled"
	ExecutionStateFailed    ExecutionState = "Failed"
	ExecutionStatePending   ExecutionState = "Pending"
	ExecutionStateRunning   ExecutionState = "Running"
	ExecutionStateSucceeded ExecutionState = "Succeeded"
	ExecutionStateTimedOut  ExecutionState = "TimedOut"
	ExecutionStateUnknown   ExecutionState = "Unknown"
)

func PossibleValuesForExecutionState() []string {
	return []string{
		string(ExecutionStateCanceled),
		string(ExecutionStateFailed),
		string(ExecutionStatePending),
		string(ExecutionStateRunning),
		string(ExecutionStateSucceeded),
		string(ExecutionStateTimedOut),
		string(ExecutionStateUnknown),
	}
}

func parseExecutionState(input string) (*ExecutionState, error) {
	vals := map[string]ExecutionState{
		"canceled":  ExecutionStateCanceled,
		"failed":    ExecutionStateFailed,
		"pending":   ExecutionStatePending,
		"running":   ExecutionStateRunning,
		"succeeded": ExecutionStateSucceeded,
		"timedout":  ExecutionStateTimedOut,
		"unknown":   ExecutionStateUnknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExecutionState(input)
	return &out, nil
}

type StatusLevelTypes string

const (
	StatusLevelTypesError   StatusLevelTypes = "Error"
	StatusLevelTypesInfo    StatusLevelTypes = "Info"
	StatusLevelTypesWarning StatusLevelTypes = "Warning"
)

func PossibleValuesForStatusLevelTypes() []string {
	return []string{
		string(StatusLevelTypesError),
		string(StatusLevelTypesInfo),
		string(StatusLevelTypesWarning),
	}
}

func parseStatusLevelTypes(input string) (*StatusLevelTypes, error) {
	vals := map[string]StatusLevelTypes{
		"error":   StatusLevelTypesError,
		"info":    StatusLevelTypesInfo,
		"warning": StatusLevelTypesWarning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StatusLevelTypes(input)
	return &out, nil
}
