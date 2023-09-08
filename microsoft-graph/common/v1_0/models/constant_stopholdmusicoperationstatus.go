package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopHoldMusicOperationStatus string

const (
	StopHoldMusicOperationStatusCompleted  StopHoldMusicOperationStatus = "Completed"
	StopHoldMusicOperationStatusFailed     StopHoldMusicOperationStatus = "Failed"
	StopHoldMusicOperationStatusNotStarted StopHoldMusicOperationStatus = "NotStarted"
	StopHoldMusicOperationStatusRunning    StopHoldMusicOperationStatus = "Running"
)

func PossibleValuesForStopHoldMusicOperationStatus() []string {
	return []string{
		string(StopHoldMusicOperationStatusCompleted),
		string(StopHoldMusicOperationStatusFailed),
		string(StopHoldMusicOperationStatusNotStarted),
		string(StopHoldMusicOperationStatusRunning),
	}
}

func parseStopHoldMusicOperationStatus(input string) (*StopHoldMusicOperationStatus, error) {
	vals := map[string]StopHoldMusicOperationStatus{
		"completed":  StopHoldMusicOperationStatusCompleted,
		"failed":     StopHoldMusicOperationStatusFailed,
		"notstarted": StopHoldMusicOperationStatusNotStarted,
		"running":    StopHoldMusicOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StopHoldMusicOperationStatus(input)
	return &out, nil
}
