package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceActionResultActionState string

const (
	DeviceActionResultActionStateactive       DeviceActionResultActionState = "Active"
	DeviceActionResultActionStatecanceled     DeviceActionResultActionState = "Canceled"
	DeviceActionResultActionStatedone         DeviceActionResultActionState = "Done"
	DeviceActionResultActionStatefailed       DeviceActionResultActionState = "Failed"
	DeviceActionResultActionStatenone         DeviceActionResultActionState = "None"
	DeviceActionResultActionStatenotSupported DeviceActionResultActionState = "NotSupported"
	DeviceActionResultActionStatepending      DeviceActionResultActionState = "Pending"
)

func PossibleValuesForDeviceActionResultActionState() []string {
	return []string{
		string(DeviceActionResultActionStateactive),
		string(DeviceActionResultActionStatecanceled),
		string(DeviceActionResultActionStatedone),
		string(DeviceActionResultActionStatefailed),
		string(DeviceActionResultActionStatenone),
		string(DeviceActionResultActionStatenotSupported),
		string(DeviceActionResultActionStatepending),
	}
}

func parseDeviceActionResultActionState(input string) (*DeviceActionResultActionState, error) {
	vals := map[string]DeviceActionResultActionState{
		"active":       DeviceActionResultActionStateactive,
		"canceled":     DeviceActionResultActionStatecanceled,
		"done":         DeviceActionResultActionStatedone,
		"failed":       DeviceActionResultActionStatefailed,
		"none":         DeviceActionResultActionStatenone,
		"notsupported": DeviceActionResultActionStatenotSupported,
		"pending":      DeviceActionResultActionStatepending,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceActionResultActionState(input)
	return &out, nil
}
