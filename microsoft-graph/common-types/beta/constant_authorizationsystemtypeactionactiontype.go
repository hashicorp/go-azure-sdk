package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthorizationSystemTypeActionActionType string

const (
	AuthorizationSystemTypeActionActionType_Delete AuthorizationSystemTypeActionActionType = "delete"
	AuthorizationSystemTypeActionActionType_Read   AuthorizationSystemTypeActionActionType = "read"
)

func PossibleValuesForAuthorizationSystemTypeActionActionType() []string {
	return []string{
		string(AuthorizationSystemTypeActionActionType_Delete),
		string(AuthorizationSystemTypeActionActionType_Read),
	}
}

func (s *AuthorizationSystemTypeActionActionType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthorizationSystemTypeActionActionType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthorizationSystemTypeActionActionType(input string) (*AuthorizationSystemTypeActionActionType, error) {
	vals := map[string]AuthorizationSystemTypeActionActionType{
		"delete": AuthorizationSystemTypeActionActionType_Delete,
		"read":   AuthorizationSystemTypeActionActionType_Read,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthorizationSystemTypeActionActionType(input)
	return &out, nil
}
