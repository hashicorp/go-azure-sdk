package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodcertificate         AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
