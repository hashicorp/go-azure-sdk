package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthScriptPolicyStateRemediationState string

const (
	DeviceHealthScriptPolicyStateRemediationStateremediationFailed DeviceHealthScriptPolicyStateRemediationState = "RemediationFailed"
	DeviceHealthScriptPolicyStateRemediationStatescriptError       DeviceHealthScriptPolicyStateRemediationState = "ScriptError"
	DeviceHealthScriptPolicyStateRemediationStateskipped           DeviceHealthScriptPolicyStateRemediationState = "Skipped"
	DeviceHealthScriptPolicyStateRemediationStatesuccess           DeviceHealthScriptPolicyStateRemediationState = "Success"
	DeviceHealthScriptPolicyStateRemediationStateunknown           DeviceHealthScriptPolicyStateRemediationState = "Unknown"
)

func PossibleValuesForDeviceHealthScriptPolicyStateRemediationState() []string {
	return []string{
		string(DeviceHealthScriptPolicyStateRemediationStateremediationFailed),
		string(DeviceHealthScriptPolicyStateRemediationStatescriptError),
		string(DeviceHealthScriptPolicyStateRemediationStateskipped),
		string(DeviceHealthScriptPolicyStateRemediationStatesuccess),
		string(DeviceHealthScriptPolicyStateRemediationStateunknown),
	}
}

func parseDeviceHealthScriptPolicyStateRemediationState(input string) (*DeviceHealthScriptPolicyStateRemediationState, error) {
	vals := map[string]DeviceHealthScriptPolicyStateRemediationState{
		"remediationfailed": DeviceHealthScriptPolicyStateRemediationStateremediationFailed,
		"scripterror":       DeviceHealthScriptPolicyStateRemediationStatescriptError,
		"skipped":           DeviceHealthScriptPolicyStateRemediationStateskipped,
		"success":           DeviceHealthScriptPolicyStateRemediationStatesuccess,
		"unknown":           DeviceHealthScriptPolicyStateRemediationStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthScriptPolicyStateRemediationState(input)
	return &out, nil
}
