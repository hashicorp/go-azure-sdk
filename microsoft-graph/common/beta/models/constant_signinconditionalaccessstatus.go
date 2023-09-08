package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInConditionalAccessStatus string

const (
	SignInConditionalAccessStatusfailure    SignInConditionalAccessStatus = "Failure"
	SignInConditionalAccessStatusnotApplied SignInConditionalAccessStatus = "NotApplied"
	SignInConditionalAccessStatussuccess    SignInConditionalAccessStatus = "Success"
)

func PossibleValuesForSignInConditionalAccessStatus() []string {
	return []string{
		string(SignInConditionalAccessStatusfailure),
		string(SignInConditionalAccessStatusnotApplied),
		string(SignInConditionalAccessStatussuccess),
	}
}

func parseSignInConditionalAccessStatus(input string) (*SignInConditionalAccessStatus, error) {
	vals := map[string]SignInConditionalAccessStatus{
		"failure":    SignInConditionalAccessStatusfailure,
		"notapplied": SignInConditionalAccessStatusnotApplied,
		"success":    SignInConditionalAccessStatussuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInConditionalAccessStatus(input)
	return &out, nil
}
