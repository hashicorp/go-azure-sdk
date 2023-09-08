package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodTargetTargetType string

const (
	AuthenticationMethodTargetTargetTypegroup AuthenticationMethodTargetTargetType = "Group"
	AuthenticationMethodTargetTargetTypeuser  AuthenticationMethodTargetTargetType = "User"
)

func PossibleValuesForAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(AuthenticationMethodTargetTargetTypegroup),
		string(AuthenticationMethodTargetTargetTypeuser),
	}
}

func parseAuthenticationMethodTargetTargetType(input string) (*AuthenticationMethodTargetTargetType, error) {
	vals := map[string]AuthenticationMethodTargetTargetType{
		"group": AuthenticationMethodTargetTargetTypegroup,
		"user":  AuthenticationMethodTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodTargetTargetType(input)
	return &out, nil
}
