package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OperationStatus string

const (
	OperationStatusCompleted  OperationStatus = "Completed"
	OperationStatusFailed     OperationStatus = "Failed"
	OperationStatusNotStarted OperationStatus = "NotStarted"
	OperationStatusRunning    OperationStatus = "Running"
)

func PossibleValuesForOperationStatus() []string {
	return []string{
		string(OperationStatusCompleted),
		string(OperationStatusFailed),
		string(OperationStatusNotStarted),
		string(OperationStatusRunning),
	}
}

func parseOperationStatus(input string) (*OperationStatus, error) {
	vals := map[string]OperationStatus{
		"completed":  OperationStatusCompleted,
		"failed":     OperationStatusFailed,
		"notstarted": OperationStatusNotStarted,
		"running":    OperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OperationStatus(input)
	return &out, nil
}
