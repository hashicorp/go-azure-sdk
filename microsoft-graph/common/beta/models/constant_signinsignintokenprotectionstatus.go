package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInSignInTokenProtectionStatus string

const (
	SignInSignInTokenProtectionStatusbound   SignInSignInTokenProtectionStatus = "Bound"
	SignInSignInTokenProtectionStatusnone    SignInSignInTokenProtectionStatus = "None"
	SignInSignInTokenProtectionStatusunbound SignInSignInTokenProtectionStatus = "Unbound"
)

func PossibleValuesForSignInSignInTokenProtectionStatus() []string {
	return []string{
		string(SignInSignInTokenProtectionStatusbound),
		string(SignInSignInTokenProtectionStatusnone),
		string(SignInSignInTokenProtectionStatusunbound),
	}
}

func parseSignInSignInTokenProtectionStatus(input string) (*SignInSignInTokenProtectionStatus, error) {
	vals := map[string]SignInSignInTokenProtectionStatus{
		"bound":   SignInSignInTokenProtectionStatusbound,
		"none":    SignInSignInTokenProtectionStatusnone,
		"unbound": SignInSignInTokenProtectionStatusunbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInSignInTokenProtectionStatus(input)
	return &out, nil
}
