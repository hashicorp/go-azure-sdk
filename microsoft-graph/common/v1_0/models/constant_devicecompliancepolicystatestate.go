package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyStateState string

const (
	DeviceCompliancePolicyStateStatecompliant     DeviceCompliancePolicyStateState = "Compliant"
	DeviceCompliancePolicyStateStateconflict      DeviceCompliancePolicyStateState = "Conflict"
	DeviceCompliancePolicyStateStateerror         DeviceCompliancePolicyStateState = "Error"
	DeviceCompliancePolicyStateStatenonCompliant  DeviceCompliancePolicyStateState = "NonCompliant"
	DeviceCompliancePolicyStateStatenotApplicable DeviceCompliancePolicyStateState = "NotApplicable"
	DeviceCompliancePolicyStateStatenotAssigned   DeviceCompliancePolicyStateState = "NotAssigned"
	DeviceCompliancePolicyStateStateremediated    DeviceCompliancePolicyStateState = "Remediated"
	DeviceCompliancePolicyStateStateunknown       DeviceCompliancePolicyStateState = "Unknown"
)

func PossibleValuesForDeviceCompliancePolicyStateState() []string {
	return []string{
		string(DeviceCompliancePolicyStateStatecompliant),
		string(DeviceCompliancePolicyStateStateconflict),
		string(DeviceCompliancePolicyStateStateerror),
		string(DeviceCompliancePolicyStateStatenonCompliant),
		string(DeviceCompliancePolicyStateStatenotApplicable),
		string(DeviceCompliancePolicyStateStatenotAssigned),
		string(DeviceCompliancePolicyStateStateremediated),
		string(DeviceCompliancePolicyStateStateunknown),
	}
}

func parseDeviceCompliancePolicyStateState(input string) (*DeviceCompliancePolicyStateState, error) {
	vals := map[string]DeviceCompliancePolicyStateState{
		"compliant":     DeviceCompliancePolicyStateStatecompliant,
		"conflict":      DeviceCompliancePolicyStateStateconflict,
		"error":         DeviceCompliancePolicyStateStateerror,
		"noncompliant":  DeviceCompliancePolicyStateStatenonCompliant,
		"notapplicable": DeviceCompliancePolicyStateStatenotApplicable,
		"notassigned":   DeviceCompliancePolicyStateStatenotAssigned,
		"remediated":    DeviceCompliancePolicyStateStateremediated,
		"unknown":       DeviceCompliancePolicyStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyStateState(input)
	return &out, nil
}
