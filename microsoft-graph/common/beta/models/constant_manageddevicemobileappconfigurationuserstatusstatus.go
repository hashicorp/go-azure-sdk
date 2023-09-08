package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationUserStatusStatus string

const (
	ManagedDeviceMobileAppConfigurationUserStatusStatuscompliant     ManagedDeviceMobileAppConfigurationUserStatusStatus = "Compliant"
	ManagedDeviceMobileAppConfigurationUserStatusStatusconflict      ManagedDeviceMobileAppConfigurationUserStatusStatus = "Conflict"
	ManagedDeviceMobileAppConfigurationUserStatusStatuserror         ManagedDeviceMobileAppConfigurationUserStatusStatus = "Error"
	ManagedDeviceMobileAppConfigurationUserStatusStatusnonCompliant  ManagedDeviceMobileAppConfigurationUserStatusStatus = "NonCompliant"
	ManagedDeviceMobileAppConfigurationUserStatusStatusnotApplicable ManagedDeviceMobileAppConfigurationUserStatusStatus = "NotApplicable"
	ManagedDeviceMobileAppConfigurationUserStatusStatusnotAssigned   ManagedDeviceMobileAppConfigurationUserStatusStatus = "NotAssigned"
	ManagedDeviceMobileAppConfigurationUserStatusStatusremediated    ManagedDeviceMobileAppConfigurationUserStatusStatus = "Remediated"
	ManagedDeviceMobileAppConfigurationUserStatusStatusunknown       ManagedDeviceMobileAppConfigurationUserStatusStatus = "Unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationUserStatusStatus() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationUserStatusStatuscompliant),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatusconflict),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatuserror),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatusnonCompliant),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatusnotApplicable),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatusnotAssigned),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatusremediated),
		string(ManagedDeviceMobileAppConfigurationUserStatusStatusunknown),
	}
}

func parseManagedDeviceMobileAppConfigurationUserStatusStatus(input string) (*ManagedDeviceMobileAppConfigurationUserStatusStatus, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationUserStatusStatus{
		"compliant":     ManagedDeviceMobileAppConfigurationUserStatusStatuscompliant,
		"conflict":      ManagedDeviceMobileAppConfigurationUserStatusStatusconflict,
		"error":         ManagedDeviceMobileAppConfigurationUserStatusStatuserror,
		"noncompliant":  ManagedDeviceMobileAppConfigurationUserStatusStatusnonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationUserStatusStatusnotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationUserStatusStatusnotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationUserStatusStatusremediated,
		"unknown":       ManagedDeviceMobileAppConfigurationUserStatusStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationUserStatusStatus(input)
	return &out, nil
}
