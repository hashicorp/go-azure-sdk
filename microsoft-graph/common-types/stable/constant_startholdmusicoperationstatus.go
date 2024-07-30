package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StartHoldMusicOperationStatus string

const (
	StartHoldMusicOperationStatus_Completed  StartHoldMusicOperationStatus = "Completed"
	StartHoldMusicOperationStatus_Failed     StartHoldMusicOperationStatus = "Failed"
	StartHoldMusicOperationStatus_NotStarted StartHoldMusicOperationStatus = "NotStarted"
	StartHoldMusicOperationStatus_Running    StartHoldMusicOperationStatus = "Running"
)

func PossibleValuesForStartHoldMusicOperationStatus() []string {
	return []string{
		string(StartHoldMusicOperationStatus_Completed),
		string(StartHoldMusicOperationStatus_Failed),
		string(StartHoldMusicOperationStatus_NotStarted),
		string(StartHoldMusicOperationStatus_Running),
	}
}

func (s *StartHoldMusicOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseStartHoldMusicOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseStartHoldMusicOperationStatus(input string) (*StartHoldMusicOperationStatus, error) {
	vals := map[string]StartHoldMusicOperationStatus{
		"completed":  StartHoldMusicOperationStatus_Completed,
		"failed":     StartHoldMusicOperationStatus_Failed,
		"notstarted": StartHoldMusicOperationStatus_NotStarted,
		"running":    StartHoldMusicOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := StartHoldMusicOperationStatus(input)
	return &out, nil
}
