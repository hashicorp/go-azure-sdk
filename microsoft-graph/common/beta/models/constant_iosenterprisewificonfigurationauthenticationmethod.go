package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	IosEnterpriseWiFiConfigurationAuthenticationMethodcertificate         IosEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	IosEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   IosEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	IosEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword IosEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForIosEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(IosEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(IosEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseIosEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*IosEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]IosEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         IosEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   IosEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": IosEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
