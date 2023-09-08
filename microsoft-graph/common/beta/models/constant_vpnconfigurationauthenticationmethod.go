package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VpnConfigurationAuthenticationMethod string

const (
	VpnConfigurationAuthenticationMethodazureAD             VpnConfigurationAuthenticationMethod = "AzureAD"
	VpnConfigurationAuthenticationMethodcertificate         VpnConfigurationAuthenticationMethod = "Certificate"
	VpnConfigurationAuthenticationMethodderivedCredential   VpnConfigurationAuthenticationMethod = "DerivedCredential"
	VpnConfigurationAuthenticationMethodsharedSecret        VpnConfigurationAuthenticationMethod = "SharedSecret"
	VpnConfigurationAuthenticationMethodusernameAndPassword VpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(VpnConfigurationAuthenticationMethodazureAD),
		string(VpnConfigurationAuthenticationMethodcertificate),
		string(VpnConfigurationAuthenticationMethodderivedCredential),
		string(VpnConfigurationAuthenticationMethodsharedSecret),
		string(VpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseVpnConfigurationAuthenticationMethod(input string) (*VpnConfigurationAuthenticationMethod, error) {
	vals := map[string]VpnConfigurationAuthenticationMethod{
		"azuread":             VpnConfigurationAuthenticationMethodazureAD,
		"certificate":         VpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   VpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        VpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": VpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := VpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
