package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RotateBitLockerKeysDeviceActionResultActionState string

const (
	RotateBitLockerKeysDeviceActionResultActionStateactive       RotateBitLockerKeysDeviceActionResultActionState = "Active"
	RotateBitLockerKeysDeviceActionResultActionStatecanceled     RotateBitLockerKeysDeviceActionResultActionState = "Canceled"
	RotateBitLockerKeysDeviceActionResultActionStatedone         RotateBitLockerKeysDeviceActionResultActionState = "Done"
	RotateBitLockerKeysDeviceActionResultActionStatefailed       RotateBitLockerKeysDeviceActionResultActionState = "Failed"
	RotateBitLockerKeysDeviceActionResultActionStatenone         RotateBitLockerKeysDeviceActionResultActionState = "None"
	RotateBitLockerKeysDeviceActionResultActionStatenotSupported RotateBitLockerKeysDeviceActionResultActionState = "NotSupported"
	RotateBitLockerKeysDeviceActionResultActionStatepending      RotateBitLockerKeysDeviceActionResultActionState = "Pending"
)

func PossibleValuesForRotateBitLockerKeysDeviceActionResultActionState() []string {
	return []string{
		string(RotateBitLockerKeysDeviceActionResultActionStateactive),
		string(RotateBitLockerKeysDeviceActionResultActionStatecanceled),
		string(RotateBitLockerKeysDeviceActionResultActionStatedone),
		string(RotateBitLockerKeysDeviceActionResultActionStatefailed),
		string(RotateBitLockerKeysDeviceActionResultActionStatenone),
		string(RotateBitLockerKeysDeviceActionResultActionStatenotSupported),
		string(RotateBitLockerKeysDeviceActionResultActionStatepending),
	}
}

func parseRotateBitLockerKeysDeviceActionResultActionState(input string) (*RotateBitLockerKeysDeviceActionResultActionState, error) {
	vals := map[string]RotateBitLockerKeysDeviceActionResultActionState{
		"active":       RotateBitLockerKeysDeviceActionResultActionStateactive,
		"canceled":     RotateBitLockerKeysDeviceActionResultActionStatecanceled,
		"done":         RotateBitLockerKeysDeviceActionResultActionStatedone,
		"failed":       RotateBitLockerKeysDeviceActionResultActionStatefailed,
		"none":         RotateBitLockerKeysDeviceActionResultActionStatenone,
		"notsupported": RotateBitLockerKeysDeviceActionResultActionStatenotSupported,
		"pending":      RotateBitLockerKeysDeviceActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RotateBitLockerKeysDeviceActionResultActionState(input)
	return &out, nil
}
