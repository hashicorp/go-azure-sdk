package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationSettingStateState string

const (
	ManagedDeviceMobileAppConfigurationSettingStateStatecompliant     ManagedDeviceMobileAppConfigurationSettingStateState = "Compliant"
	ManagedDeviceMobileAppConfigurationSettingStateStateconflict      ManagedDeviceMobileAppConfigurationSettingStateState = "Conflict"
	ManagedDeviceMobileAppConfigurationSettingStateStateerror         ManagedDeviceMobileAppConfigurationSettingStateState = "Error"
	ManagedDeviceMobileAppConfigurationSettingStateStatenonCompliant  ManagedDeviceMobileAppConfigurationSettingStateState = "NonCompliant"
	ManagedDeviceMobileAppConfigurationSettingStateStatenotApplicable ManagedDeviceMobileAppConfigurationSettingStateState = "NotApplicable"
	ManagedDeviceMobileAppConfigurationSettingStateStatenotAssigned   ManagedDeviceMobileAppConfigurationSettingStateState = "NotAssigned"
	ManagedDeviceMobileAppConfigurationSettingStateStateremediated    ManagedDeviceMobileAppConfigurationSettingStateState = "Remediated"
	ManagedDeviceMobileAppConfigurationSettingStateStateunknown       ManagedDeviceMobileAppConfigurationSettingStateState = "Unknown"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationSettingStateState() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationSettingStateStatecompliant),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateconflict),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateerror),
		string(ManagedDeviceMobileAppConfigurationSettingStateStatenonCompliant),
		string(ManagedDeviceMobileAppConfigurationSettingStateStatenotApplicable),
		string(ManagedDeviceMobileAppConfigurationSettingStateStatenotAssigned),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateremediated),
		string(ManagedDeviceMobileAppConfigurationSettingStateStateunknown),
	}
}

func parseManagedDeviceMobileAppConfigurationSettingStateState(input string) (*ManagedDeviceMobileAppConfigurationSettingStateState, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationSettingStateState{
		"compliant":     ManagedDeviceMobileAppConfigurationSettingStateStatecompliant,
		"conflict":      ManagedDeviceMobileAppConfigurationSettingStateStateconflict,
		"error":         ManagedDeviceMobileAppConfigurationSettingStateStateerror,
		"noncompliant":  ManagedDeviceMobileAppConfigurationSettingStateStatenonCompliant,
		"notapplicable": ManagedDeviceMobileAppConfigurationSettingStateStatenotApplicable,
		"notassigned":   ManagedDeviceMobileAppConfigurationSettingStateStatenotAssigned,
		"remediated":    ManagedDeviceMobileAppConfigurationSettingStateStateremediated,
		"unknown":       ManagedDeviceMobileAppConfigurationSettingStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationSettingStateState(input)
	return &out, nil
}
