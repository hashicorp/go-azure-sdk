package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemInfoAuthorizationSystemType string

const (
	AuthorizationSystemInfoAuthorizationSystemType_Aws   AuthorizationSystemInfoAuthorizationSystemType = "aws"
	AuthorizationSystemInfoAuthorizationSystemType_Azure AuthorizationSystemInfoAuthorizationSystemType = "azure"
	AuthorizationSystemInfoAuthorizationSystemType_Gcp   AuthorizationSystemInfoAuthorizationSystemType = "gcp"
)

func PossibleValuesForAuthorizationSystemInfoAuthorizationSystemType() []string {
	return []string{
		string(AuthorizationSystemInfoAuthorizationSystemType_Aws),
		string(AuthorizationSystemInfoAuthorizationSystemType_Azure),
		string(AuthorizationSystemInfoAuthorizationSystemType_Gcp),
	}
}

func (s *AuthorizationSystemInfoAuthorizationSystemType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationSystemInfoAuthorizationSystemType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationSystemInfoAuthorizationSystemType(input string) (*AuthorizationSystemInfoAuthorizationSystemType, error) {
	vals := map[string]AuthorizationSystemInfoAuthorizationSystemType{
		"aws":   AuthorizationSystemInfoAuthorizationSystemType_Aws,
		"azure": AuthorizationSystemInfoAuthorizationSystemType_Azure,
		"gcp":   AuthorizationSystemInfoAuthorizationSystemType_Gcp,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationSystemInfoAuthorizationSystemType(input)
	return &out, nil
}
