package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationAuthenticationMethod string

const (
	AndroidForWorkNineWorkEasConfigurationAuthenticationMethodcertificate         AndroidForWorkNineWorkEasConfigurationAuthenticationMethod = "Certificate"
	AndroidForWorkNineWorkEasConfigurationAuthenticationMethodderivedCredential   AndroidForWorkNineWorkEasConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidForWorkNineWorkEasConfigurationAuthenticationMethodusernameAndPassword AndroidForWorkNineWorkEasConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationAuthenticationMethodcertificate),
		string(AndroidForWorkNineWorkEasConfigurationAuthenticationMethodderivedCredential),
		string(AndroidForWorkNineWorkEasConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidForWorkNineWorkEasConfigurationAuthenticationMethod(input string) (*AndroidForWorkNineWorkEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationAuthenticationMethod{
		"certificate":         AndroidForWorkNineWorkEasConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidForWorkNineWorkEasConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidForWorkNineWorkEasConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
