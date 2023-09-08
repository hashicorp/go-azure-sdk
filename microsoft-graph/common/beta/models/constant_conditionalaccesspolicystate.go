package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessPolicyState string

const (
	ConditionalAccessPolicyStatedisabled                          ConditionalAccessPolicyState = "Disabled"
	ConditionalAccessPolicyStateenabled                           ConditionalAccessPolicyState = "Enabled"
	ConditionalAccessPolicyStateenabledForReportingButNotEnforced ConditionalAccessPolicyState = "EnabledForReportingButNotEnforced"
)

func PossibleValuesForConditionalAccessPolicyState() []string {
	return []string{
		string(ConditionalAccessPolicyStatedisabled),
		string(ConditionalAccessPolicyStateenabled),
		string(ConditionalAccessPolicyStateenabledForReportingButNotEnforced),
	}
}

func parseConditionalAccessPolicyState(input string) (*ConditionalAccessPolicyState, error) {
	vals := map[string]ConditionalAccessPolicyState{
		"disabled":                          ConditionalAccessPolicyStatedisabled,
		"enabled":                           ConditionalAccessPolicyStateenabled,
		"enabledforreportingbutnotenforced": ConditionalAccessPolicyStateenabledForReportingButNotEnforced,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ConditionalAccessPolicyState(input)
	return &out, nil
}
