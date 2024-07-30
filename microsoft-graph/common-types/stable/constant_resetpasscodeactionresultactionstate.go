package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResetPasscodeActionResultActionState string

const (
	ResetPasscodeActionResultActionState_Active       ResetPasscodeActionResultActionState = "active"
	ResetPasscodeActionResultActionState_Canceled     ResetPasscodeActionResultActionState = "canceled"
	ResetPasscodeActionResultActionState_Done         ResetPasscodeActionResultActionState = "done"
	ResetPasscodeActionResultActionState_Failed       ResetPasscodeActionResultActionState = "failed"
	ResetPasscodeActionResultActionState_None         ResetPasscodeActionResultActionState = "none"
	ResetPasscodeActionResultActionState_NotSupported ResetPasscodeActionResultActionState = "notSupported"
	ResetPasscodeActionResultActionState_Pending      ResetPasscodeActionResultActionState = "pending"
)

func PossibleValuesForResetPasscodeActionResultActionState() []string {
	return []string{
		string(ResetPasscodeActionResultActionState_Active),
		string(ResetPasscodeActionResultActionState_Canceled),
		string(ResetPasscodeActionResultActionState_Done),
		string(ResetPasscodeActionResultActionState_Failed),
		string(ResetPasscodeActionResultActionState_None),
		string(ResetPasscodeActionResultActionState_NotSupported),
		string(ResetPasscodeActionResultActionState_Pending),
	}
}

func (s *ResetPasscodeActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseResetPasscodeActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseResetPasscodeActionResultActionState(input string) (*ResetPasscodeActionResultActionState, error) {
	vals := map[string]ResetPasscodeActionResultActionState{
		"active":       ResetPasscodeActionResultActionState_Active,
		"canceled":     ResetPasscodeActionResultActionState_Canceled,
		"done":         ResetPasscodeActionResultActionState_Done,
		"failed":       ResetPasscodeActionResultActionState_Failed,
		"none":         ResetPasscodeActionResultActionState_None,
		"notsupported": ResetPasscodeActionResultActionState_NotSupported,
		"pending":      ResetPasscodeActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ResetPasscodeActionResultActionState(input)
	return &out, nil
}
