package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StopHoldMusicOperationStatus string

const (
	StopHoldMusicOperationStatus_Completed  StopHoldMusicOperationStatus = "Completed"
	StopHoldMusicOperationStatus_Failed     StopHoldMusicOperationStatus = "Failed"
	StopHoldMusicOperationStatus_NotStarted StopHoldMusicOperationStatus = "NotStarted"
	StopHoldMusicOperationStatus_Running    StopHoldMusicOperationStatus = "Running"
)

func PossibleValuesForStopHoldMusicOperationStatus() []string {
	return []string{
		string(StopHoldMusicOperationStatus_Completed),
		string(StopHoldMusicOperationStatus_Failed),
		string(StopHoldMusicOperationStatus_NotStarted),
		string(StopHoldMusicOperationStatus_Running),
	}
}

func (s *StopHoldMusicOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStopHoldMusicOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStopHoldMusicOperationStatus(input string) (*StopHoldMusicOperationStatus, error) {
	vals := map[string]StopHoldMusicOperationStatus{
		"completed":  StopHoldMusicOperationStatus_Completed,
		"failed":     StopHoldMusicOperationStatus_Failed,
		"notstarted": StopHoldMusicOperationStatus_NotStarted,
		"running":    StopHoldMusicOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StopHoldMusicOperationStatus(input)
	return &out, nil
}
