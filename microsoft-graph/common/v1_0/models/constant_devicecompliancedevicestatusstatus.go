package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceComplianceDeviceStatusStatus string

const (
	DeviceComplianceDeviceStatusStatuscompliant     DeviceComplianceDeviceStatusStatus = "Compliant"
	DeviceComplianceDeviceStatusStatusconflict      DeviceComplianceDeviceStatusStatus = "Conflict"
	DeviceComplianceDeviceStatusStatuserror         DeviceComplianceDeviceStatusStatus = "Error"
	DeviceComplianceDeviceStatusStatusnonCompliant  DeviceComplianceDeviceStatusStatus = "NonCompliant"
	DeviceComplianceDeviceStatusStatusnotApplicable DeviceComplianceDeviceStatusStatus = "NotApplicable"
	DeviceComplianceDeviceStatusStatusnotAssigned   DeviceComplianceDeviceStatusStatus = "NotAssigned"
	DeviceComplianceDeviceStatusStatusremediated    DeviceComplianceDeviceStatusStatus = "Remediated"
	DeviceComplianceDeviceStatusStatusunknown       DeviceComplianceDeviceStatusStatus = "Unknown"
)

func PossibleValuesForDeviceComplianceDeviceStatusStatus() []string {
	return []string{
		string(DeviceComplianceDeviceStatusStatuscompliant),
		string(DeviceComplianceDeviceStatusStatusconflict),
		string(DeviceComplianceDeviceStatusStatuserror),
		string(DeviceComplianceDeviceStatusStatusnonCompliant),
		string(DeviceComplianceDeviceStatusStatusnotApplicable),
		string(DeviceComplianceDeviceStatusStatusnotAssigned),
		string(DeviceComplianceDeviceStatusStatusremediated),
		string(DeviceComplianceDeviceStatusStatusunknown),
	}
}

func parseDeviceComplianceDeviceStatusStatus(input string) (*DeviceComplianceDeviceStatusStatus, error) {
	vals := map[string]DeviceComplianceDeviceStatusStatus{
		"compliant":     DeviceComplianceDeviceStatusStatuscompliant,
		"conflict":      DeviceComplianceDeviceStatusStatusconflict,
		"error":         DeviceComplianceDeviceStatusStatuserror,
		"noncompliant":  DeviceComplianceDeviceStatusStatusnonCompliant,
		"notapplicable": DeviceComplianceDeviceStatusStatusnotApplicable,
		"notassigned":   DeviceComplianceDeviceStatusStatusnotAssigned,
		"remediated":    DeviceComplianceDeviceStatusStatusremediated,
		"unknown":       DeviceComplianceDeviceStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceComplianceDeviceStatusStatus(input)
	return &out, nil
}
