package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VppTokenActionResultActionState string

const (
	VppTokenActionResultActionState_Active       VppTokenActionResultActionState = "active"
	VppTokenActionResultActionState_Canceled     VppTokenActionResultActionState = "canceled"
	VppTokenActionResultActionState_Done         VppTokenActionResultActionState = "done"
	VppTokenActionResultActionState_Failed       VppTokenActionResultActionState = "failed"
	VppTokenActionResultActionState_None         VppTokenActionResultActionState = "none"
	VppTokenActionResultActionState_NotSupported VppTokenActionResultActionState = "notSupported"
	VppTokenActionResultActionState_Pending      VppTokenActionResultActionState = "pending"
)

func PossibleValuesForVppTokenActionResultActionState() []string {
	return []string{
		string(VppTokenActionResultActionState_Active),
		string(VppTokenActionResultActionState_Canceled),
		string(VppTokenActionResultActionState_Done),
		string(VppTokenActionResultActionState_Failed),
		string(VppTokenActionResultActionState_None),
		string(VppTokenActionResultActionState_NotSupported),
		string(VppTokenActionResultActionState_Pending),
	}
}

func (s *VppTokenActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseVppTokenActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseVppTokenActionResultActionState(input string) (*VppTokenActionResultActionState, error) {
	vals := map[string]VppTokenActionResultActionState{
		"active":       VppTokenActionResultActionState_Active,
		"canceled":     VppTokenActionResultActionState_Canceled,
		"done":         VppTokenActionResultActionState_Done,
		"failed":       VppTokenActionResultActionState_Failed,
		"none":         VppTokenActionResultActionState_None,
		"notsupported": VppTokenActionResultActionState_NotSupported,
		"pending":      VppTokenActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VppTokenActionResultActionState(input)
	return &out, nil
}
