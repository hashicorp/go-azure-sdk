package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationAppEvaluation string

const (
	AuthenticationAppEvaluation_Failure AuthenticationAppEvaluation = "failure"
	AuthenticationAppEvaluation_Success AuthenticationAppEvaluation = "success"
)

func PossibleValuesForAuthenticationAppEvaluation() []string {
	return []string{
		string(AuthenticationAppEvaluation_Failure),
		string(AuthenticationAppEvaluation_Success),
	}
}

func (s *AuthenticationAppEvaluation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationAppEvaluation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationAppEvaluation(input string) (*AuthenticationAppEvaluation, error) {
	vals := map[string]AuthenticationAppEvaluation{
		"failure": AuthenticationAppEvaluation_Failure,
		"success": AuthenticationAppEvaluation_Success,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationAppEvaluation(input)
	return &out, nil
}
