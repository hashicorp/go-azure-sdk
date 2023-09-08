package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10VpnConfigurationAuthenticationMethod string

const (
	Windows10VpnConfigurationAuthenticationMethodcertificate         Windows10VpnConfigurationAuthenticationMethod = "Certificate"
	Windows10VpnConfigurationAuthenticationMethodcustomEapXml        Windows10VpnConfigurationAuthenticationMethod = "CustomEapXml"
	Windows10VpnConfigurationAuthenticationMethodderivedCredential   Windows10VpnConfigurationAuthenticationMethod = "DerivedCredential"
	Windows10VpnConfigurationAuthenticationMethodusernameAndPassword Windows10VpnConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForWindows10VpnConfigurationAuthenticationMethod() []string {
	return []string{
		string(Windows10VpnConfigurationAuthenticationMethodcertificate),
		string(Windows10VpnConfigurationAuthenticationMethodcustomEapXml),
		string(Windows10VpnConfigurationAuthenticationMethodderivedCredential),
		string(Windows10VpnConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseWindows10VpnConfigurationAuthenticationMethod(input string) (*Windows10VpnConfigurationAuthenticationMethod, error) {
	vals := map[string]Windows10VpnConfigurationAuthenticationMethod{
		"certificate":         Windows10VpnConfigurationAuthenticationMethodcertificate,
		"customeapxml":        Windows10VpnConfigurationAuthenticationMethodcustomEapXml,
		"derivedcredential":   Windows10VpnConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": Windows10VpnConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10VpnConfigurationAuthenticationMethod(input)
	return &out, nil
}
