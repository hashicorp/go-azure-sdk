package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodcertificate         AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
