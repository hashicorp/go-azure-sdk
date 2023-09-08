package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSVpnConfigurationAuthenticationMethod string

const (
	MacOSVpnConfigurationAuthenticationMethodazureAD             MacOSVpnConfigurationAuthenticationMethod = "AzureAD"
	MacOSVpnConfigurationAuthenticationMethodcertificate         MacOSVpnConfigurationAuthenticationMethod = "Certificate"
	MacOSVpnConfigurationAuthenticationMethodderivedCredential   MacOSVpnConfigurationAuthenticationMethod = "DerivedCredential"
	MacOSVpnConfigurationAuthenticationMethodsharedSecret        MacOSVpnConfigurationAuthenticationMethod = "SharedSecret"
	MacOSVpnConfigurationAuthenticationMethodusernameAndPassword MacOSVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForMacOSVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(MacOSVpnConfigurationAuthenticationMethodazureAD),
		string(MacOSVpnConfigurationAuthenticationMethodcertificate),
		string(MacOSVpnConfigurationAuthenticationMethodderivedCredential),
		string(MacOSVpnConfigurationAuthenticationMethodsharedSecret),
		string(MacOSVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseMacOSVpnConfigurationAuthenticationMethod(input string) (*MacOSVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]MacOSVpnConfigurationAuthenticationMethod{
		"azuread":             MacOSVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         MacOSVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   MacOSVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        MacOSVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": MacOSVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
