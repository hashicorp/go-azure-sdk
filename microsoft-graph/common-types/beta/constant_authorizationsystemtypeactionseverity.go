package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemTypeActionSeverity string

const (
	AuthorizationSystemTypeActionSeverity_High   AuthorizationSystemTypeActionSeverity = "high"
	AuthorizationSystemTypeActionSeverity_Normal AuthorizationSystemTypeActionSeverity = "normal"
)

func PossibleValuesForAuthorizationSystemTypeActionSeverity() []string {
	return []string{
		string(AuthorizationSystemTypeActionSeverity_High),
		string(AuthorizationSystemTypeActionSeverity_Normal),
	}
}

func (s *AuthorizationSystemTypeActionSeverity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationSystemTypeActionSeverity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationSystemTypeActionSeverity(input string) (*AuthorizationSystemTypeActionSeverity, error) {
	vals := map[string]AuthorizationSystemTypeActionSeverity{
		"high":   AuthorizationSystemTypeActionSeverity_High,
		"normal": AuthorizationSystemTypeActionSeverity_Normal,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationSystemTypeActionSeverity(input)
	return &out, nil
}
