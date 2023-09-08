package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CancelMediaProcessingOperationStatus string

const (
	CancelMediaProcessingOperationStatusCompleted  CancelMediaProcessingOperationStatus = "Completed"
	CancelMediaProcessingOperationStatusFailed     CancelMediaProcessingOperationStatus = "Failed"
	CancelMediaProcessingOperationStatusNotStarted CancelMediaProcessingOperationStatus = "NotStarted"
	CancelMediaProcessingOperationStatusRunning    CancelMediaProcessingOperationStatus = "Running"
)

func PossibleValuesForCancelMediaProcessingOperationStatus() []string {
	return []string{
		string(CancelMediaProcessingOperationStatusCompleted),
		string(CancelMediaProcessingOperationStatusFailed),
		string(CancelMediaProcessingOperationStatusNotStarted),
		string(CancelMediaProcessingOperationStatusRunning),
	}
}

func parseCancelMediaProcessingOperationStatus(input string) (*CancelMediaProcessingOperationStatus, error) {
	vals := map[string]CancelMediaProcessingOperationStatus{
		"completed":  CancelMediaProcessingOperationStatusCompleted,
		"failed":     CancelMediaProcessingOperationStatusFailed,
		"notstarted": CancelMediaProcessingOperationStatusNotStarted,
		"running":    CancelMediaProcessingOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CancelMediaProcessingOperationStatus(input)
	return &out, nil
}
