package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkVpnConfigurationAuthenticationMethod string

const (
	AndroidForWorkVpnConfigurationAuthenticationMethodazureAD             AndroidForWorkVpnConfigurationAuthenticationMethod = "AzureAD"
	AndroidForWorkVpnConfigurationAuthenticationMethodcertificate         AndroidForWorkVpnConfigurationAuthenticationMethod = "Certificate"
	AndroidForWorkVpnConfigurationAuthenticationMethodderivedCredential   AndroidForWorkVpnConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidForWorkVpnConfigurationAuthenticationMethodsharedSecret        AndroidForWorkVpnConfigurationAuthenticationMethod = "SharedSecret"
	AndroidForWorkVpnConfigurationAuthenticationMethodusernameAndPassword AndroidForWorkVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidForWorkVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkVpnConfigurationAuthenticationMethodazureAD),
		string(AndroidForWorkVpnConfigurationAuthenticationMethodcertificate),
		string(AndroidForWorkVpnConfigurationAuthenticationMethodderivedCredential),
		string(AndroidForWorkVpnConfigurationAuthenticationMethodsharedSecret),
		string(AndroidForWorkVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidForWorkVpnConfigurationAuthenticationMethod(input string) (*AndroidForWorkVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidForWorkVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         AndroidForWorkVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidForWorkVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        AndroidForWorkVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": AndroidForWorkVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
