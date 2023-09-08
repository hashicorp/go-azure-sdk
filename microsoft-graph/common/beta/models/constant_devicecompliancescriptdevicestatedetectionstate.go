package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceScriptDeviceStateDetectionState string

const (
	DeviceComplianceScriptDeviceStateDetectionStatefail          DeviceComplianceScriptDeviceStateDetectionState = "Fail"
	DeviceComplianceScriptDeviceStateDetectionStatenotApplicable DeviceComplianceScriptDeviceStateDetectionState = "NotApplicable"
	DeviceComplianceScriptDeviceStateDetectionStatepending       DeviceComplianceScriptDeviceStateDetectionState = "Pending"
	DeviceComplianceScriptDeviceStateDetectionStatescriptError   DeviceComplianceScriptDeviceStateDetectionState = "ScriptError"
	DeviceComplianceScriptDeviceStateDetectionStatesuccess       DeviceComplianceScriptDeviceStateDetectionState = "Success"
	DeviceComplianceScriptDeviceStateDetectionStateunknown       DeviceComplianceScriptDeviceStateDetectionState = "Unknown"
)

func PossibleValuesForDeviceComplianceScriptDeviceStateDetectionState() []string {
	return []string{
		string(DeviceComplianceScriptDeviceStateDetectionStatefail),
		string(DeviceComplianceScriptDeviceStateDetectionStatenotApplicable),
		string(DeviceComplianceScriptDeviceStateDetectionStatepending),
		string(DeviceComplianceScriptDeviceStateDetectionStatescriptError),
		string(DeviceComplianceScriptDeviceStateDetectionStatesuccess),
		string(DeviceComplianceScriptDeviceStateDetectionStateunknown),
	}
}

func parseDeviceComplianceScriptDeviceStateDetectionState(input string) (*DeviceComplianceScriptDeviceStateDetectionState, error) {
	vals := map[string]DeviceComplianceScriptDeviceStateDetectionState{
		"fail":          DeviceComplianceScriptDeviceStateDetectionStatefail,
		"notapplicable": DeviceComplianceScriptDeviceStateDetectionStatenotApplicable,
		"pending":       DeviceComplianceScriptDeviceStateDetectionStatepending,
		"scripterror":   DeviceComplianceScriptDeviceStateDetectionStatescriptError,
		"success":       DeviceComplianceScriptDeviceStateDetectionStatesuccess,
		"unknown":       DeviceComplianceScriptDeviceStateDetectionStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceScriptDeviceStateDetectionState(input)
	return &out, nil
}
