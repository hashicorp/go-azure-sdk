package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingDeviceSettingStateState string

const (
	AdvancedThreatProtectionOnboardingDeviceSettingStateStatecompliant     AdvancedThreatProtectionOnboardingDeviceSettingStateState = "Compliant"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStateconflict      AdvancedThreatProtectionOnboardingDeviceSettingStateState = "Conflict"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStateerror         AdvancedThreatProtectionOnboardingDeviceSettingStateState = "Error"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStatenonCompliant  AdvancedThreatProtectionOnboardingDeviceSettingStateState = "NonCompliant"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStatenotApplicable AdvancedThreatProtectionOnboardingDeviceSettingStateState = "NotApplicable"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStatenotAssigned   AdvancedThreatProtectionOnboardingDeviceSettingStateState = "NotAssigned"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStateremediated    AdvancedThreatProtectionOnboardingDeviceSettingStateState = "Remediated"
	AdvancedThreatProtectionOnboardingDeviceSettingStateStateunknown       AdvancedThreatProtectionOnboardingDeviceSettingStateState = "Unknown"
)

func PossibleValuesForAdvancedThreatProtectionOnboardingDeviceSettingStateState() []string {
	return []string{
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStatecompliant),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStateconflict),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStateerror),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStatenonCompliant),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStatenotApplicable),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStatenotAssigned),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStateremediated),
		string(AdvancedThreatProtectionOnboardingDeviceSettingStateStateunknown),
	}
}

func parseAdvancedThreatProtectionOnboardingDeviceSettingStateState(input string) (*AdvancedThreatProtectionOnboardingDeviceSettingStateState, error) {
	vals := map[string]AdvancedThreatProtectionOnboardingDeviceSettingStateState{
		"compliant":     AdvancedThreatProtectionOnboardingDeviceSettingStateStatecompliant,
		"conflict":      AdvancedThreatProtectionOnboardingDeviceSettingStateStateconflict,
		"error":         AdvancedThreatProtectionOnboardingDeviceSettingStateStateerror,
		"noncompliant":  AdvancedThreatProtectionOnboardingDeviceSettingStateStatenonCompliant,
		"notapplicable": AdvancedThreatProtectionOnboardingDeviceSettingStateStatenotApplicable,
		"notassigned":   AdvancedThreatProtectionOnboardingDeviceSettingStateStatenotAssigned,
		"remediated":    AdvancedThreatProtectionOnboardingDeviceSettingStateStateremediated,
		"unknown":       AdvancedThreatProtectionOnboardingDeviceSettingStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AdvancedThreatProtectionOnboardingDeviceSettingStateState(input)
	return &out, nil
}
