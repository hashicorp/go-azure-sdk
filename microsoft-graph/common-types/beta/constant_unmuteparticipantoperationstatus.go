package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnmuteParticipantOperationStatus string

const (
	UnmuteParticipantOperationStatus_Completed  UnmuteParticipantOperationStatus = "Completed"
	UnmuteParticipantOperationStatus_Failed     UnmuteParticipantOperationStatus = "Failed"
	UnmuteParticipantOperationStatus_NotStarted UnmuteParticipantOperationStatus = "NotStarted"
	UnmuteParticipantOperationStatus_Running    UnmuteParticipantOperationStatus = "Running"
)

func PossibleValuesForUnmuteParticipantOperationStatus() []string {
	return []string{
		string(UnmuteParticipantOperationStatus_Completed),
		string(UnmuteParticipantOperationStatus_Failed),
		string(UnmuteParticipantOperationStatus_NotStarted),
		string(UnmuteParticipantOperationStatus_Running),
	}
}

func (s *UnmuteParticipantOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUnmuteParticipantOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUnmuteParticipantOperationStatus(input string) (*UnmuteParticipantOperationStatus, error) {
	vals := map[string]UnmuteParticipantOperationStatus{
		"completed":  UnmuteParticipantOperationStatus_Completed,
		"failed":     UnmuteParticipantOperationStatus_Failed,
		"notstarted": UnmuteParticipantOperationStatus_NotStarted,
		"running":    UnmuteParticipantOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UnmuteParticipantOperationStatus(input)
	return &out, nil
}
