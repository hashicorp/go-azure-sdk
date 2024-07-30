package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MuteParticipantsOperationStatus string

const (
	MuteParticipantsOperationStatus_Completed  MuteParticipantsOperationStatus = "Completed"
	MuteParticipantsOperationStatus_Failed     MuteParticipantsOperationStatus = "Failed"
	MuteParticipantsOperationStatus_NotStarted MuteParticipantsOperationStatus = "NotStarted"
	MuteParticipantsOperationStatus_Running    MuteParticipantsOperationStatus = "Running"
)

func PossibleValuesForMuteParticipantsOperationStatus() []string {
	return []string{
		string(MuteParticipantsOperationStatus_Completed),
		string(MuteParticipantsOperationStatus_Failed),
		string(MuteParticipantsOperationStatus_NotStarted),
		string(MuteParticipantsOperationStatus_Running),
	}
}

func (s *MuteParticipantsOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMuteParticipantsOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMuteParticipantsOperationStatus(input string) (*MuteParticipantsOperationStatus, error) {
	vals := map[string]MuteParticipantsOperationStatus{
		"completed":  MuteParticipantsOperationStatus_Completed,
		"failed":     MuteParticipantsOperationStatus_Failed,
		"notstarted": MuteParticipantsOperationStatus_NotStarted,
		"running":    MuteParticipantsOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MuteParticipantsOperationStatus(input)
	return &out, nil
}
