package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppPolicyDetailsAuthenticationEvaluation string

const (
	AuthenticationAppPolicyDetailsAuthenticationEvaluation_Failure AuthenticationAppPolicyDetailsAuthenticationEvaluation = "failure"
	AuthenticationAppPolicyDetailsAuthenticationEvaluation_Success AuthenticationAppPolicyDetailsAuthenticationEvaluation = "success"
)

func PossibleValuesForAuthenticationAppPolicyDetailsAuthenticationEvaluation() []string {
	return []string{
		string(AuthenticationAppPolicyDetailsAuthenticationEvaluation_Failure),
		string(AuthenticationAppPolicyDetailsAuthenticationEvaluation_Success),
	}
}

func (s *AuthenticationAppPolicyDetailsAuthenticationEvaluation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAppPolicyDetailsAuthenticationEvaluation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAppPolicyDetailsAuthenticationEvaluation(input string) (*AuthenticationAppPolicyDetailsAuthenticationEvaluation, error) {
	vals := map[string]AuthenticationAppPolicyDetailsAuthenticationEvaluation{
		"failure": AuthenticationAppPolicyDetailsAuthenticationEvaluation_Failure,
		"success": AuthenticationAppPolicyDetailsAuthenticationEvaluation_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppPolicyDetailsAuthenticationEvaluation(input)
	return &out, nil
}
