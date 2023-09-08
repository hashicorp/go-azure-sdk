package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordOperationStatus string

const (
	RecordOperationStatusCompleted  RecordOperationStatus = "Completed"
	RecordOperationStatusFailed     RecordOperationStatus = "Failed"
	RecordOperationStatusNotStarted RecordOperationStatus = "NotStarted"
	RecordOperationStatusRunning    RecordOperationStatus = "Running"
)

func PossibleValuesForRecordOperationStatus() []string {
	return []string{
		string(RecordOperationStatusCompleted),
		string(RecordOperationStatusFailed),
		string(RecordOperationStatusNotStarted),
		string(RecordOperationStatusRunning),
	}
}

func parseRecordOperationStatus(input string) (*RecordOperationStatus, error) {
	vals := map[string]RecordOperationStatus{
		"completed":  RecordOperationStatusCompleted,
		"failed":     RecordOperationStatusFailed,
		"notstarted": RecordOperationStatusNotStarted,
		"running":    RecordOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordOperationStatus(input)
	return &out, nil
}
