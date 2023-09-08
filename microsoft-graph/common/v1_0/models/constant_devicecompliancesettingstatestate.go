package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceSettingStateState string

const (
	DeviceComplianceSettingStateStatecompliant     DeviceComplianceSettingStateState = "Compliant"
	DeviceComplianceSettingStateStateconflict      DeviceComplianceSettingStateState = "Conflict"
	DeviceComplianceSettingStateStateerror         DeviceComplianceSettingStateState = "Error"
	DeviceComplianceSettingStateStatenonCompliant  DeviceComplianceSettingStateState = "NonCompliant"
	DeviceComplianceSettingStateStatenotApplicable DeviceComplianceSettingStateState = "NotApplicable"
	DeviceComplianceSettingStateStatenotAssigned   DeviceComplianceSettingStateState = "NotAssigned"
	DeviceComplianceSettingStateStateremediated    DeviceComplianceSettingStateState = "Remediated"
	DeviceComplianceSettingStateStateunknown       DeviceComplianceSettingStateState = "Unknown"
)

func PossibleValuesForDeviceComplianceSettingStateState() []string {
	return []string{
		string(DeviceComplianceSettingStateStatecompliant),
		string(DeviceComplianceSettingStateStateconflict),
		string(DeviceComplianceSettingStateStateerror),
		string(DeviceComplianceSettingStateStatenonCompliant),
		string(DeviceComplianceSettingStateStatenotApplicable),
		string(DeviceComplianceSettingStateStatenotAssigned),
		string(DeviceComplianceSettingStateStateremediated),
		string(DeviceComplianceSettingStateStateunknown),
	}
}

func parseDeviceComplianceSettingStateState(input string) (*DeviceComplianceSettingStateState, error) {
	vals := map[string]DeviceComplianceSettingStateState{
		"compliant":     DeviceComplianceSettingStateStatecompliant,
		"conflict":      DeviceComplianceSettingStateStateconflict,
		"error":         DeviceComplianceSettingStateStateerror,
		"noncompliant":  DeviceComplianceSettingStateStatenonCompliant,
		"notapplicable": DeviceComplianceSettingStateStatenotApplicable,
		"notassigned":   DeviceComplianceSettingStateStatenotAssigned,
		"remediated":    DeviceComplianceSettingStateStateremediated,
		"unknown":       DeviceComplianceSettingStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceSettingStateState(input)
	return &out, nil
}
