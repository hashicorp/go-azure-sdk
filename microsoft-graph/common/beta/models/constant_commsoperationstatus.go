package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommsOperationStatus string

const (
	CommsOperationStatusCompleted  CommsOperationStatus = "Completed"
	CommsOperationStatusFailed     CommsOperationStatus = "Failed"
	CommsOperationStatusNotStarted CommsOperationStatus = "NotStarted"
	CommsOperationStatusRunning    CommsOperationStatus = "Running"
)

func PossibleValuesForCommsOperationStatus() []string {
	return []string{
		string(CommsOperationStatusCompleted),
		string(CommsOperationStatusFailed),
		string(CommsOperationStatusNotStarted),
		string(CommsOperationStatusRunning),
	}
}

func parseCommsOperationStatus(input string) (*CommsOperationStatus, error) {
	vals := map[string]CommsOperationStatus{
		"completed":  CommsOperationStatusCompleted,
		"failed":     CommsOperationStatusFailed,
		"notstarted": CommsOperationStatusNotStarted,
		"running":    CommsOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CommsOperationStatus(input)
	return &out, nil
}
