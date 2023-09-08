package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlayPromptOperationStatus string

const (
	PlayPromptOperationStatusCompleted  PlayPromptOperationStatus = "Completed"
	PlayPromptOperationStatusFailed     PlayPromptOperationStatus = "Failed"
	PlayPromptOperationStatusNotStarted PlayPromptOperationStatus = "NotStarted"
	PlayPromptOperationStatusRunning    PlayPromptOperationStatus = "Running"
)

func PossibleValuesForPlayPromptOperationStatus() []string {
	return []string{
		string(PlayPromptOperationStatusCompleted),
		string(PlayPromptOperationStatusFailed),
		string(PlayPromptOperationStatusNotStarted),
		string(PlayPromptOperationStatusRunning),
	}
}

func parsePlayPromptOperationStatus(input string) (*PlayPromptOperationStatus, error) {
	vals := map[string]PlayPromptOperationStatus{
		"completed":  PlayPromptOperationStatusCompleted,
		"failed":     PlayPromptOperationStatusFailed,
		"notstarted": PlayPromptOperationStatusNotStarted,
		"running":    PlayPromptOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlayPromptOperationStatus(input)
	return &out, nil
}
