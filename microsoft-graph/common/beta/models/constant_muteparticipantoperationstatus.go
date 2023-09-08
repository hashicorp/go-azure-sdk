package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MuteParticipantOperationStatus string

const (
	MuteParticipantOperationStatusCompleted  MuteParticipantOperationStatus = "Completed"
	MuteParticipantOperationStatusFailed     MuteParticipantOperationStatus = "Failed"
	MuteParticipantOperationStatusNotStarted MuteParticipantOperationStatus = "NotStarted"
	MuteParticipantOperationStatusRunning    MuteParticipantOperationStatus = "Running"
)

func PossibleValuesForMuteParticipantOperationStatus() []string {
	return []string{
		string(MuteParticipantOperationStatusCompleted),
		string(MuteParticipantOperationStatusFailed),
		string(MuteParticipantOperationStatusNotStarted),
		string(MuteParticipantOperationStatusRunning),
	}
}

func parseMuteParticipantOperationStatus(input string) (*MuteParticipantOperationStatus, error) {
	vals := map[string]MuteParticipantOperationStatus{
		"completed":  MuteParticipantOperationStatusCompleted,
		"failed":     MuteParticipantOperationStatusFailed,
		"notstarted": MuteParticipantOperationStatusNotStarted,
		"running":    MuteParticipantOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MuteParticipantOperationStatus(input)
	return &out, nil
}
