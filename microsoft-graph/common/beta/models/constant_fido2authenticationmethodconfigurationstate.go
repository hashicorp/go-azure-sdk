package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Fido2AuthenticationMethodConfigurationState string

const (
	Fido2AuthenticationMethodConfigurationStatedisabled Fido2AuthenticationMethodConfigurationState = "Disabled"
	Fido2AuthenticationMethodConfigurationStateenabled  Fido2AuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForFido2AuthenticationMethodConfigurationState() []string {
	return []string{
		string(Fido2AuthenticationMethodConfigurationStatedisabled),
		string(Fido2AuthenticationMethodConfigurationStateenabled),
	}
}

func parseFido2AuthenticationMethodConfigurationState(input string) (*Fido2AuthenticationMethodConfigurationState, error) {
	vals := map[string]Fido2AuthenticationMethodConfigurationState{
		"disabled": Fido2AuthenticationMethodConfigurationStatedisabled,
		"enabled":  Fido2AuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Fido2AuthenticationMethodConfigurationState(input)
	return &out, nil
}
