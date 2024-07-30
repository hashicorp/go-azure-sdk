package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlayPromptOperationStatus string

const (
	PlayPromptOperationStatus_Completed  PlayPromptOperationStatus = "Completed"
	PlayPromptOperationStatus_Failed     PlayPromptOperationStatus = "Failed"
	PlayPromptOperationStatus_NotStarted PlayPromptOperationStatus = "NotStarted"
	PlayPromptOperationStatus_Running    PlayPromptOperationStatus = "Running"
)

func PossibleValuesForPlayPromptOperationStatus() []string {
	return []string{
		string(PlayPromptOperationStatus_Completed),
		string(PlayPromptOperationStatus_Failed),
		string(PlayPromptOperationStatus_NotStarted),
		string(PlayPromptOperationStatus_Running),
	}
}

func (s *PlayPromptOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePlayPromptOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePlayPromptOperationStatus(input string) (*PlayPromptOperationStatus, error) {
	vals := map[string]PlayPromptOperationStatus{
		"completed":  PlayPromptOperationStatus_Completed,
		"failed":     PlayPromptOperationStatus_Failed,
		"notstarted": PlayPromptOperationStatus_NotStarted,
		"running":    PlayPromptOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PlayPromptOperationStatus(input)
	return &out, nil
}
