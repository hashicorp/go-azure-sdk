package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LocateDeviceActionResultActionState string

const (
	LocateDeviceActionResultActionStateactive       LocateDeviceActionResultActionState = "Active"
	LocateDeviceActionResultActionStatecanceled     LocateDeviceActionResultActionState = "Canceled"
	LocateDeviceActionResultActionStatedone         LocateDeviceActionResultActionState = "Done"
	LocateDeviceActionResultActionStatefailed       LocateDeviceActionResultActionState = "Failed"
	LocateDeviceActionResultActionStatenone         LocateDeviceActionResultActionState = "None"
	LocateDeviceActionResultActionStatenotSupported LocateDeviceActionResultActionState = "NotSupported"
	LocateDeviceActionResultActionStatepending      LocateDeviceActionResultActionState = "Pending"
)

func PossibleValuesForLocateDeviceActionResultActionState() []string {
	return []string{
		string(LocateDeviceActionResultActionStateactive),
		string(LocateDeviceActionResultActionStatecanceled),
		string(LocateDeviceActionResultActionStatedone),
		string(LocateDeviceActionResultActionStatefailed),
		string(LocateDeviceActionResultActionStatenone),
		string(LocateDeviceActionResultActionStatenotSupported),
		string(LocateDeviceActionResultActionStatepending),
	}
}

func parseLocateDeviceActionResultActionState(input string) (*LocateDeviceActionResultActionState, error) {
	vals := map[string]LocateDeviceActionResultActionState{
		"active":       LocateDeviceActionResultActionStateactive,
		"canceled":     LocateDeviceActionResultActionStatecanceled,
		"done":         LocateDeviceActionResultActionStatedone,
		"failed":       LocateDeviceActionResultActionStatefailed,
		"none":         LocateDeviceActionResultActionStatenone,
		"notsupported": LocateDeviceActionResultActionStatenotSupported,
		"pending":      LocateDeviceActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := LocateDeviceActionResultActionState(input)
	return &out, nil
}
