package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsAuthenticationEvaluation string

const (
	AuthenticationAppPolicyDetailsAuthenticationEvaluationfailure AuthenticationAppPolicyDetailsAuthenticationEvaluation = "Failure"
	AuthenticationAppPolicyDetailsAuthenticationEvaluationsuccess AuthenticationAppPolicyDetailsAuthenticationEvaluation = "Success"
)

func PossibleValuesForAuthenticationAppPolicyDetailsAuthenticationEvaluation() []string {
	return []string{
		string(AuthenticationAppPolicyDetailsAuthenticationEvaluationfailure),
		string(AuthenticationAppPolicyDetailsAuthenticationEvaluationsuccess),
	}
}

func parseAuthenticationAppPolicyDetailsAuthenticationEvaluation(input string) (*AuthenticationAppPolicyDetailsAuthenticationEvaluation, error) {
	vals := map[string]AuthenticationAppPolicyDetailsAuthenticationEvaluation{
		"failure": AuthenticationAppPolicyDetailsAuthenticationEvaluationfailure,
		"success": AuthenticationAppPolicyDetailsAuthenticationEvaluationsuccess,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyDetailsAuthenticationEvaluation(input)
	return &out, nil
}
