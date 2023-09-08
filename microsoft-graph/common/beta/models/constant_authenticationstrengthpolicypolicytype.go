package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyPolicyType string

const (
	AuthenticationStrengthPolicyPolicyTypebuiltIn AuthenticationStrengthPolicyPolicyType = "BuiltIn"
	AuthenticationStrengthPolicyPolicyTypecustom  AuthenticationStrengthPolicyPolicyType = "Custom"
)

func PossibleValuesForAuthenticationStrengthPolicyPolicyType() []string {
	return []string{
		string(AuthenticationStrengthPolicyPolicyTypebuiltIn),
		string(AuthenticationStrengthPolicyPolicyTypecustom),
	}
}

func parseAuthenticationStrengthPolicyPolicyType(input string) (*AuthenticationStrengthPolicyPolicyType, error) {
	vals := map[string]AuthenticationStrengthPolicyPolicyType{
		"builtin": AuthenticationStrengthPolicyPolicyTypebuiltIn,
		"custom":  AuthenticationStrengthPolicyPolicyTypecustom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyPolicyType(input)
	return &out, nil
}
