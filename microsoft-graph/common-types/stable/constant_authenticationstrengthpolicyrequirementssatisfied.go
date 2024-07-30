package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyRequirementsSatisfied string

const (
	AuthenticationStrengthPolicyRequirementsSatisfied_Mfa  AuthenticationStrengthPolicyRequirementsSatisfied = "mfa"
	AuthenticationStrengthPolicyRequirementsSatisfied_None AuthenticationStrengthPolicyRequirementsSatisfied = "none"
)

func PossibleValuesForAuthenticationStrengthPolicyRequirementsSatisfied() []string {
	return []string{
		string(AuthenticationStrengthPolicyRequirementsSatisfied_Mfa),
		string(AuthenticationStrengthPolicyRequirementsSatisfied_None),
	}
}

func (s *AuthenticationStrengthPolicyRequirementsSatisfied) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthPolicyRequirementsSatisfied(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthPolicyRequirementsSatisfied(input string) (*AuthenticationStrengthPolicyRequirementsSatisfied, error) {
	vals := map[string]AuthenticationStrengthPolicyRequirementsSatisfied{
		"mfa":  AuthenticationStrengthPolicyRequirementsSatisfied_Mfa,
		"none": AuthenticationStrengthPolicyRequirementsSatisfied_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyRequirementsSatisfied(input)
	return &out, nil
}
