package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidEnterpriseWiFiConfigurationAuthenticationMethodcertificate         AndroidEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	AndroidEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   AndroidEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword AndroidEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(AndroidEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(AndroidEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
