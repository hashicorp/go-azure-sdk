package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodcertificate         AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
