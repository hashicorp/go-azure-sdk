package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementIntentUserStateState string

const (
	DeviceManagementIntentUserStateStatecompliant     DeviceManagementIntentUserStateState = "Compliant"
	DeviceManagementIntentUserStateStateconflict      DeviceManagementIntentUserStateState = "Conflict"
	DeviceManagementIntentUserStateStateerror         DeviceManagementIntentUserStateState = "Error"
	DeviceManagementIntentUserStateStatenonCompliant  DeviceManagementIntentUserStateState = "NonCompliant"
	DeviceManagementIntentUserStateStatenotApplicable DeviceManagementIntentUserStateState = "NotApplicable"
	DeviceManagementIntentUserStateStatenotAssigned   DeviceManagementIntentUserStateState = "NotAssigned"
	DeviceManagementIntentUserStateStateremediated    DeviceManagementIntentUserStateState = "Remediated"
	DeviceManagementIntentUserStateStateunknown       DeviceManagementIntentUserStateState = "Unknown"
)

func PossibleValuesForDeviceManagementIntentUserStateState() []string {
	return []string{
		string(DeviceManagementIntentUserStateStatecompliant),
		string(DeviceManagementIntentUserStateStateconflict),
		string(DeviceManagementIntentUserStateStateerror),
		string(DeviceManagementIntentUserStateStatenonCompliant),
		string(DeviceManagementIntentUserStateStatenotApplicable),
		string(DeviceManagementIntentUserStateStatenotAssigned),
		string(DeviceManagementIntentUserStateStateremediated),
		string(DeviceManagementIntentUserStateStateunknown),
	}
}

func parseDeviceManagementIntentUserStateState(input string) (*DeviceManagementIntentUserStateState, error) {
	vals := map[string]DeviceManagementIntentUserStateState{
		"compliant":     DeviceManagementIntentUserStateStatecompliant,
		"conflict":      DeviceManagementIntentUserStateStateconflict,
		"error":         DeviceManagementIntentUserStateStateerror,
		"noncompliant":  DeviceManagementIntentUserStateStatenonCompliant,
		"notapplicable": DeviceManagementIntentUserStateStatenotApplicable,
		"notassigned":   DeviceManagementIntentUserStateStatenotAssigned,
		"remediated":    DeviceManagementIntentUserStateStateremediated,
		"unknown":       DeviceManagementIntentUserStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementIntentUserStateState(input)
	return &out, nil
}
