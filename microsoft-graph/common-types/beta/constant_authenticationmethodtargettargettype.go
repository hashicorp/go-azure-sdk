package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodTargetTargetType string

const (
	AuthenticationMethodTargetTargetType_Group AuthenticationMethodTargetTargetType = "group"
	AuthenticationMethodTargetTargetType_User  AuthenticationMethodTargetTargetType = "user"
)

func PossibleValuesForAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(AuthenticationMethodTargetTargetType_Group),
		string(AuthenticationMethodTargetTargetType_User),
	}
}

func (s *AuthenticationMethodTargetTargetType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthenticationMethodTargetTargetType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthenticationMethodTargetTargetType(input string) (*AuthenticationMethodTargetTargetType, error) {
	vals := map[string]AuthenticationMethodTargetTargetType{
		"group": AuthenticationMethodTargetTargetType_Group,
		"user":  AuthenticationMethodTargetTargetType_User,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodTargetTargetType(input)
	return &out, nil
}
