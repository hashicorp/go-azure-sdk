package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod string

const (
	AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodcertificate         AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod = "Certificate"
	AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodderivedCredential   AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod = "DerivedCredential"
	AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodusernameAndPassword AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodcertificate),
		string(AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodderivedCredential),
		string(AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidWorkProfileEasEmailProfileBaseAuthenticationMethod(input string) (*AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod{
		"certificate":         AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodcertificate,
		"derivedcredential":   AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidWorkProfileEasEmailProfileBaseAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseAuthenticationMethod(input)
	return &out, nil
}
