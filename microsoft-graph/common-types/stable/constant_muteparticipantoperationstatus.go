package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MuteParticipantOperationStatus string

const (
	MuteParticipantOperationStatus_Completed  MuteParticipantOperationStatus = "Completed"
	MuteParticipantOperationStatus_Failed     MuteParticipantOperationStatus = "Failed"
	MuteParticipantOperationStatus_NotStarted MuteParticipantOperationStatus = "NotStarted"
	MuteParticipantOperationStatus_Running    MuteParticipantOperationStatus = "Running"
)

func PossibleValuesForMuteParticipantOperationStatus() []string {
	return []string{
		string(MuteParticipantOperationStatus_Completed),
		string(MuteParticipantOperationStatus_Failed),
		string(MuteParticipantOperationStatus_NotStarted),
		string(MuteParticipantOperationStatus_Running),
	}
}

func (s *MuteParticipantOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMuteParticipantOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMuteParticipantOperationStatus(input string) (*MuteParticipantOperationStatus, error) {
	vals := map[string]MuteParticipantOperationStatus{
		"completed":  MuteParticipantOperationStatus_Completed,
		"failed":     MuteParticipantOperationStatus_Failed,
		"notstarted": MuteParticipantOperationStatus_NotStarted,
		"running":    MuteParticipantOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MuteParticipantOperationStatus(input)
	return &out, nil
}
