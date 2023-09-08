package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	MacOSEnterpriseWiFiConfigurationAuthenticationMethodcertificate         MacOSEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	MacOSEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   MacOSEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	MacOSEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword MacOSEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForMacOSEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(MacOSEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(MacOSEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(MacOSEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseMacOSEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*MacOSEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]MacOSEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         MacOSEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   MacOSEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": MacOSEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
