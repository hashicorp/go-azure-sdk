package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartHoldMusicOperationStatus string

const (
	StartHoldMusicOperationStatusCompleted  StartHoldMusicOperationStatus = "Completed"
	StartHoldMusicOperationStatusFailed     StartHoldMusicOperationStatus = "Failed"
	StartHoldMusicOperationStatusNotStarted StartHoldMusicOperationStatus = "NotStarted"
	StartHoldMusicOperationStatusRunning    StartHoldMusicOperationStatus = "Running"
)

func PossibleValuesForStartHoldMusicOperationStatus() []string {
	return []string{
		string(StartHoldMusicOperationStatusCompleted),
		string(StartHoldMusicOperationStatusFailed),
		string(StartHoldMusicOperationStatusNotStarted),
		string(StartHoldMusicOperationStatusRunning),
	}
}

func parseStartHoldMusicOperationStatus(input string) (*StartHoldMusicOperationStatus, error) {
	vals := map[string]StartHoldMusicOperationStatus{
		"completed":  StartHoldMusicOperationStatusCompleted,
		"failed":     StartHoldMusicOperationStatusFailed,
		"notstarted": StartHoldMusicOperationStatusNotStarted,
		"running":    StartHoldMusicOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StartHoldMusicOperationStatus(input)
	return &out, nil
}
