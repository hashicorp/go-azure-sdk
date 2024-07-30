package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivateDeviceEsimActionResultActionState string

const (
	ActivateDeviceEsimActionResultActionState_Active       ActivateDeviceEsimActionResultActionState = "active"
	ActivateDeviceEsimActionResultActionState_Canceled     ActivateDeviceEsimActionResultActionState = "canceled"
	ActivateDeviceEsimActionResultActionState_Done         ActivateDeviceEsimActionResultActionState = "done"
	ActivateDeviceEsimActionResultActionState_Failed       ActivateDeviceEsimActionResultActionState = "failed"
	ActivateDeviceEsimActionResultActionState_None         ActivateDeviceEsimActionResultActionState = "none"
	ActivateDeviceEsimActionResultActionState_NotSupported ActivateDeviceEsimActionResultActionState = "notSupported"
	ActivateDeviceEsimActionResultActionState_Pending      ActivateDeviceEsimActionResultActionState = "pending"
)

func PossibleValuesForActivateDeviceEsimActionResultActionState() []string {
	return []string{
		string(ActivateDeviceEsimActionResultActionState_Active),
		string(ActivateDeviceEsimActionResultActionState_Canceled),
		string(ActivateDeviceEsimActionResultActionState_Done),
		string(ActivateDeviceEsimActionResultActionState_Failed),
		string(ActivateDeviceEsimActionResultActionState_None),
		string(ActivateDeviceEsimActionResultActionState_NotSupported),
		string(ActivateDeviceEsimActionResultActionState_Pending),
	}
}

func (s *ActivateDeviceEsimActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseActivateDeviceEsimActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseActivateDeviceEsimActionResultActionState(input string) (*ActivateDeviceEsimActionResultActionState, error) {
	vals := map[string]ActivateDeviceEsimActionResultActionState{
		"active":       ActivateDeviceEsimActionResultActionState_Active,
		"canceled":     ActivateDeviceEsimActionResultActionState_Canceled,
		"done":         ActivateDeviceEsimActionResultActionState_Done,
		"failed":       ActivateDeviceEsimActionResultActionState_Failed,
		"none":         ActivateDeviceEsimActionResultActionState_None,
		"notsupported": ActivateDeviceEsimActionResultActionState_NotSupported,
		"pending":      ActivateDeviceEsimActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ActivateDeviceEsimActionResultActionState(input)
	return &out, nil
}
