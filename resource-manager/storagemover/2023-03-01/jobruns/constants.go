package jobruns

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobRunScanStatus string

const (
	JobRunScanStatusCompleted  JobRunScanStatus = "Completed"
	JobRunScanStatusNotStarted JobRunScanStatus = "NotStarted"
	JobRunScanStatusScanning   JobRunScanStatus = "Scanning"
)

func PossibleValuesForJobRunScanStatus() []string {
	return []string{
		string(JobRunScanStatusCompleted),
		string(JobRunScanStatusNotStarted),
		string(JobRunScanStatusScanning),
	}
}

func parseJobRunScanStatus(input string) (*JobRunScanStatus, error) {
	vals := map[string]JobRunScanStatus{
		"completed":  JobRunScanStatusCompleted,
		"notstarted": JobRunScanStatusNotStarted,
		"scanning":   JobRunScanStatusScanning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobRunScanStatus(input)
	return &out, nil
}

type JobRunStatus string

const (
	JobRunStatusCancelRequested JobRunStatus = "CancelRequested"
	JobRunStatusCanceled        JobRunStatus = "Canceled"
	JobRunStatusCanceling       JobRunStatus = "Canceling"
	JobRunStatusFailed          JobRunStatus = "Failed"
	JobRunStatusQueued          JobRunStatus = "Queued"
	JobRunStatusRunning         JobRunStatus = "Running"
	JobRunStatusStarted         JobRunStatus = "Started"
	JobRunStatusSucceeded       JobRunStatus = "Succeeded"
)

func PossibleValuesForJobRunStatus() []string {
	return []string{
		string(JobRunStatusCancelRequested),
		string(JobRunStatusCanceled),
		string(JobRunStatusCanceling),
		string(JobRunStatusFailed),
		string(JobRunStatusQueued),
		string(JobRunStatusRunning),
		string(JobRunStatusStarted),
		string(JobRunStatusSucceeded),
	}
}

func parseJobRunStatus(input string) (*JobRunStatus, error) {
	vals := map[string]JobRunStatus{
		"cancelrequested": JobRunStatusCancelRequested,
		"canceled":        JobRunStatusCanceled,
		"canceling":       JobRunStatusCanceling,
		"failed":          JobRunStatusFailed,
		"queued":          JobRunStatusQueued,
		"running":         JobRunStatusRunning,
		"started":         JobRunStatusStarted,
		"succeeded":       JobRunStatusSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := JobRunStatus(input)
	return &out, nil
}

type ProvisioningState string

const (
	ProvisioningStateSucceeded ProvisioningState = "Succeeded"
)

func PossibleValuesForProvisioningState() []string {
	return []string{
		string(ProvisioningStateSucceeded),
	}
}

func parseProvisioningState(input string) (*ProvisioningState, error) {
	vals := map[string]ProvisioningState{
		"succeeded": ProvisioningStateSucceeded,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ProvisioningState(input)
	return &out, nil
}
