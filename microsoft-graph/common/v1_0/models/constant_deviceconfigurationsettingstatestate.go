package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationSettingStateState string

const (
	DeviceConfigurationSettingStateStatecompliant     DeviceConfigurationSettingStateState = "Compliant"
	DeviceConfigurationSettingStateStateconflict      DeviceConfigurationSettingStateState = "Conflict"
	DeviceConfigurationSettingStateStateerror         DeviceConfigurationSettingStateState = "Error"
	DeviceConfigurationSettingStateStatenonCompliant  DeviceConfigurationSettingStateState = "NonCompliant"
	DeviceConfigurationSettingStateStatenotApplicable DeviceConfigurationSettingStateState = "NotApplicable"
	DeviceConfigurationSettingStateStatenotAssigned   DeviceConfigurationSettingStateState = "NotAssigned"
	DeviceConfigurationSettingStateStateremediated    DeviceConfigurationSettingStateState = "Remediated"
	DeviceConfigurationSettingStateStateunknown       DeviceConfigurationSettingStateState = "Unknown"
)

func PossibleValuesForDeviceConfigurationSettingStateState() []string {
	return []string{
		string(DeviceConfigurationSettingStateStatecompliant),
		string(DeviceConfigurationSettingStateStateconflict),
		string(DeviceConfigurationSettingStateStateerror),
		string(DeviceConfigurationSettingStateStatenonCompliant),
		string(DeviceConfigurationSettingStateStatenotApplicable),
		string(DeviceConfigurationSettingStateStatenotAssigned),
		string(DeviceConfigurationSettingStateStateremediated),
		string(DeviceConfigurationSettingStateStateunknown),
	}
}

func parseDeviceConfigurationSettingStateState(input string) (*DeviceConfigurationSettingStateState, error) {
	vals := map[string]DeviceConfigurationSettingStateState{
		"compliant":     DeviceConfigurationSettingStateStatecompliant,
		"conflict":      DeviceConfigurationSettingStateStateconflict,
		"error":         DeviceConfigurationSettingStateStateerror,
		"noncompliant":  DeviceConfigurationSettingStateStatenonCompliant,
		"notapplicable": DeviceConfigurationSettingStateStatenotApplicable,
		"notassigned":   DeviceConfigurationSettingStateStatenotAssigned,
		"remediated":    DeviceConfigurationSettingStateStateremediated,
		"unknown":       DeviceConfigurationSettingStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationSettingStateState(input)
	return &out, nil
}
