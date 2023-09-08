package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodConfigurationState string

const (
	MicrosoftAuthenticatorAuthenticationMethodConfigurationStatedisabled MicrosoftAuthenticatorAuthenticationMethodConfigurationState = "Disabled"
	MicrosoftAuthenticatorAuthenticationMethodConfigurationStateenabled  MicrosoftAuthenticatorAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodConfigurationState() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodConfigurationStatedisabled),
		string(MicrosoftAuthenticatorAuthenticationMethodConfigurationStateenabled),
	}
}

func parseMicrosoftAuthenticatorAuthenticationMethodConfigurationState(input string) (*MicrosoftAuthenticatorAuthenticationMethodConfigurationState, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodConfigurationState{
		"disabled": MicrosoftAuthenticatorAuthenticationMethodConfigurationStatedisabled,
		"enabled":  MicrosoftAuthenticatorAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodConfigurationState(input)
	return &out, nil
}
