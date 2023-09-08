package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseAuthenticationMethod string

const (
	AndroidForWorkEasEmailProfileBaseAuthenticationMethodcertificate         AndroidForWorkEasEmailProfileBaseAuthenticationMethod = "Certificate"
	AndroidForWorkEasEmailProfileBaseAuthenticationMethodderivedCredential   AndroidForWorkEasEmailProfileBaseAuthenticationMethod = "DerivedCredential"
	AndroidForWorkEasEmailProfileBaseAuthenticationMethodusernameAndPassword AndroidForWorkEasEmailProfileBaseAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseAuthenticationMethodcertificate),
		string(AndroidForWorkEasEmailProfileBaseAuthenticationMethodderivedCredential),
		string(AndroidForWorkEasEmailProfileBaseAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidForWorkEasEmailProfileBaseAuthenticationMethod(input string) (*AndroidForWorkEasEmailProfileBaseAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseAuthenticationMethod{
		"certificate":         AndroidForWorkEasEmailProfileBaseAuthenticationMethodcertificate,
		"derivedcredential":   AndroidForWorkEasEmailProfileBaseAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidForWorkEasEmailProfileBaseAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseAuthenticationMethod(input)
	return &out, nil
}
