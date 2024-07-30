package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InviteParticipantsOperationStatus string

const (
	InviteParticipantsOperationStatus_Completed  InviteParticipantsOperationStatus = "Completed"
	InviteParticipantsOperationStatus_Failed     InviteParticipantsOperationStatus = "Failed"
	InviteParticipantsOperationStatus_NotStarted InviteParticipantsOperationStatus = "NotStarted"
	InviteParticipantsOperationStatus_Running    InviteParticipantsOperationStatus = "Running"
)

func PossibleValuesForInviteParticipantsOperationStatus() []string {
	return []string{
		string(InviteParticipantsOperationStatus_Completed),
		string(InviteParticipantsOperationStatus_Failed),
		string(InviteParticipantsOperationStatus_NotStarted),
		string(InviteParticipantsOperationStatus_Running),
	}
}

func (s *InviteParticipantsOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseInviteParticipantsOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseInviteParticipantsOperationStatus(input string) (*InviteParticipantsOperationStatus, error) {
	vals := map[string]InviteParticipantsOperationStatus{
		"completed":  InviteParticipantsOperationStatus_Completed,
		"failed":     InviteParticipantsOperationStatus_Failed,
		"notstarted": InviteParticipantsOperationStatus_NotStarted,
		"running":    InviteParticipantsOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := InviteParticipantsOperationStatus(input)
	return &out, nil
}
