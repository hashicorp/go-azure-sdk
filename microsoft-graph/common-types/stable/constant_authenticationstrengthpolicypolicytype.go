package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationStrengthPolicyPolicyType string

const (
	AuthenticationStrengthPolicyPolicyType_BuiltIn AuthenticationStrengthPolicyPolicyType = "builtIn"
	AuthenticationStrengthPolicyPolicyType_Custom  AuthenticationStrengthPolicyPolicyType = "custom"
)

func PossibleValuesForAuthenticationStrengthPolicyPolicyType() []string {
	return []string{
		string(AuthenticationStrengthPolicyPolicyType_BuiltIn),
		string(AuthenticationStrengthPolicyPolicyType_Custom),
	}
}

func (s *AuthenticationStrengthPolicyPolicyType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationStrengthPolicyPolicyType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationStrengthPolicyPolicyType(input string) (*AuthenticationStrengthPolicyPolicyType, error) {
	vals := map[string]AuthenticationStrengthPolicyPolicyType{
		"builtin": AuthenticationStrengthPolicyPolicyType_BuiltIn,
		"custom":  AuthenticationStrengthPolicyPolicyType_Custom,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationStrengthPolicyPolicyType(input)
	return &out, nil
}
