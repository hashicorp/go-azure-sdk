package sqlpoolssecurityalertpolicies

import "strings"

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertPolicyState string

const (
	SecurityAlertPolicyStateDisabled SecurityAlertPolicyState = "Disabled"
	SecurityAlertPolicyStateEnabled  SecurityAlertPolicyState = "Enabled"
	SecurityAlertPolicyStateNew      SecurityAlertPolicyState = "New"
)

func PossibleValuesForSecurityAlertPolicyState() []string {
	return []string{
		string(SecurityAlertPolicyStateDisabled),
		string(SecurityAlertPolicyStateEnabled),
		string(SecurityAlertPolicyStateNew),
	}
}

func parseSecurityAlertPolicyState(input string) (*SecurityAlertPolicyState, error) {
	vals := map[string]SecurityAlertPolicyState{
		"disabled": SecurityAlertPolicyStateDisabled,
		"enabled":  SecurityAlertPolicyStateEnabled,
		"new":      SecurityAlertPolicyStateNew,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAlertPolicyState(input)
	return &out, nil
}
