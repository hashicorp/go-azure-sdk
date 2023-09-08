package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateState string

const (
	DeviceCompliancePolicySettingStateStatecompliant     DeviceCompliancePolicySettingStateState = "Compliant"
	DeviceCompliancePolicySettingStateStateconflict      DeviceCompliancePolicySettingStateState = "Conflict"
	DeviceCompliancePolicySettingStateStateerror         DeviceCompliancePolicySettingStateState = "Error"
	DeviceCompliancePolicySettingStateStatenonCompliant  DeviceCompliancePolicySettingStateState = "NonCompliant"
	DeviceCompliancePolicySettingStateStatenotApplicable DeviceCompliancePolicySettingStateState = "NotApplicable"
	DeviceCompliancePolicySettingStateStatenotAssigned   DeviceCompliancePolicySettingStateState = "NotAssigned"
	DeviceCompliancePolicySettingStateStateremediated    DeviceCompliancePolicySettingStateState = "Remediated"
	DeviceCompliancePolicySettingStateStateunknown       DeviceCompliancePolicySettingStateState = "Unknown"
)

func PossibleValuesForDeviceCompliancePolicySettingStateState() []string {
	return []string{
		string(DeviceCompliancePolicySettingStateStatecompliant),
		string(DeviceCompliancePolicySettingStateStateconflict),
		string(DeviceCompliancePolicySettingStateStateerror),
		string(DeviceCompliancePolicySettingStateStatenonCompliant),
		string(DeviceCompliancePolicySettingStateStatenotApplicable),
		string(DeviceCompliancePolicySettingStateStatenotAssigned),
		string(DeviceCompliancePolicySettingStateStateremediated),
		string(DeviceCompliancePolicySettingStateStateunknown),
	}
}

func parseDeviceCompliancePolicySettingStateState(input string) (*DeviceCompliancePolicySettingStateState, error) {
	vals := map[string]DeviceCompliancePolicySettingStateState{
		"compliant":     DeviceCompliancePolicySettingStateStatecompliant,
		"conflict":      DeviceCompliancePolicySettingStateStateconflict,
		"error":         DeviceCompliancePolicySettingStateStateerror,
		"noncompliant":  DeviceCompliancePolicySettingStateStatenonCompliant,
		"notapplicable": DeviceCompliancePolicySettingStateStatenotApplicable,
		"notassigned":   DeviceCompliancePolicySettingStateStatenotAssigned,
		"remediated":    DeviceCompliancePolicySettingStateStateremediated,
		"unknown":       DeviceCompliancePolicySettingStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicySettingStateState(input)
	return &out, nil
}
