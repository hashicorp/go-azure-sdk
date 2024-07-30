package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RotateBitLockerKeysDeviceActionResultActionState string

const (
	RotateBitLockerKeysDeviceActionResultActionState_Active       RotateBitLockerKeysDeviceActionResultActionState = "active"
	RotateBitLockerKeysDeviceActionResultActionState_Canceled     RotateBitLockerKeysDeviceActionResultActionState = "canceled"
	RotateBitLockerKeysDeviceActionResultActionState_Done         RotateBitLockerKeysDeviceActionResultActionState = "done"
	RotateBitLockerKeysDeviceActionResultActionState_Failed       RotateBitLockerKeysDeviceActionResultActionState = "failed"
	RotateBitLockerKeysDeviceActionResultActionState_None         RotateBitLockerKeysDeviceActionResultActionState = "none"
	RotateBitLockerKeysDeviceActionResultActionState_NotSupported RotateBitLockerKeysDeviceActionResultActionState = "notSupported"
	RotateBitLockerKeysDeviceActionResultActionState_Pending      RotateBitLockerKeysDeviceActionResultActionState = "pending"
)

func PossibleValuesForRotateBitLockerKeysDeviceActionResultActionState() []string {
	return []string{
		string(RotateBitLockerKeysDeviceActionResultActionState_Active),
		string(RotateBitLockerKeysDeviceActionResultActionState_Canceled),
		string(RotateBitLockerKeysDeviceActionResultActionState_Done),
		string(RotateBitLockerKeysDeviceActionResultActionState_Failed),
		string(RotateBitLockerKeysDeviceActionResultActionState_None),
		string(RotateBitLockerKeysDeviceActionResultActionState_NotSupported),
		string(RotateBitLockerKeysDeviceActionResultActionState_Pending),
	}
}

func (s *RotateBitLockerKeysDeviceActionResultActionState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRotateBitLockerKeysDeviceActionResultActionState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRotateBitLockerKeysDeviceActionResultActionState(input string) (*RotateBitLockerKeysDeviceActionResultActionState, error) {
	vals := map[string]RotateBitLockerKeysDeviceActionResultActionState{
		"active":       RotateBitLockerKeysDeviceActionResultActionState_Active,
		"canceled":     RotateBitLockerKeysDeviceActionResultActionState_Canceled,
		"done":         RotateBitLockerKeysDeviceActionResultActionState_Done,
		"failed":       RotateBitLockerKeysDeviceActionResultActionState_Failed,
		"none":         RotateBitLockerKeysDeviceActionResultActionState_None,
		"notsupported": RotateBitLockerKeysDeviceActionResultActionState_NotSupported,
		"pending":      RotateBitLockerKeysDeviceActionResultActionState_Pending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RotateBitLockerKeysDeviceActionResultActionState(input)
	return &out, nil
}
