package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInRiskLevelDuringSignIn string

const (
	SignInRiskLevelDuringSignIn_Hidden SignInRiskLevelDuringSignIn = "hidden"
	SignInRiskLevelDuringSignIn_High   SignInRiskLevelDuringSignIn = "high"
	SignInRiskLevelDuringSignIn_Low    SignInRiskLevelDuringSignIn = "low"
	SignInRiskLevelDuringSignIn_Medium SignInRiskLevelDuringSignIn = "medium"
	SignInRiskLevelDuringSignIn_None   SignInRiskLevelDuringSignIn = "none"
)

func PossibleValuesForSignInRiskLevelDuringSignIn() []string {
	return []string{
		string(SignInRiskLevelDuringSignIn_Hidden),
		string(SignInRiskLevelDuringSignIn_High),
		string(SignInRiskLevelDuringSignIn_Low),
		string(SignInRiskLevelDuringSignIn_Medium),
		string(SignInRiskLevelDuringSignIn_None),
	}
}

func (s *SignInRiskLevelDuringSignIn) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSignInRiskLevelDuringSignIn(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSignInRiskLevelDuringSignIn(input string) (*SignInRiskLevelDuringSignIn, error) {
	vals := map[string]SignInRiskLevelDuringSignIn{
		"hidden": SignInRiskLevelDuringSignIn_Hidden,
		"high":   SignInRiskLevelDuringSignIn_High,
		"low":    SignInRiskLevelDuringSignIn_Low,
		"medium": SignInRiskLevelDuringSignIn_Medium,
		"none":   SignInRiskLevelDuringSignIn_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInRiskLevelDuringSignIn(input)
	return &out, nil
}
