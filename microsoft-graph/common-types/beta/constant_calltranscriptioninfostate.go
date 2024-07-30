package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscriptionInfoState string

const (
	CallTranscriptionInfoState_Active     CallTranscriptionInfoState = "active"
	CallTranscriptionInfoState_Inactive   CallTranscriptionInfoState = "inactive"
	CallTranscriptionInfoState_NotStarted CallTranscriptionInfoState = "notStarted"
)

func PossibleValuesForCallTranscriptionInfoState() []string {
	return []string{
		string(CallTranscriptionInfoState_Active),
		string(CallTranscriptionInfoState_Inactive),
		string(CallTranscriptionInfoState_NotStarted),
	}
}

func (s *CallTranscriptionInfoState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCallTranscriptionInfoState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCallTranscriptionInfoState(input string) (*CallTranscriptionInfoState, error) {
	vals := map[string]CallTranscriptionInfoState{
		"active":     CallTranscriptionInfoState_Active,
		"inactive":   CallTranscriptionInfoState_Inactive,
		"notstarted": CallTranscriptionInfoState_NotStarted,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CallTranscriptionInfoState(input)
	return &out, nil
}
