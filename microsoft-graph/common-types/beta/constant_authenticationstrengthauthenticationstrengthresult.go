package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthAuthenticationStrengthResult string

const (
	AuthenticationStrengthAuthenticationStrengthResult_CannotSatisfy                              AuthenticationStrengthAuthenticationStrengthResult = "cannotSatisfy"
	AuthenticationStrengthAuthenticationStrengthResult_CannotSatisfyDueToCombinationConfiguration AuthenticationStrengthAuthenticationStrengthResult = "cannotSatisfyDueToCombinationConfiguration"
	AuthenticationStrengthAuthenticationStrengthResult_MultipleChallengesRequired                 AuthenticationStrengthAuthenticationStrengthResult = "multipleChallengesRequired"
	AuthenticationStrengthAuthenticationStrengthResult_MultipleRegistrationsRequired              AuthenticationStrengthAuthenticationStrengthResult = "multipleRegistrationsRequired"
	AuthenticationStrengthAuthenticationStrengthResult_NotSet                                     AuthenticationStrengthAuthenticationStrengthResult = "notSet"
	AuthenticationStrengthAuthenticationStrengthResult_Satisfied                                  AuthenticationStrengthAuthenticationStrengthResult = "satisfied"
	AuthenticationStrengthAuthenticationStrengthResult_SingleChallengeRequired                    AuthenticationStrengthAuthenticationStrengthResult = "singleChallengeRequired"
	AuthenticationStrengthAuthenticationStrengthResult_SingleRegistrationRequired                 AuthenticationStrengthAuthenticationStrengthResult = "singleRegistrationRequired"
	AuthenticationStrengthAuthenticationStrengthResult_SkippedForProofUp                          AuthenticationStrengthAuthenticationStrengthResult = "skippedForProofUp"
)

func PossibleValuesForAuthenticationStrengthAuthenticationStrengthResult() []string {
	return []string{
		string(AuthenticationStrengthAuthenticationStrengthResult_CannotSatisfy),
		string(AuthenticationStrengthAuthenticationStrengthResult_CannotSatisfyDueToCombinationConfiguration),
		string(AuthenticationStrengthAuthenticationStrengthResult_MultipleChallengesRequired),
		string(AuthenticationStrengthAuthenticationStrengthResult_MultipleRegistrationsRequired),
		string(AuthenticationStrengthAuthenticationStrengthResult_NotSet),
		string(AuthenticationStrengthAuthenticationStrengthResult_Satisfied),
		string(AuthenticationStrengthAuthenticationStrengthResult_SingleChallengeRequired),
		string(AuthenticationStrengthAuthenticationStrengthResult_SingleRegistrationRequired),
		string(AuthenticationStrengthAuthenticationStrengthResult_SkippedForProofUp),
	}
}

func (s *AuthenticationStrengthAuthenticationStrengthResult) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthAuthenticationStrengthResult(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthAuthenticationStrengthResult(input string) (*AuthenticationStrengthAuthenticationStrengthResult, error) {
	vals := map[string]AuthenticationStrengthAuthenticationStrengthResult{
		"cannotsatisfy": AuthenticationStrengthAuthenticationStrengthResult_CannotSatisfy,
		"cannotsatisfyduetocombinationconfiguration": AuthenticationStrengthAuthenticationStrengthResult_CannotSatisfyDueToCombinationConfiguration,
		"multiplechallengesrequired":                 AuthenticationStrengthAuthenticationStrengthResult_MultipleChallengesRequired,
		"multipleregistrationsrequired":              AuthenticationStrengthAuthenticationStrengthResult_MultipleRegistrationsRequired,
		"notset":                                     AuthenticationStrengthAuthenticationStrengthResult_NotSet,
		"satisfied":                                  AuthenticationStrengthAuthenticationStrengthResult_Satisfied,
		"singlechallengerequired":                    AuthenticationStrengthAuthenticationStrengthResult_SingleChallengeRequired,
		"singleregistrationrequired":                 AuthenticationStrengthAuthenticationStrengthResult_SingleRegistrationRequired,
		"skippedforproofup":                          AuthenticationStrengthAuthenticationStrengthResult_SkippedForProofUp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthAuthenticationStrengthResult(input)
	return &out, nil
}
