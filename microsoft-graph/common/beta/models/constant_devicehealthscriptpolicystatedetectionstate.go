package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptPolicyStateDetectionState string

const (
	DeviceHealthScriptPolicyStateDetectionStatefail          DeviceHealthScriptPolicyStateDetectionState = "Fail"
	DeviceHealthScriptPolicyStateDetectionStatenotApplicable DeviceHealthScriptPolicyStateDetectionState = "NotApplicable"
	DeviceHealthScriptPolicyStateDetectionStatepending       DeviceHealthScriptPolicyStateDetectionState = "Pending"
	DeviceHealthScriptPolicyStateDetectionStatescriptError   DeviceHealthScriptPolicyStateDetectionState = "ScriptError"
	DeviceHealthScriptPolicyStateDetectionStatesuccess       DeviceHealthScriptPolicyStateDetectionState = "Success"
	DeviceHealthScriptPolicyStateDetectionStateunknown       DeviceHealthScriptPolicyStateDetectionState = "Unknown"
)

func PossibleValuesForDeviceHealthScriptPolicyStateDetectionState() []string {
	return []string{
		string(DeviceHealthScriptPolicyStateDetectionStatefail),
		string(DeviceHealthScriptPolicyStateDetectionStatenotApplicable),
		string(DeviceHealthScriptPolicyStateDetectionStatepending),
		string(DeviceHealthScriptPolicyStateDetectionStatescriptError),
		string(DeviceHealthScriptPolicyStateDetectionStatesuccess),
		string(DeviceHealthScriptPolicyStateDetectionStateunknown),
	}
}

func parseDeviceHealthScriptPolicyStateDetectionState(input string) (*DeviceHealthScriptPolicyStateDetectionState, error) {
	vals := map[string]DeviceHealthScriptPolicyStateDetectionState{
		"fail":          DeviceHealthScriptPolicyStateDetectionStatefail,
		"notapplicable": DeviceHealthScriptPolicyStateDetectionStatenotApplicable,
		"pending":       DeviceHealthScriptPolicyStateDetectionStatepending,
		"scripterror":   DeviceHealthScriptPolicyStateDetectionStatescriptError,
		"success":       DeviceHealthScriptPolicyStateDetectionStatesuccess,
		"unknown":       DeviceHealthScriptPolicyStateDetectionStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptPolicyStateDetectionState(input)
	return &out, nil
}
