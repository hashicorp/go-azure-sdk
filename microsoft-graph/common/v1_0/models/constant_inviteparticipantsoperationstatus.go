package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InviteParticipantsOperationStatus string

const (
	InviteParticipantsOperationStatusCompleted  InviteParticipantsOperationStatus = "Completed"
	InviteParticipantsOperationStatusFailed     InviteParticipantsOperationStatus = "Failed"
	InviteParticipantsOperationStatusNotStarted InviteParticipantsOperationStatus = "NotStarted"
	InviteParticipantsOperationStatusRunning    InviteParticipantsOperationStatus = "Running"
)

func PossibleValuesForInviteParticipantsOperationStatus() []string {
	return []string{
		string(InviteParticipantsOperationStatusCompleted),
		string(InviteParticipantsOperationStatusFailed),
		string(InviteParticipantsOperationStatusNotStarted),
		string(InviteParticipantsOperationStatusRunning),
	}
}

func parseInviteParticipantsOperationStatus(input string) (*InviteParticipantsOperationStatus, error) {
	vals := map[string]InviteParticipantsOperationStatus{
		"completed":  InviteParticipantsOperationStatusCompleted,
		"failed":     InviteParticipantsOperationStatusFailed,
		"notstarted": InviteParticipantsOperationStatusNotStarted,
		"running":    InviteParticipantsOperationStatusRunning,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InviteParticipantsOperationStatus(input)
	return &out, nil
}
