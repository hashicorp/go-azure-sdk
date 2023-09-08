package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodcertificate         AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod = "Certificate"
	AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodderivedCredential   AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodusernameAndPassword AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodcertificate),
		string(AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodderivedCredential),
		string(AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod{
		"certificate":         AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
