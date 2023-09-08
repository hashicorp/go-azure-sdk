package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSWiredNetworkConfigurationAuthenticationMethod string

const (
	MacOSWiredNetworkConfigurationAuthenticationMethodcertificate         MacOSWiredNetworkConfigurationAuthenticationMethod = "Certificate"
	MacOSWiredNetworkConfigurationAuthenticationMethodderivedCredential   MacOSWiredNetworkConfigurationAuthenticationMethod = "DerivedCredential"
	MacOSWiredNetworkConfigurationAuthenticationMethodusernameAndPassword MacOSWiredNetworkConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForMacOSWiredNetworkConfigurationAuthenticationMethod() []string {
	return []string{
		string(MacOSWiredNetworkConfigurationAuthenticationMethodcertificate),
		string(MacOSWiredNetworkConfigurationAuthenticationMethodderivedCredential),
		string(MacOSWiredNetworkConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseMacOSWiredNetworkConfigurationAuthenticationMethod(input string) (*MacOSWiredNetworkConfigurationAuthenticationMethod, error) {
	vals := map[string]MacOSWiredNetworkConfigurationAuthenticationMethod{
		"certificate":         MacOSWiredNetworkConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   MacOSWiredNetworkConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": MacOSWiredNetworkConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSWiredNetworkConfigurationAuthenticationMethod(input)
	return &out, nil
}
