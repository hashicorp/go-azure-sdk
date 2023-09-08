package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VoiceAuthenticationMethodTargetTargetType string

const (
	VoiceAuthenticationMethodTargetTargetTypegroup VoiceAuthenticationMethodTargetTargetType = "Group"
	VoiceAuthenticationMethodTargetTargetTypeuser  VoiceAuthenticationMethodTargetTargetType = "User"
)

func PossibleValuesForVoiceAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(VoiceAuthenticationMethodTargetTargetTypegroup),
		string(VoiceAuthenticationMethodTargetTargetTypeuser),
	}
}

func parseVoiceAuthenticationMethodTargetTargetType(input string) (*VoiceAuthenticationMethodTargetTargetType, error) {
	vals := map[string]VoiceAuthenticationMethodTargetTargetType{
		"group": VoiceAuthenticationMethodTargetTargetTypegroup,
		"user":  VoiceAuthenticationMethodTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VoiceAuthenticationMethodTargetTargetType(input)
	return &out, nil
}
