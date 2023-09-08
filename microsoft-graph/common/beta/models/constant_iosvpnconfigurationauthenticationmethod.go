package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosVpnConfigurationAuthenticationMethod string

const (
	IosVpnConfigurationAuthenticationMethodazureAD             IosVpnConfigurationAuthenticationMethod = "AzureAD"
	IosVpnConfigurationAuthenticationMethodcertificate         IosVpnConfigurationAuthenticationMethod = "Certificate"
	IosVpnConfigurationAuthenticationMethodderivedCredential   IosVpnConfigurationAuthenticationMethod = "DerivedCredential"
	IosVpnConfigurationAuthenticationMethodsharedSecret        IosVpnConfigurationAuthenticationMethod = "SharedSecret"
	IosVpnConfigurationAuthenticationMethodusernameAndPassword IosVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForIosVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosVpnConfigurationAuthenticationMethodazureAD),
		string(IosVpnConfigurationAuthenticationMethodcertificate),
		string(IosVpnConfigurationAuthenticationMethodderivedCredential),
		string(IosVpnConfigurationAuthenticationMethodsharedSecret),
		string(IosVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseIosVpnConfigurationAuthenticationMethod(input string) (*IosVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]IosVpnConfigurationAuthenticationMethod{
		"azuread":             IosVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         IosVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   IosVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        IosVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": IosVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
