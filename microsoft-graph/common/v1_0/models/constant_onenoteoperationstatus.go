package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteOperationStatus string

const (
	OnenoteOperationStatusCompleted  OnenoteOperationStatus = "Completed"
	OnenoteOperationStatusFailed     OnenoteOperationStatus = "Failed"
	OnenoteOperationStatusNotStarted OnenoteOperationStatus = "NotStarted"
	OnenoteOperationStatusRunning    OnenoteOperationStatus = "Running"
)

func PossibleValuesForOnenoteOperationStatus() []string {
	return []string{
		string(OnenoteOperationStatusCompleted),
		string(OnenoteOperationStatusFailed),
		string(OnenoteOperationStatusNotStarted),
		string(OnenoteOperationStatusRunning),
	}
}

func parseOnenoteOperationStatus(input string) (*OnenoteOperationStatus, error) {
	vals := map[string]OnenoteOperationStatus{
		"completed":  OnenoteOperationStatusCompleted,
		"failed":     OnenoteOperationStatusFailed,
		"notstarted": OnenoteOperationStatusNotStarted,
		"running":    OnenoteOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := OnenoteOperationStatus(input)
	return &out, nil
}
