package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceUserStatusStatus string

const (
	DeviceComplianceUserStatusStatuscompliant     DeviceComplianceUserStatusStatus = "Compliant"
	DeviceComplianceUserStatusStatusconflict      DeviceComplianceUserStatusStatus = "Conflict"
	DeviceComplianceUserStatusStatuserror         DeviceComplianceUserStatusStatus = "Error"
	DeviceComplianceUserStatusStatusnonCompliant  DeviceComplianceUserStatusStatus = "NonCompliant"
	DeviceComplianceUserStatusStatusnotApplicable DeviceComplianceUserStatusStatus = "NotApplicable"
	DeviceComplianceUserStatusStatusnotAssigned   DeviceComplianceUserStatusStatus = "NotAssigned"
	DeviceComplianceUserStatusStatusremediated    DeviceComplianceUserStatusStatus = "Remediated"
	DeviceComplianceUserStatusStatusunknown       DeviceComplianceUserStatusStatus = "Unknown"
)

func PossibleValuesForDeviceComplianceUserStatusStatus() []string {
	return []string{
		string(DeviceComplianceUserStatusStatuscompliant),
		string(DeviceComplianceUserStatusStatusconflict),
		string(DeviceComplianceUserStatusStatuserror),
		string(DeviceComplianceUserStatusStatusnonCompliant),
		string(DeviceComplianceUserStatusStatusnotApplicable),
		string(DeviceComplianceUserStatusStatusnotAssigned),
		string(DeviceComplianceUserStatusStatusremediated),
		string(DeviceComplianceUserStatusStatusunknown),
	}
}

func parseDeviceComplianceUserStatusStatus(input string) (*DeviceComplianceUserStatusStatus, error) {
	vals := map[string]DeviceComplianceUserStatusStatus{
		"compliant":     DeviceComplianceUserStatusStatuscompliant,
		"conflict":      DeviceComplianceUserStatusStatusconflict,
		"error":         DeviceComplianceUserStatusStatuserror,
		"noncompliant":  DeviceComplianceUserStatusStatusnonCompliant,
		"notapplicable": DeviceComplianceUserStatusStatusnotApplicable,
		"notassigned":   DeviceComplianceUserStatusStatusnotAssigned,
		"remediated":    DeviceComplianceUserStatusStatusremediated,
		"unknown":       DeviceComplianceUserStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceUserStatusStatus(input)
	return &out, nil
}
