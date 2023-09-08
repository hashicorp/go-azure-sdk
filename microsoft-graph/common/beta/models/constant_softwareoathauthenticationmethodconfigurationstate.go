package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SoftwareOathAuthenticationMethodConfigurationState string

const (
	SoftwareOathAuthenticationMethodConfigurationStatedisabled SoftwareOathAuthenticationMethodConfigurationState = "Disabled"
	SoftwareOathAuthenticationMethodConfigurationStateenabled  SoftwareOathAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForSoftwareOathAuthenticationMethodConfigurationState() []string {
	return []string{
		string(SoftwareOathAuthenticationMethodConfigurationStatedisabled),
		string(SoftwareOathAuthenticationMethodConfigurationStateenabled),
	}
}

func parseSoftwareOathAuthenticationMethodConfigurationState(input string) (*SoftwareOathAuthenticationMethodConfigurationState, error) {
	vals := map[string]SoftwareOathAuthenticationMethodConfigurationState{
		"disabled": SoftwareOathAuthenticationMethodConfigurationStatedisabled,
		"enabled":  SoftwareOathAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SoftwareOathAuthenticationMethodConfigurationState(input)
	return &out, nil
}
