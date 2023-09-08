package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationAuthenticationMethod string

const (
	AndroidForWorkGmailEasConfigurationAuthenticationMethodcertificate         AndroidForWorkGmailEasConfigurationAuthenticationMethod = "Certificate"
	AndroidForWorkGmailEasConfigurationAuthenticationMethodderivedCredential   AndroidForWorkGmailEasConfigurationAuthenticationMethod = "DerivedCredential"
	AndroidForWorkGmailEasConfigurationAuthenticationMethodusernameAndPassword AndroidForWorkGmailEasConfigurationAuthenticationMethod = "UsernameAndPassword"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationAuthenticationMethod() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationAuthenticationMethodcertificate),
		string(AndroidForWorkGmailEasConfigurationAuthenticationMethodderivedCredential),
		string(AndroidForWorkGmailEasConfigurationAuthenticationMethodusernameAndPassword),
	}
}

func parseAndroidForWorkGmailEasConfigurationAuthenticationMethod(input string) (*AndroidForWorkGmailEasConfigurationAuthenticationMethod, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationAuthenticationMethod{
		"certificate":         AndroidForWorkGmailEasConfigurationAuthenticationMethodcertificate,
		"derivedcredential":   AndroidForWorkGmailEasConfigurationAuthenticationMethodderivedCredential,
		"usernameandpassword": AndroidForWorkGmailEasConfigurationAuthenticationMethodusernameAndPassword,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationAuthenticationMethod(input)
	return &out, nil
}
