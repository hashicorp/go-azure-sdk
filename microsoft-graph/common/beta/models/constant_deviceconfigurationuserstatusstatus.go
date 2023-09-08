package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationUserStatusStatus string

const (
	DeviceConfigurationUserStatusStatuscompliant     DeviceConfigurationUserStatusStatus = "Compliant"
	DeviceConfigurationUserStatusStatusconflict      DeviceConfigurationUserStatusStatus = "Conflict"
	DeviceConfigurationUserStatusStatuserror         DeviceConfigurationUserStatusStatus = "Error"
	DeviceConfigurationUserStatusStatusnonCompliant  DeviceConfigurationUserStatusStatus = "NonCompliant"
	DeviceConfigurationUserStatusStatusnotApplicable DeviceConfigurationUserStatusStatus = "NotApplicable"
	DeviceConfigurationUserStatusStatusnotAssigned   DeviceConfigurationUserStatusStatus = "NotAssigned"
	DeviceConfigurationUserStatusStatusremediated    DeviceConfigurationUserStatusStatus = "Remediated"
	DeviceConfigurationUserStatusStatusunknown       DeviceConfigurationUserStatusStatus = "Unknown"
)

func PossibleValuesForDeviceConfigurationUserStatusStatus() []string {
	return []string{
		string(DeviceConfigurationUserStatusStatuscompliant),
		string(DeviceConfigurationUserStatusStatusconflict),
		string(DeviceConfigurationUserStatusStatuserror),
		string(DeviceConfigurationUserStatusStatusnonCompliant),
		string(DeviceConfigurationUserStatusStatusnotApplicable),
		string(DeviceConfigurationUserStatusStatusnotAssigned),
		string(DeviceConfigurationUserStatusStatusremediated),
		string(DeviceConfigurationUserStatusStatusunknown),
	}
}

func parseDeviceConfigurationUserStatusStatus(input string) (*DeviceConfigurationUserStatusStatus, error) {
	vals := map[string]DeviceConfigurationUserStatusStatus{
		"compliant":     DeviceConfigurationUserStatusStatuscompliant,
		"conflict":      DeviceConfigurationUserStatusStatusconflict,
		"error":         DeviceConfigurationUserStatusStatuserror,
		"noncompliant":  DeviceConfigurationUserStatusStatusnonCompliant,
		"notapplicable": DeviceConfigurationUserStatusStatusnotApplicable,
		"notassigned":   DeviceConfigurationUserStatusStatusnotAssigned,
		"remediated":    DeviceConfigurationUserStatusStatusremediated,
		"unknown":       DeviceConfigurationUserStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationUserStatusStatus(input)
	return &out, nil
}
