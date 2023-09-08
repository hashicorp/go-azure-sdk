package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationAuthenticationMethod string

const (
	WindowsWiredNetworkConfigurationAuthenticationMethodcertificate         WindowsWiredNetworkConfigurationAuthenticationMethod = "Certificate"
	WindowsWiredNetworkConfigurationAuthenticationMethodderivedCredential   WindowsWiredNetworkConfigurationAuthenticationMethod = "DerivedCredential"
	WindowsWiredNetworkConfigurationAuthenticationMethodusernameAndPassword WindowsWiredNetworkConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForWindowsWiredNetworkConfigurationAuthenticationMethod() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationAuthenticationMethodcertificate),
		string(WindowsWiredNetworkConfigurationAuthenticationMethodderivedCredential),
		string(WindowsWiredNetworkConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseWindowsWiredNetworkConfigurationAuthenticationMethod(input string) (*WindowsWiredNetworkConfigurationAuthenticationMethod, error) {
	vals := map[string]WindowsWiredNetworkConfigurationAuthenticationMethod{
		"certificate":         WindowsWiredNetworkConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   WindowsWiredNetworkConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": WindowsWiredNetworkConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationAuthenticationMethod(input)
	return &out, nil
}
