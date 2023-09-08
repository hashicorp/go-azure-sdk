package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationAuthenticationMethod string

const (
	AndroidWorkProfileGmailEasConfigurationAuthenticationMethodcertificate         AndroidWorkProfileGmailEasConfigurationAuthenticationMethod = "Certificate"
	AndroidWorkProfileGmailEasConfigurationAuthenticationMethodderivedCredential   AndroidWorkProfileGmailEasConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidWorkProfileGmailEasConfigurationAuthenticationMethodusernameAndPassword AndroidWorkProfileGmailEasConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationAuthenticationMethodcertificate),
		string(AndroidWorkProfileGmailEasConfigurationAuthenticationMethodderivedCredential),
		string(AndroidWorkProfileGmailEasConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidWorkProfileGmailEasConfigurationAuthenticationMethod(input string) (*AndroidWorkProfileGmailEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationAuthenticationMethod{
		"certificate":         AndroidWorkProfileGmailEasConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidWorkProfileGmailEasConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidWorkProfileGmailEasConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
