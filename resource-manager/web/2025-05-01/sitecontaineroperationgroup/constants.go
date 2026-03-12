package sitecontaineroperationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthType string

const (
	AuthTypeAnonymous       AuthType = "Anonymous"
	AuthTypeSystemIdentity  AuthType = "SystemIdentity"
	AuthTypeUserAssigned    AuthType = "UserAssigned"
	AuthTypeUserCredentials AuthType = "UserCredentials"
)

func PossibleValuesForAuthType() []string {
	return []string{
		string(AuthTypeAnonymous),
		string(AuthTypeSystemIdentity),
		string(AuthTypeUserAssigned),
		string(AuthTypeUserCredentials),
	}
}

func (s *AuthType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAuthType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAuthType(input string) (*AuthType, error) {
	vals := map[string]AuthType{
		"anonymous":       AuthTypeAnonymous,
		"systemidentity":  AuthTypeSystemIdentity,
		"userassigned":    AuthTypeUserAssigned,
		"usercredentials": AuthTypeUserCredentials,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthType(input)
	return &out, nil
}
