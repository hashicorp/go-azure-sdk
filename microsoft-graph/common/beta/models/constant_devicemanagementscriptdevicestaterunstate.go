package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptDeviceStateRunState string

const (
	DeviceManagementScriptDeviceStateRunStatefail          DeviceManagementScriptDeviceStateRunState = "Fail"
	DeviceManagementScriptDeviceStateRunStatenotApplicable DeviceManagementScriptDeviceStateRunState = "NotApplicable"
	DeviceManagementScriptDeviceStateRunStatepending       DeviceManagementScriptDeviceStateRunState = "Pending"
	DeviceManagementScriptDeviceStateRunStatescriptError   DeviceManagementScriptDeviceStateRunState = "ScriptError"
	DeviceManagementScriptDeviceStateRunStatesuccess       DeviceManagementScriptDeviceStateRunState = "Success"
	DeviceManagementScriptDeviceStateRunStateunknown       DeviceManagementScriptDeviceStateRunState = "Unknown"
)

func PossibleValuesForDeviceManagementScriptDeviceStateRunState() []string {
	return []string{
		string(DeviceManagementScriptDeviceStateRunStatefail),
		string(DeviceManagementScriptDeviceStateRunStatenotApplicable),
		string(DeviceManagementScriptDeviceStateRunStatepending),
		string(DeviceManagementScriptDeviceStateRunStatescriptError),
		string(DeviceManagementScriptDeviceStateRunStatesuccess),
		string(DeviceManagementScriptDeviceStateRunStateunknown),
	}
}

func parseDeviceManagementScriptDeviceStateRunState(input string) (*DeviceManagementScriptDeviceStateRunState, error) {
	vals := map[string]DeviceManagementScriptDeviceStateRunState{
		"fail":          DeviceManagementScriptDeviceStateRunStatefail,
		"notapplicable": DeviceManagementScriptDeviceStateRunStatenotApplicable,
		"pending":       DeviceManagementScriptDeviceStateRunStatepending,
		"scripterror":   DeviceManagementScriptDeviceStateRunStatescriptError,
		"success":       DeviceManagementScriptDeviceStateRunStatesuccess,
		"unknown":       DeviceManagementScriptDeviceStateRunStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptDeviceStateRunState(input)
	return &out, nil
}
