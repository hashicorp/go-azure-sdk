package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceStateDetectionState string

const (
	DeviceHealthScriptDeviceStateDetectionStatefail          DeviceHealthScriptDeviceStateDetectionState = "Fail"
	DeviceHealthScriptDeviceStateDetectionStatenotApplicable DeviceHealthScriptDeviceStateDetectionState = "NotApplicable"
	DeviceHealthScriptDeviceStateDetectionStatepending       DeviceHealthScriptDeviceStateDetectionState = "Pending"
	DeviceHealthScriptDeviceStateDetectionStatescriptError   DeviceHealthScriptDeviceStateDetectionState = "ScriptError"
	DeviceHealthScriptDeviceStateDetectionStatesuccess       DeviceHealthScriptDeviceStateDetectionState = "Success"
	DeviceHealthScriptDeviceStateDetectionStateunknown       DeviceHealthScriptDeviceStateDetectionState = "Unknown"
)

func PossibleValuesForDeviceHealthScriptDeviceStateDetectionState() []string {
	return []string{
		string(DeviceHealthScriptDeviceStateDetectionStatefail),
		string(DeviceHealthScriptDeviceStateDetectionStatenotApplicable),
		string(DeviceHealthScriptDeviceStateDetectionStatepending),
		string(DeviceHealthScriptDeviceStateDetectionStatescriptError),
		string(DeviceHealthScriptDeviceStateDetectionStatesuccess),
		string(DeviceHealthScriptDeviceStateDetectionStateunknown),
	}
}

func parseDeviceHealthScriptDeviceStateDetectionState(input string) (*DeviceHealthScriptDeviceStateDetectionState, error) {
	vals := map[string]DeviceHealthScriptDeviceStateDetectionState{
		"fail":          DeviceHealthScriptDeviceStateDetectionStatefail,
		"notapplicable": DeviceHealthScriptDeviceStateDetectionStatenotApplicable,
		"pending":       DeviceHealthScriptDeviceStateDetectionStatepending,
		"scripterror":   DeviceHealthScriptDeviceStateDetectionStatescriptError,
		"success":       DeviceHealthScriptDeviceStateDetectionStatesuccess,
		"unknown":       DeviceHealthScriptDeviceStateDetectionStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptDeviceStateDetectionState(input)
	return &out, nil
}
