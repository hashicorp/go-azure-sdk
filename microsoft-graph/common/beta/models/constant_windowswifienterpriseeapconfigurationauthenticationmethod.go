package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod string

const (
	WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodcertificate         WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod = "Certificate"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodderivedCredential   WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod = "DerivedCredential"
	WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodusernameAndPassword WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForWindowsWifiEnterpriseEAPConfigurationAuthenticationMethod() []string {
	return []string{
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodcertificate),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodderivedCredential),
		string(WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseWindowsWifiEnterpriseEAPConfigurationAuthenticationMethod(input string) (*WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod, error) {
	vals := map[string]WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod{
		"certificate":         WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": WindowsWifiEnterpriseEAPConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWifiEnterpriseEAPConfigurationAuthenticationMethod(input)
	return &out, nil
}
