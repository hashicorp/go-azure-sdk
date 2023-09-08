package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MicrosoftAuthenticatorAuthenticationMethodTargetTargetType string

const (
	MicrosoftAuthenticatorAuthenticationMethodTargetTargetTypegroup MicrosoftAuthenticatorAuthenticationMethodTargetTargetType = "Group"
	MicrosoftAuthenticatorAuthenticationMethodTargetTargetTypeuser  MicrosoftAuthenticatorAuthenticationMethodTargetTargetType = "User"
)

func PossibleValuesForMicrosoftAuthenticatorAuthenticationMethodTargetTargetType() []string {
	return []string{
		string(MicrosoftAuthenticatorAuthenticationMethodTargetTargetTypegroup),
		string(MicrosoftAuthenticatorAuthenticationMethodTargetTargetTypeuser),
	}
}

func parseMicrosoftAuthenticatorAuthenticationMethodTargetTargetType(input string) (*MicrosoftAuthenticatorAuthenticationMethodTargetTargetType, error) {
	vals := map[string]MicrosoftAuthenticatorAuthenticationMethodTargetTargetType{
		"group": MicrosoftAuthenticatorAuthenticationMethodTargetTargetTypegroup,
		"user":  MicrosoftAuthenticatorAuthenticationMethodTargetTargetTypeuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MicrosoftAuthenticatorAuthenticationMethodTargetTargetType(input)
	return &out, nil
}
