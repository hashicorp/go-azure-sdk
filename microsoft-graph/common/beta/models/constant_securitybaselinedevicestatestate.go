package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineDeviceStateState string

const (
	SecurityBaselineDeviceStateStateconflict      SecurityBaselineDeviceStateState = "Conflict"
	SecurityBaselineDeviceStateStateerror         SecurityBaselineDeviceStateState = "Error"
	SecurityBaselineDeviceStateStatenotApplicable SecurityBaselineDeviceStateState = "NotApplicable"
	SecurityBaselineDeviceStateStatenotSecure     SecurityBaselineDeviceStateState = "NotSecure"
	SecurityBaselineDeviceStateStatesecure        SecurityBaselineDeviceStateState = "Secure"
	SecurityBaselineDeviceStateStateunknown       SecurityBaselineDeviceStateState = "Unknown"
)

func PossibleValuesForSecurityBaselineDeviceStateState() []string {
	return []string{
		string(SecurityBaselineDeviceStateStateconflict),
		string(SecurityBaselineDeviceStateStateerror),
		string(SecurityBaselineDeviceStateStatenotApplicable),
		string(SecurityBaselineDeviceStateStatenotSecure),
		string(SecurityBaselineDeviceStateStatesecure),
		string(SecurityBaselineDeviceStateStateunknown),
	}
}

func parseSecurityBaselineDeviceStateState(input string) (*SecurityBaselineDeviceStateState, error) {
	vals := map[string]SecurityBaselineDeviceStateState{
		"conflict":      SecurityBaselineDeviceStateStateconflict,
		"error":         SecurityBaselineDeviceStateStateerror,
		"notapplicable": SecurityBaselineDeviceStateStatenotApplicable,
		"notsecure":     SecurityBaselineDeviceStateStatenotSecure,
		"secure":        SecurityBaselineDeviceStateStatesecure,
		"unknown":       SecurityBaselineDeviceStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineDeviceStateState(input)
	return &out, nil
}
