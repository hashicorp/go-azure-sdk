package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleVpnConfigurationAuthenticationMethod string

const (
	AppleVpnConfigurationAuthenticationMethodazureAD             AppleVpnConfigurationAuthenticationMethod = "AzureAD"
	AppleVpnConfigurationAuthenticationMethodcertificate         AppleVpnConfigurationAuthenticationMethod = "Certificate"
	AppleVpnConfigurationAuthenticationMethodderivedCredential   AppleVpnConfigurationAuthenticationMethod = "DerivedCredential"
	AppleVpnConfigurationAuthenticationMethodsharedSecret        AppleVpnConfigurationAuthenticationMethod = "SharedSecret"
	AppleVpnConfigurationAuthenticationMethodusernameAndPassword AppleVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAppleVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AppleVpnConfigurationAuthenticationMethodazureAD),
		string(AppleVpnConfigurationAuthenticationMethodcertificate),
		string(AppleVpnConfigurationAuthenticationMethodderivedCredential),
		string(AppleVpnConfigurationAuthenticationMethodsharedSecret),
		string(AppleVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAppleVpnConfigurationAuthenticationMethod(input string) (*AppleVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AppleVpnConfigurationAuthenticationMethod{
		"azuread":             AppleVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         AppleVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AppleVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        AppleVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": AppleVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
