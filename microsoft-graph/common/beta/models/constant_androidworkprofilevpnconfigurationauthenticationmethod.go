package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileVpnConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileVpnConfigurationAuthenticationMethodazureAD             AndroidWorkProfileVpnConfigurationAuthenticationMethod = "AzureAD"
	AndroidWorkProfileVpnConfigurationAuthenticationMethodcertificate         AndroidWorkProfileVpnConfigurationAuthenticationMethod = "Certificate"
	AndroidWorkProfileVpnConfigurationAuthenticationMethodderivedCredential   AndroidWorkProfileVpnConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidWorkProfileVpnConfigurationAuthenticationMethodsharedSecret        AndroidWorkProfileVpnConfigurationAuthenticationMethod = "SharedSecret"
	AndroidWorkProfileVpnConfigurationAuthenticationMethodusernameAndPassword AndroidWorkProfileVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethodazureAD),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethodcertificate),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethodderivedCredential),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethodsharedSecret),
		string(AndroidWorkProfileVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidWorkProfileVpnConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidWorkProfileVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         AndroidWorkProfileVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidWorkProfileVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        AndroidWorkProfileVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": AndroidWorkProfileVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
