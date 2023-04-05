package virtualmachinescalesetvmruncommands

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
