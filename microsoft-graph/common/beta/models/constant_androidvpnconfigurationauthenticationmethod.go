package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidVpnConfigurationAuthenticationMethod string

const (
	AndroidVpnConfigurationAuthenticationMethodazureAD             AndroidVpnConfigurationAuthenticationMethod = "AzureAD"
	AndroidVpnConfigurationAuthenticationMethodcertificate         AndroidVpnConfigurationAuthenticationMethod = "Certificate"
	AndroidVpnConfigurationAuthenticationMethodderivedCredential   AndroidVpnConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidVpnConfigurationAuthenticationMethodsharedSecret        AndroidVpnConfigurationAuthenticationMethod = "SharedSecret"
	AndroidVpnConfigurationAuthenticationMethodusernameAndPassword AndroidVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidVpnConfigurationAuthenticationMethodazureAD),
		string(AndroidVpnConfigurationAuthenticationMethodcertificate),
		string(AndroidVpnConfigurationAuthenticationMethodderivedCredential),
		string(AndroidVpnConfigurationAuthenticationMethodsharedSecret),
		string(AndroidVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidVpnConfigurationAuthenticationMethod(input string) (*AndroidVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         AndroidVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        AndroidVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": AndroidVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
