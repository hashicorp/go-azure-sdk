package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationAuthenticationMethod string

const (
	AndroidEasEmailProfileConfigurationAuthenticationMethodcertificate         AndroidEasEmailProfileConfigurationAuthenticationMethod = "Certificate"
	AndroidEasEmailProfileConfigurationAuthenticationMethodderivedCredential   AndroidEasEmailProfileConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidEasEmailProfileConfigurationAuthenticationMethodusernameAndPassword AndroidEasEmailProfileConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationAuthenticationMethodcertificate),
		string(AndroidEasEmailProfileConfigurationAuthenticationMethodderivedCredential),
		string(AndroidEasEmailProfileConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidEasEmailProfileConfigurationAuthenticationMethod(input string) (*AndroidEasEmailProfileConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationAuthenticationMethod{
		"certificate":         AndroidEasEmailProfileConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidEasEmailProfileConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidEasEmailProfileConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationAuthenticationMethod(input)
	return &out, nil
}
