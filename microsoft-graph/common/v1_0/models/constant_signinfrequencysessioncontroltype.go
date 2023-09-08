package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInFrequencySessionControlType string

const (
	SignInFrequencySessionControlTypedays  SignInFrequencySessionControlType = "Days"
	SignInFrequencySessionControlTypehours SignInFrequencySessionControlType = "Hours"
)

func PossibleValuesForSignInFrequencySessionControlType() []string {
	return []string{
		string(SignInFrequencySessionControlTypedays),
		string(SignInFrequencySessionControlTypehours),
	}
}

func parseSignInFrequencySessionControlType(input string) (*SignInFrequencySessionControlType, error) {
	vals := map[string]SignInFrequencySessionControlType{
		"days":  SignInFrequencySessionControlTypedays,
		"hours": SignInFrequencySessionControlTypehours,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInFrequencySessionControlType(input)
	return &out, nil
}
