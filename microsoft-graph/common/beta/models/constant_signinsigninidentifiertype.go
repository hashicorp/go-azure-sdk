package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignInSignInIdentifierType string

const (
	SignInSignInIdentifierTypeonPremisesUserPrincipalName SignInSignInIdentifierType = "OnPremisesUserPrincipalName"
	SignInSignInIdentifierTypephoneNumber                 SignInSignInIdentifierType = "PhoneNumber"
	SignInSignInIdentifierTypeproxyAddress                SignInSignInIdentifierType = "ProxyAddress"
	SignInSignInIdentifierTypeqrCode                      SignInSignInIdentifierType = "QrCode"
	SignInSignInIdentifierTypeuserPrincipalName           SignInSignInIdentifierType = "UserPrincipalName"
)

func PossibleValuesForSignInSignInIdentifierType() []string {
	return []string{
		string(SignInSignInIdentifierTypeonPremisesUserPrincipalName),
		string(SignInSignInIdentifierTypephoneNumber),
		string(SignInSignInIdentifierTypeproxyAddress),
		string(SignInSignInIdentifierTypeqrCode),
		string(SignInSignInIdentifierTypeuserPrincipalName),
	}
}

func parseSignInSignInIdentifierType(input string) (*SignInSignInIdentifierType, error) {
	vals := map[string]SignInSignInIdentifierType{
		"onpremisesuserprincipalname": SignInSignInIdentifierTypeonPremisesUserPrincipalName,
		"phonenumber":                 SignInSignInIdentifierTypephoneNumber,
		"proxyaddress":                SignInSignInIdentifierTypeproxyAddress,
		"qrcode":                      SignInSignInIdentifierTypeqrCode,
		"userprincipalname":           SignInSignInIdentifierTypeuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SignInSignInIdentifierType(input)
	return &out, nil
}
