package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UpdateRecordingStatusOperationStatus string

const (
	UpdateRecordingStatusOperationStatusCompleted  UpdateRecordingStatusOperationStatus = "Completed"
	UpdateRecordingStatusOperationStatusFailed     UpdateRecordingStatusOperationStatus = "Failed"
	UpdateRecordingStatusOperationStatusNotStarted UpdateRecordingStatusOperationStatus = "NotStarted"
	UpdateRecordingStatusOperationStatusRunning    UpdateRecordingStatusOperationStatus = "Running"
)

func PossibleValuesForUpdateRecordingStatusOperationStatus() []string {
	return []string{
		string(UpdateRecordingStatusOperationStatusCompleted),
		string(UpdateRecordingStatusOperationStatusFailed),
		string(UpdateRecordingStatusOperationStatusNotStarted),
		string(UpdateRecordingStatusOperationStatusRunning),
	}
}

func parseUpdateRecordingStatusOperationStatus(input string) (*UpdateRecordingStatusOperationStatus, error) {
	vals := map[string]UpdateRecordingStatusOperationStatus{
		"completed":  UpdateRecordingStatusOperationStatusCompleted,
		"failed":     UpdateRecordingStatusOperationStatusFailed,
		"notstarted": UpdateRecordingStatusOperationStatusNotStarted,
		"running":    UpdateRecordingStatusOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UpdateRecordingStatusOperationStatus(input)
	return &out, nil
}
