package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodConfigurationState string

const (
	VoiceAuthenticationMethodConfigurationStatedisabled VoiceAuthenticationMethodConfigurationState = "Disabled"
	VoiceAuthenticationMethodConfigurationStateenabled  VoiceAuthenticationMethodConfigurationState = "Enabled"
)

func PossibleValuesForVoiceAuthenticationMethodConfigurationState() []string {
	return []string{
		string(VoiceAuthenticationMethodConfigurationStatedisabled),
		string(VoiceAuthenticationMethodConfigurationStateenabled),
	}
}

func parseVoiceAuthenticationMethodConfigurationState(input string) (*VoiceAuthenticationMethodConfigurationState, error) {
	vals := map[string]VoiceAuthenticationMethodConfigurationState{
		"disabled": VoiceAuthenticationMethodConfigurationStatedisabled,
		"enabled":  VoiceAuthenticationMethodConfigurationStateenabled,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VoiceAuthenticationMethodConfigurationState(input)
	return &out, nil
}
