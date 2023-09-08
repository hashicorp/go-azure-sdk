package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmuteParticipantOperationStatus string

const (
	UnmuteParticipantOperationStatusCompleted  UnmuteParticipantOperationStatus = "Completed"
	UnmuteParticipantOperationStatusFailed     UnmuteParticipantOperationStatus = "Failed"
	UnmuteParticipantOperationStatusNotStarted UnmuteParticipantOperationStatus = "NotStarted"
	UnmuteParticipantOperationStatusRunning    UnmuteParticipantOperationStatus = "Running"
)

func PossibleValuesForUnmuteParticipantOperationStatus() []string {
	return []string{
		string(UnmuteParticipantOperationStatusCompleted),
		string(UnmuteParticipantOperationStatusFailed),
		string(UnmuteParticipantOperationStatusNotStarted),
		string(UnmuteParticipantOperationStatusRunning),
	}
}

func parseUnmuteParticipantOperationStatus(input string) (*UnmuteParticipantOperationStatus, error) {
	vals := map[string]UnmuteParticipantOperationStatus{
		"completed":  UnmuteParticipantOperationStatusCompleted,
		"failed":     UnmuteParticipantOperationStatusFailed,
		"notstarted": UnmuteParticipantOperationStatusNotStarted,
		"running":    UnmuteParticipantOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmuteParticipantOperationStatus(input)
	return &out, nil
}
