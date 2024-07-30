package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocateDeviceActionResultActionState string

const (
	LocateDeviceActionResultActionState_Active       LocateDeviceActionResultActionState = "active"
	LocateDeviceActionResultActionState_Canceled     LocateDeviceActionResultActionState = "canceled"
	LocateDeviceActionResultActionState_Done         LocateDeviceActionResultActionState = "done"
	LocateDeviceActionResultActionState_Failed       LocateDeviceActionResultActionState = "failed"
	LocateDeviceActionResultActionState_None         LocateDeviceActionResultActionState = "none"
	LocateDeviceActionResultActionState_NotSupported LocateDeviceActionResultActionState = "notSupported"
	LocateDeviceActionResultActionState_Pending      LocateDeviceActionResultActionState = "pending"
)

func PossibleValuesForLocateDeviceActionResultActionState() []string {
	return []string{
		string(LocateDeviceActionResultActionState_Active),
		string(LocateDeviceActionResultActionState_Canceled),
		string(LocateDeviceActionResultActionState_Done),
		string(LocateDeviceActionResultActionState_Failed),
		string(LocateDeviceActionResultActionState_None),
		string(LocateDeviceActionResultActionState_NotSupported),
		string(LocateDeviceActionResultActionState_Pending),
	}
}

func (s *LocateDeviceActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLocateDeviceActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLocateDeviceActionResultActionState(input string) (*LocateDeviceActionResultActionState, error) {
	vals := map[string]LocateDeviceActionResultActionState{
		"active":       LocateDeviceActionResultActionState_Active,
		"canceled":     LocateDeviceActionResultActionState_Canceled,
		"done":         LocateDeviceActionResultActionState_Done,
		"failed":       LocateDeviceActionResultActionState_Failed,
		"none":         LocateDeviceActionResultActionState_None,
		"notsupported": LocateDeviceActionResultActionState_NotSupported,
		"pending":      LocateDeviceActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocateDeviceActionResultActionState(input)
	return &out, nil
}
