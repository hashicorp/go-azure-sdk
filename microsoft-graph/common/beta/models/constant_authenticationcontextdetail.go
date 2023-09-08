package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationContextDetail string

const (
	AuthenticationContextDetailnotApplicable       AuthenticationContextDetail = "NotApplicable"
	AuthenticationContextDetailpreviouslySatisfied AuthenticationContextDetail = "PreviouslySatisfied"
	AuthenticationContextDetailrequired            AuthenticationContextDetail = "Required"
)

func PossibleValuesForAuthenticationContextDetail() []string {
	return []string{
		string(AuthenticationContextDetailnotApplicable),
		string(AuthenticationContextDetailpreviouslySatisfied),
		string(AuthenticationContextDetailrequired),
	}
}

func parseAuthenticationContextDetail(input string) (*AuthenticationContextDetail, error) {
	vals := map[string]AuthenticationContextDetail{
		"notapplicable":       AuthenticationContextDetailnotApplicable,
		"previouslysatisfied": AuthenticationContextDetailpreviouslySatisfied,
		"required":            AuthenticationContextDetailrequired,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationContextDetail(input)
	return &out, nil
}
