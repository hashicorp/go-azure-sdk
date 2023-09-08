package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TemporaryAccessPassAuthenticationMethodConfigurationState string

const (
	TemporaryAccessPassAuthenticationMethodConfigurationStatedisabled TemporaryAccessPassAuthenticationMethodConfigurationState = "Disabled"
	TemporaryAccessPassAuthenticationMethodConfigurationStateenabled  TemporaryAccessPassAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForTemporaryAccessPassAuthenticationMethodConfigurationState() []string {
	return []string{
		string(TemporaryAccessPassAuthenticationMethodConfigurationStatedisabled),
		string(TemporaryAccessPassAuthenticationMethodConfigurationStateenabled),
	}
}

func parseTemporaryAccessPassAuthenticationMethodConfigurationState(input string) (*TemporaryAccessPassAuthenticationMethodConfigurationState, error) {
	vals := map[string]TemporaryAccessPassAuthenticationMethodConfigurationState{
		"disabled": TemporaryAccessPassAuthenticationMethodConfigurationStatedisabled,
		"enabled":  TemporaryAccessPassAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TemporaryAccessPassAuthenticationMethodConfigurationState(input)
	return &out, nil
}
