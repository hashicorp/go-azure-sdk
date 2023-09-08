package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyRequirementsSatisfied string

const (
	AuthenticationStrengthPolicyRequirementsSatisfiedmfa  AuthenticationStrengthPolicyRequirementsSatisfied = "Mfa"
	AuthenticationStrengthPolicyRequirementsSatisfiednone AuthenticationStrengthPolicyRequirementsSatisfied = "None"
)

func PossibleValuesForAuthenticationStrengthPolicyRequirementsSatisfied() []string {
	return []string{
		string(AuthenticationStrengthPolicyRequirementsSatisfiedmfa),
		string(AuthenticationStrengthPolicyRequirementsSatisfiednone),
	}
}

func parseAuthenticationStrengthPolicyRequirementsSatisfied(input string) (*AuthenticationStrengthPolicyRequirementsSatisfied, error) {
	vals := map[string]AuthenticationStrengthPolicyRequirementsSatisfied{
		"mfa":  AuthenticationStrengthPolicyRequirementsSatisfiedmfa,
		"none": AuthenticationStrengthPolicyRequirementsSatisfiednone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyRequirementsSatisfied(input)
	return &out, nil
}
