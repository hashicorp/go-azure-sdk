package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppliedConditionalAccessPolicyResult string

const (
	AppliedConditionalAccessPolicyResultfailure    AppliedConditionalAccessPolicyResult = "Failure"
	AppliedConditionalAccessPolicyResultnotApplied AppliedConditionalAccessPolicyResult = "NotApplied"
	AppliedConditionalAccessPolicyResultnotEnabled AppliedConditionalAccessPolicyResult = "NotEnabled"
	AppliedConditionalAccessPolicyResultsuccess    AppliedConditionalAccessPolicyResult = "Success"
	AppliedConditionalAccessPolicyResultunknown    AppliedConditionalAccessPolicyResult = "Unknown"
)

func PossibleValuesForAppliedConditionalAccessPolicyResult() []string {
	return []string{
		string(AppliedConditionalAccessPolicyResultfailure),
		string(AppliedConditionalAccessPolicyResultnotApplied),
		string(AppliedConditionalAccessPolicyResultnotEnabled),
		string(AppliedConditionalAccessPolicyResultsuccess),
		string(AppliedConditionalAccessPolicyResultunknown),
	}
}

func parseAppliedConditionalAccessPolicyResult(input string) (*AppliedConditionalAccessPolicyResult, error) {
	vals := map[string]AppliedConditionalAccessPolicyResult{
		"failure":    AppliedConditionalAccessPolicyResultfailure,
		"notapplied": AppliedConditionalAccessPolicyResultnotApplied,
		"notenabled": AppliedConditionalAccessPolicyResultnotEnabled,
		"success":    AppliedConditionalAccessPolicyResultsuccess,
		"unknown":    AppliedConditionalAccessPolicyResultunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppliedConditionalAccessPolicyResult(input)
	return &out, nil
}
