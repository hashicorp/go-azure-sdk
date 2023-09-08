package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptDeviceStateRemediationState string

const (
	DeviceHealthScriptDeviceStateRemediationStateremediationFailed DeviceHealthScriptDeviceStateRemediationState = "RemediationFailed"
	DeviceHealthScriptDeviceStateRemediationStatescriptError       DeviceHealthScriptDeviceStateRemediationState = "ScriptError"
	DeviceHealthScriptDeviceStateRemediationStateskipped           DeviceHealthScriptDeviceStateRemediationState = "Skipped"
	DeviceHealthScriptDeviceStateRemediationStatesuccess           DeviceHealthScriptDeviceStateRemediationState = "Success"
	DeviceHealthScriptDeviceStateRemediationStateunknown           DeviceHealthScriptDeviceStateRemediationState = "Unknown"
)

func PossibleValuesForDeviceHealthScriptDeviceStateRemediationState() []string {
	return []string{
		string(DeviceHealthScriptDeviceStateRemediationStateremediationFailed),
		string(DeviceHealthScriptDeviceStateRemediationStatescriptError),
		string(DeviceHealthScriptDeviceStateRemediationStateskipped),
		string(DeviceHealthScriptDeviceStateRemediationStatesuccess),
		string(DeviceHealthScriptDeviceStateRemediationStateunknown),
	}
}

func parseDeviceHealthScriptDeviceStateRemediationState(input string) (*DeviceHealthScriptDeviceStateRemediationState, error) {
	vals := map[string]DeviceHealthScriptDeviceStateRemediationState{
		"remediationfailed": DeviceHealthScriptDeviceStateRemediationStateremediationFailed,
		"scripterror":       DeviceHealthScriptDeviceStateRemediationStatescriptError,
		"skipped":           DeviceHealthScriptDeviceStateRemediationStateskipped,
		"success":           DeviceHealthScriptDeviceStateRemediationStatesuccess,
		"unknown":           DeviceHealthScriptDeviceStateRemediationStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptDeviceStateRemediationState(input)
	return &out, nil
}
