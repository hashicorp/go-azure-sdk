package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationStateState string

const (
	DeviceConfigurationStateStatecompliant     DeviceConfigurationStateState = "Compliant"
	DeviceConfigurationStateStateconflict      DeviceConfigurationStateState = "Conflict"
	DeviceConfigurationStateStateerror         DeviceConfigurationStateState = "Error"
	DeviceConfigurationStateStatenonCompliant  DeviceConfigurationStateState = "NonCompliant"
	DeviceConfigurationStateStatenotApplicable DeviceConfigurationStateState = "NotApplicable"
	DeviceConfigurationStateStatenotAssigned   DeviceConfigurationStateState = "NotAssigned"
	DeviceConfigurationStateStateremediated    DeviceConfigurationStateState = "Remediated"
	DeviceConfigurationStateStateunknown       DeviceConfigurationStateState = "Unknown"
)

func PossibleValuesForDeviceConfigurationStateState() []string {
	return []string{
		string(DeviceConfigurationStateStatecompliant),
		string(DeviceConfigurationStateStateconflict),
		string(DeviceConfigurationStateStateerror),
		string(DeviceConfigurationStateStatenonCompliant),
		string(DeviceConfigurationStateStatenotApplicable),
		string(DeviceConfigurationStateStatenotAssigned),
		string(DeviceConfigurationStateStateremediated),
		string(DeviceConfigurationStateStateunknown),
	}
}

func parseDeviceConfigurationStateState(input string) (*DeviceConfigurationStateState, error) {
	vals := map[string]DeviceConfigurationStateState{
		"compliant":     DeviceConfigurationStateStatecompliant,
		"conflict":      DeviceConfigurationStateStateconflict,
		"error":         DeviceConfigurationStateStateerror,
		"noncompliant":  DeviceConfigurationStateStatenonCompliant,
		"notapplicable": DeviceConfigurationStateStatenotApplicable,
		"notassigned":   DeviceConfigurationStateStatenotAssigned,
		"remediated":    DeviceConfigurationStateStateremediated,
		"unknown":       DeviceConfigurationStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationStateState(input)
	return &out, nil
}
