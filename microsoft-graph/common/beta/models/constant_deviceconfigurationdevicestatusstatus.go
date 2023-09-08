package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationDeviceStatusStatus string

const (
	DeviceConfigurationDeviceStatusStatuscompliant     DeviceConfigurationDeviceStatusStatus = "Compliant"
	DeviceConfigurationDeviceStatusStatusconflict      DeviceConfigurationDeviceStatusStatus = "Conflict"
	DeviceConfigurationDeviceStatusStatuserror         DeviceConfigurationDeviceStatusStatus = "Error"
	DeviceConfigurationDeviceStatusStatusnonCompliant  DeviceConfigurationDeviceStatusStatus = "NonCompliant"
	DeviceConfigurationDeviceStatusStatusnotApplicable DeviceConfigurationDeviceStatusStatus = "NotApplicable"
	DeviceConfigurationDeviceStatusStatusnotAssigned   DeviceConfigurationDeviceStatusStatus = "NotAssigned"
	DeviceConfigurationDeviceStatusStatusremediated    DeviceConfigurationDeviceStatusStatus = "Remediated"
	DeviceConfigurationDeviceStatusStatusunknown       DeviceConfigurationDeviceStatusStatus = "Unknown"
)

func PossibleValuesForDeviceConfigurationDeviceStatusStatus() []string {
	return []string{
		string(DeviceConfigurationDeviceStatusStatuscompliant),
		string(DeviceConfigurationDeviceStatusStatusconflict),
		string(DeviceConfigurationDeviceStatusStatuserror),
		string(DeviceConfigurationDeviceStatusStatusnonCompliant),
		string(DeviceConfigurationDeviceStatusStatusnotApplicable),
		string(DeviceConfigurationDeviceStatusStatusnotAssigned),
		string(DeviceConfigurationDeviceStatusStatusremediated),
		string(DeviceConfigurationDeviceStatusStatusunknown),
	}
}

func parseDeviceConfigurationDeviceStatusStatus(input string) (*DeviceConfigurationDeviceStatusStatus, error) {
	vals := map[string]DeviceConfigurationDeviceStatusStatus{
		"compliant":     DeviceConfigurationDeviceStatusStatuscompliant,
		"conflict":      DeviceConfigurationDeviceStatusStatusconflict,
		"error":         DeviceConfigurationDeviceStatusStatuserror,
		"noncompliant":  DeviceConfigurationDeviceStatusStatusnonCompliant,
		"notapplicable": DeviceConfigurationDeviceStatusStatusnotApplicable,
		"notassigned":   DeviceConfigurationDeviceStatusStatusnotAssigned,
		"remediated":    DeviceConfigurationDeviceStatusStatusremediated,
		"unknown":       DeviceConfigurationDeviceStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationDeviceStatusStatus(input)
	return &out, nil
}
