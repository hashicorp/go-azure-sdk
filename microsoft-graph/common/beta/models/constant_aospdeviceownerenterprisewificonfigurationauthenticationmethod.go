package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod string

const (
	AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodcertificate         AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "Certificate"
	AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential   AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "DerivedCredential"
	AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod() []string {
	return []string{
		string(AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodcertificate),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential),
		string(AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input string) (*AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod, error) {
	vals := map[string]AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod{
		"certificate":         AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AospDeviceOwnerEnterpriseWiFiConfigurationAuthenticationMethod(input)
	return &out, nil
}
