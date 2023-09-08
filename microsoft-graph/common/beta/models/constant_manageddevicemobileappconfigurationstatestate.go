package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationStateState string

const (
	ManagedDeviceMobileAppConfigurationStateStatecompliant     ManagedDeviceMobileAppConfigurationStateState = "Compliant"
	ManagedDeviceMobileAppConfigurationStateStateconflict      ManagedDeviceMobileAppConfigurationStateState = "Conflict"
	ManagedDeviceMobileAppConfigurationStateStateerror         ManagedDeviceMobileAppConfigurationStateState = "Error"
	ManagedDeviceMobileAppConfigurationStateStatenonCompliant  ManagedDeviceMobileAppConfigurationStateState = "NonCompliant"
	ManagedDeviceMobileAppConfigurationStateStatenotApplicable ManagedDeviceMobileAppConfigurationStateState = "NotApplicable"
	ManagedDeviceMobileAppConfigurationStateStatenotAssigned   ManagedDeviceMobileAppConfigurationStateState = "NotAssigned"
	ManagedDeviceMobileAppConfigurationStateStateremediated    ManagedDeviceMobileAppConfigurationStateState = "Remediated"
	ManagedDeviceMobileAppConfigurationStateStateunknown       ManagedDeviceMobileAppConfigurationStateState = "Unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationStateState() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationStateStatecompliant),
		string(ManagedDeviceMobileAppConfigurationStateStateconflict),
		string(ManagedDeviceMobileAppConfigurationStateStateerror),
		string(ManagedDeviceMobileAppConfigurationStateStatenonCompliant),
		string(ManagedDeviceMobileAppConfigurationStateStatenotApplicable),
		string(ManagedDeviceMobileAppConfigurationStateStatenotAssigned),
		string(ManagedDeviceMobileAppConfigurationStateStateremediated),
		string(ManagedDeviceMobileAppConfigurationStateStateunknown),
	}
}

func parseManagedDeviceMobileAppConfigurationStateState(input string) (*ManagedDeviceMobileAppConfigurationStateState, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationStateState{
		"compliant":     ManagedDeviceMobileAppConfigurationStateStatecompliant,
		"conflict":      ManagedDeviceMobileAppConfigurationStateStateconflict,
		"error":         ManagedDeviceMobileAppConfigurationStateStateerror,
		"noncompliant":  ManagedDeviceMobileAppConfigurationStateStatenonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationStateStatenotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationStateStatenotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationStateStateremediated,
		"unknown":       ManagedDeviceMobileAppConfigurationStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationStateState(input)
	return &out, nil
}
