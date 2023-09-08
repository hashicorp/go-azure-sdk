package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerVpnConfigurationAuthenticationMethod string

const (
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethodazureAD             AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "AzureAD"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethodcertificate         AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "Certificate"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethodderivedCredential   AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethodsharedSecret        AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "SharedSecret"
	AndroidDeviceOwnerVpnConfigurationAuthenticationMethodusernameAndPassword AndroidDeviceOwnerVpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidDeviceOwnerVpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethodazureAD),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethodcertificate),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethodderivedCredential),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethodsharedSecret),
		string(AndroidDeviceOwnerVpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidDeviceOwnerVpnConfigurationAuthenticationMethod(input string) (*AndroidDeviceOwnerVpnConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidDeviceOwnerVpnConfigurationAuthenticationMethod{
		"azuread":             AndroidDeviceOwnerVpnConfigurationAuthenticationMethodazureAD,
		"certificate":         AndroidDeviceOwnerVpnConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidDeviceOwnerVpnConfigurationAuthenticationMethodderivedCredential,
		"sharedsecret":        AndroidDeviceOwnerVpnConfigurationAuthenticationMethodsharedSecret,
		"usernameandpassword": AndroidDeviceOwnerVpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerVpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
