package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineSettingStateState string

const (
	SecurityBaselineSettingStateStateconflict      SecurityBaselineSettingStateState = "Conflict"
	SecurityBaselineSettingStateStateerror         SecurityBaselineSettingStateState = "Error"
	SecurityBaselineSettingStateStatenotApplicable SecurityBaselineSettingStateState = "NotApplicable"
	SecurityBaselineSettingStateStatenotSecure     SecurityBaselineSettingStateState = "NotSecure"
	SecurityBaselineSettingStateStatesecure        SecurityBaselineSettingStateState = "Secure"
	SecurityBaselineSettingStateStateunknown       SecurityBaselineSettingStateState = "Unknown"
)

func PossibleValuesForSecurityBaselineSettingStateState() []string {
	return []string{
		string(SecurityBaselineSettingStateStateconflict),
		string(SecurityBaselineSettingStateStateerror),
		string(SecurityBaselineSettingStateStatenotApplicable),
		string(SecurityBaselineSettingStateStatenotSecure),
		string(SecurityBaselineSettingStateStatesecure),
		string(SecurityBaselineSettingStateStateunknown),
	}
}

func parseSecurityBaselineSettingStateState(input string) (*SecurityBaselineSettingStateState, error) {
	vals := map[string]SecurityBaselineSettingStateState{
		"conflict":      SecurityBaselineSettingStateStateconflict,
		"error":         SecurityBaselineSettingStateStateerror,
		"notapplicable": SecurityBaselineSettingStateStatenotApplicable,
		"notsecure":     SecurityBaselineSettingStateStatenotSecure,
		"secure":        SecurityBaselineSettingStateStatesecure,
		"unknown":       SecurityBaselineSettingStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineSettingStateState(input)
	return &out, nil
}
