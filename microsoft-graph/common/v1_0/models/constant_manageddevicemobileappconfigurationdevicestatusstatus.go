package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationDeviceStatusStatus string

const (
	ManagedDeviceMobileAppConfigurationDeviceStatusStatuscompliant     ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "Compliant"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatusconflict      ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "Conflict"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatuserror         ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "Error"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatusnonCompliant  ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "NonCompliant"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatusnotApplicable ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "NotApplicable"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatusnotAssigned   ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "NotAssigned"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatusremediated    ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "Remediated"
	ManagedDeviceMobileAppConfigurationDeviceStatusStatusunknown       ManagedDeviceMobileAppConfigurationDeviceStatusStatus = "Unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationDeviceStatusStatus() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatuscompliant),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatusconflict),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatuserror),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatusnonCompliant),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatusnotApplicable),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatusnotAssigned),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatusremediated),
		string(ManagedDeviceMobileAppConfigurationDeviceStatusStatusunknown),
	}
}

func parseManagedDeviceMobileAppConfigurationDeviceStatusStatus(input string) (*ManagedDeviceMobileAppConfigurationDeviceStatusStatus, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationDeviceStatusStatus{
		"compliant":     ManagedDeviceMobileAppConfigurationDeviceStatusStatuscompliant,
		"conflict":      ManagedDeviceMobileAppConfigurationDeviceStatusStatusconflict,
		"error":         ManagedDeviceMobileAppConfigurationDeviceStatusStatuserror,
		"noncompliant":  ManagedDeviceMobileAppConfigurationDeviceStatusStatusnonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationDeviceStatusStatusnotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationDeviceStatusStatusnotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationDeviceStatusStatusremediated,
		"unknown":       ManagedDeviceMobileAppConfigurationDeviceStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationDeviceStatusStatus(input)
	return &out, nil
}
