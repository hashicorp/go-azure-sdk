package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscribeToToneOperationStatus string

const (
	SubscribeToToneOperationStatus_Completed  SubscribeToToneOperationStatus = "Completed"
	SubscribeToToneOperationStatus_Failed     SubscribeToToneOperationStatus = "Failed"
	SubscribeToToneOperationStatus_NotStarted SubscribeToToneOperationStatus = "NotStarted"
	SubscribeToToneOperationStatus_Running    SubscribeToToneOperationStatus = "Running"
)

func PossibleValuesForSubscribeToToneOperationStatus() []string {
	return []string{
		string(SubscribeToToneOperationStatus_Completed),
		string(SubscribeToToneOperationStatus_Failed),
		string(SubscribeToToneOperationStatus_NotStarted),
		string(SubscribeToToneOperationStatus_Running),
	}
}

func (s *SubscribeToToneOperationStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSubscribeToToneOperationStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSubscribeToToneOperationStatus(input string) (*SubscribeToToneOperationStatus, error) {
	vals := map[string]SubscribeToToneOperationStatus{
		"completed":  SubscribeToToneOperationStatus_Completed,
		"failed":     SubscribeToToneOperationStatus_Failed,
		"notstarted": SubscribeToToneOperationStatus_NotStarted,
		"running":    SubscribeToToneOperationStatus_Running,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SubscribeToToneOperationStatus(input)
	return &out, nil
}
