package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuthenticationMethodConfigurationState string

const (
	AuthenticationMethodConfigurationStatedisabled AuthenticationMethodConfigurationState = "Disabled"
	AuthenticationMethodConfigurationStateenabled  AuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForAuthenticationMethodConfigurationState() []string {
	return []string{
		string(AuthenticationMethodConfigurationStatedisabled),
		string(AuthenticationMethodConfigurationStateenabled),
	}
}

func parseAuthenticationMethodConfigurationState(input string) (*AuthenticationMethodConfigurationState, error) {
	vals := map[string]AuthenticationMethodConfigurationState{
		"disabled": AuthenticationMethodConfigurationStatedisabled,
		"enabled":  AuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AuthenticationMethodConfigurationState(input)
	return &out, nil
}
