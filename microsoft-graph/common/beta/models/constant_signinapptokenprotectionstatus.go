package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInAppTokenProtectionStatus string

const (
	SignInAppTokenProtectionStatusbound   SignInAppTokenProtectionStatus = "Bound"
	SignInAppTokenProtectionStatusnone    SignInAppTokenProtectionStatus = "None"
	SignInAppTokenProtectionStatusunbound SignInAppTokenProtectionStatus = "Unbound"
)

func PossibleValuesForSignInAppTokenProtectionStatus() []string {
	return []string{
		string(SignInAppTokenProtectionStatusbound),
		string(SignInAppTokenProtectionStatusnone),
		string(SignInAppTokenProtectionStatusunbound),
	}
}

func parseSignInAppTokenProtectionStatus(input string) (*SignInAppTokenProtectionStatus, error) {
	vals := map[string]SignInAppTokenProtectionStatus{
		"bound":   SignInAppTokenProtectionStatusbound,
		"none":    SignInAppTokenProtectionStatusnone,
		"unbound": SignInAppTokenProtectionStatusunbound,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInAppTokenProtectionStatus(input)
	return &out, nil
}
