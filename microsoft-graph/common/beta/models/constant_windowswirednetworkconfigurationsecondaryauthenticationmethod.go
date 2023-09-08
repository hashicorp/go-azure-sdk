package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod string

const (
	WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodcertificate         WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod = "Certificate"
	WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodderivedCredential   WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod = "DerivedCredential"
	WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodusernameAndPassword WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForWindowsWiredNetworkConfigurationSecondaryAuthenticationMethod() []string {
	return []string{
		string(WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodcertificate),
		string(WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodderivedCredential),
		string(WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodusernameAndPassword),
	}
}

func parseWindowsWiredNetworkConfigurationSecondaryAuthenticationMethod(input string) (*WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod, error) {
	vals := map[string]WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod{
		"certificate":         WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodcertificate,
		"derivedcredential":   WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodderivedCredential,
		"usernameandpassword": WindowsWiredNetworkConfigurationSecondaryAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsWiredNetworkConfigurationSecondaryAuthenticationMethod(input)
	return &out, nil
}
