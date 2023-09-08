package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HardwareOathAuthenticationMethodConfigurationState string

const (
	HardwareOathAuthenticationMethodConfigurationStatedisabled HardwareOathAuthenticationMethodConfigurationState = "Disabled"
	HardwareOathAuthenticationMethodConfigurationStateenabled  HardwareOathAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForHardwareOathAuthenticationMethodConfigurationState() []string {
	return []string{
		string(HardwareOathAuthenticationMethodConfigurationStatedisabled),
		string(HardwareOathAuthenticationMethodConfigurationStateenabled),
	}
}

func parseHardwareOathAuthenticationMethodConfigurationState(input string) (*HardwareOathAuthenticationMethodConfigurationState, error) {
	vals := map[string]HardwareOathAuthenticationMethodConfigurationState{
		"disabled": HardwareOathAuthenticationMethodConfigurationStatedisabled,
		"enabled":  HardwareOathAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := HardwareOathAuthenticationMethodConfigurationState(input)
	return &out, nil
}
