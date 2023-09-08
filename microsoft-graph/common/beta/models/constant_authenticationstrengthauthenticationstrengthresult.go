package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthAuthenticationStrengthResult string

const (
	AuthenticationStrengthAuthenticationStrengthResultcannotSatisfy                              AuthenticationStrengthAuthenticationStrengthResult = "CannotSatisfy"
	AuthenticationStrengthAuthenticationStrengthResultcannotSatisfyDueToCombinationConfiguration AuthenticationStrengthAuthenticationStrengthResult = "CannotSatisfyDueToCombinationConfiguration"
	AuthenticationStrengthAuthenticationStrengthResultmultipleChallengesRequired                 AuthenticationStrengthAuthenticationStrengthResult = "MultipleChallengesRequired"
	AuthenticationStrengthAuthenticationStrengthResultmultipleRegistrationsRequired              AuthenticationStrengthAuthenticationStrengthResult = "MultipleRegistrationsRequired"
	AuthenticationStrengthAuthenticationStrengthResultnotSet                                     AuthenticationStrengthAuthenticationStrengthResult = "NotSet"
	AuthenticationStrengthAuthenticationStrengthResultsatisfied                                  AuthenticationStrengthAuthenticationStrengthResult = "Satisfied"
	AuthenticationStrengthAuthenticationStrengthResultsingleChallengeRequired                    AuthenticationStrengthAuthenticationStrengthResult = "SingleChallengeRequired"
	AuthenticationStrengthAuthenticationStrengthResultsingleRegistrationRequired                 AuthenticationStrengthAuthenticationStrengthResult = "SingleRegistrationRequired"
	AuthenticationStrengthAuthenticationStrengthResultskippedForProofUp                          AuthenticationStrengthAuthenticationStrengthResult = "SkippedForProofUp"
)

func PossibleValuesForAuthenticationStrengthAuthenticationStrengthResult() []string {
	return []string{
		string(AuthenticationStrengthAuthenticationStrengthResultcannotSatisfy),
		string(AuthenticationStrengthAuthenticationStrengthResultcannotSatisfyDueToCombinationConfiguration),
		string(AuthenticationStrengthAuthenticationStrengthResultmultipleChallengesRequired),
		string(AuthenticationStrengthAuthenticationStrengthResultmultipleRegistrationsRequired),
		string(AuthenticationStrengthAuthenticationStrengthResultnotSet),
		string(AuthenticationStrengthAuthenticationStrengthResultsatisfied),
		string(AuthenticationStrengthAuthenticationStrengthResultsingleChallengeRequired),
		string(AuthenticationStrengthAuthenticationStrengthResultsingleRegistrationRequired),
		string(AuthenticationStrengthAuthenticationStrengthResultskippedForProofUp),
	}
}

func parseAuthenticationStrengthAuthenticationStrengthResult(input string) (*AuthenticationStrengthAuthenticationStrengthResult, error) {
	vals := map[string]AuthenticationStrengthAuthenticationStrengthResult{
		"cannotsatisfy": AuthenticationStrengthAuthenticationStrengthResultcannotSatisfy,
		"cannotsatisfyduetocombinationconfiguration": AuthenticationStrengthAuthenticationStrengthResultcannotSatisfyDueToCombinationConfiguration,
		"multiplechallengesrequired":                 AuthenticationStrengthAuthenticationStrengthResultmultipleChallengesRequired,
		"multipleregistrationsrequired":              AuthenticationStrengthAuthenticationStrengthResultmultipleRegistrationsRequired,
		"notset":                                     AuthenticationStrengthAuthenticationStrengthResultnotSet,
		"satisfied":                                  AuthenticationStrengthAuthenticationStrengthResultsatisfied,
		"singlechallengerequired":                    AuthenticationStrengthAuthenticationStrengthResultsingleChallengeRequired,
		"singleregistrationrequired":                 AuthenticationStrengthAuthenticationStrengthResultsingleRegistrationRequired,
		"skippedforproofup":                          AuthenticationStrengthAuthenticationStrengthResultskippedForProofUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthAuthenticationStrengthResult(input)
	return &out, nil
}
