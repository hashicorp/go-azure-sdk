package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineStateState string

const (
	SecurityBaselineStateStateconflict      SecurityBaselineStateState = "Conflict"
	SecurityBaselineStateStateerror         SecurityBaselineStateState = "Error"
	SecurityBaselineStateStatenotApplicable SecurityBaselineStateState = "NotApplicable"
	SecurityBaselineStateStatenotSecure     SecurityBaselineStateState = "NotSecure"
	SecurityBaselineStateStatesecure        SecurityBaselineStateState = "Secure"
	SecurityBaselineStateStateunknown       SecurityBaselineStateState = "Unknown"
)

func PossibleValuesForSecurityBaselineStateState() []string {
	return []string{
		string(SecurityBaselineStateStateconflict),
		string(SecurityBaselineStateStateerror),
		string(SecurityBaselineStateStatenotApplicable),
		string(SecurityBaselineStateStatenotSecure),
		string(SecurityBaselineStateStatesecure),
		string(SecurityBaselineStateStateunknown),
	}
}

func parseSecurityBaselineStateState(input string) (*SecurityBaselineStateState, error) {
	vals := map[string]SecurityBaselineStateState{
		"conflict":      SecurityBaselineStateStateconflict,
		"error":         SecurityBaselineStateStateerror,
		"notapplicable": SecurityBaselineStateStatenotApplicable,
		"notsecure":     SecurityBaselineStateStatenotSecure,
		"secure":        SecurityBaselineStateStatesecure,
		"unknown":       SecurityBaselineStateStateunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineStateState(input)
	return &out, nil
}
