package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInIncomingTokenType string

const (
	SignInIncomingTokenTypenone                SignInIncomingTokenType = "None"
	SignInIncomingTokenTypeprimaryRefreshToken SignInIncomingTokenType = "PrimaryRefreshToken"
	SignInIncomingTokenTyperemoteDesktopToken  SignInIncomingTokenType = "RemoteDesktopToken"
	SignInIncomingTokenTypesaml11              SignInIncomingTokenType = "Saml11"
	SignInIncomingTokenTypesaml20              SignInIncomingTokenType = "Saml20"
)

func PossibleValuesForSignInIncomingTokenType() []string {
	return []string{
		string(SignInIncomingTokenTypenone),
		string(SignInIncomingTokenTypeprimaryRefreshToken),
		string(SignInIncomingTokenTyperemoteDesktopToken),
		string(SignInIncomingTokenTypesaml11),
		string(SignInIncomingTokenTypesaml20),
	}
}

func parseSignInIncomingTokenType(input string) (*SignInIncomingTokenType, error) {
	vals := map[string]SignInIncomingTokenType{
		"none":                SignInIncomingTokenTypenone,
		"primaryrefreshtoken": SignInIncomingTokenTypeprimaryRefreshToken,
		"remotedesktoptoken":  SignInIncomingTokenTyperemoteDesktopToken,
		"saml11":              SignInIncomingTokenTypesaml11,
		"saml20":              SignInIncomingTokenTypesaml20,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInIncomingTokenType(input)
	return &out, nil
}
