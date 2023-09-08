package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInAuthenticationProtocol string

const (
	SignInAuthenticationProtocolauthenticationTransfer SignInAuthenticationProtocol = "AuthenticationTransfer"
	SignInAuthenticationProtocoldeviceCode             SignInAuthenticationProtocol = "DeviceCode"
	SignInAuthenticationProtocolnone                   SignInAuthenticationProtocol = "None"
	SignInAuthenticationProtocoloAuth2                 SignInAuthenticationProtocol = "OAuth2"
	SignInAuthenticationProtocolropc                   SignInAuthenticationProtocol = "Ropc"
	SignInAuthenticationProtocolsaml20                 SignInAuthenticationProtocol = "Saml20"
	SignInAuthenticationProtocolwsFederation           SignInAuthenticationProtocol = "WsFederation"
)

func PossibleValuesForSignInAuthenticationProtocol() []string {
	return []string{
		string(SignInAuthenticationProtocolauthenticationTransfer),
		string(SignInAuthenticationProtocoldeviceCode),
		string(SignInAuthenticationProtocolnone),
		string(SignInAuthenticationProtocoloAuth2),
		string(SignInAuthenticationProtocolropc),
		string(SignInAuthenticationProtocolsaml20),
		string(SignInAuthenticationProtocolwsFederation),
	}
}

func parseSignInAuthenticationProtocol(input string) (*SignInAuthenticationProtocol, error) {
	vals := map[string]SignInAuthenticationProtocol{
		"authenticationtransfer": SignInAuthenticationProtocolauthenticationTransfer,
		"devicecode":             SignInAuthenticationProtocoldeviceCode,
		"none":                   SignInAuthenticationProtocolnone,
		"oauth2":                 SignInAuthenticationProtocoloAuth2,
		"ropc":                   SignInAuthenticationProtocolropc,
		"saml20":                 SignInAuthenticationProtocolsaml20,
		"wsfederation":           SignInAuthenticationProtocolwsFederation,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInAuthenticationProtocol(input)
	return &out, nil
}
