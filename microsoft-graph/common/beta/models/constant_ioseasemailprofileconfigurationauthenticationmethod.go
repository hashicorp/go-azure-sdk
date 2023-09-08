package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationAuthenticationMethod string

const (
	IosEasEmailProfileConfigurationAuthenticationMethodcertificate         IosEasEmailProfileConfigurationAuthenticationMethod = "Certificate"
	IosEasEmailProfileConfigurationAuthenticationMethodderivedCredential   IosEasEmailProfileConfigurationAuthenticationMethod = "DerivedCredential"
	IosEasEmailProfileConfigurationAuthenticationMethodusernameAndPassword IosEasEmailProfileConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForIosEasEmailProfileConfigurationAuthenticationMethod() []string {
	return []string{
		string(IosEasEmailProfileConfigurationAuthenticationMethodcertificate),
		string(IosEasEmailProfileConfigurationAuthenticationMethodderivedCredential),
		string(IosEasEmailProfileConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseIosEasEmailProfileConfigurationAuthenticationMethod(input string) (*IosEasEmailProfileConfigurationAuthenticationMethod, error) {
	vals := map[string]IosEasEmailProfileConfigurationAuthenticationMethod{
		"certificate":         IosEasEmailProfileConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   IosEasEmailProfileConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": IosEasEmailProfileConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationAuthenticationMethod(input)
	return &out, nil
}
