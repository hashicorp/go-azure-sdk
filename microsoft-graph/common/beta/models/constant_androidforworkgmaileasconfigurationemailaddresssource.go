package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationEmailAddressSource string

const (
	AndroidForWorkGmailEasConfigurationEmailAddressSourceprimarySmtpAddress AndroidForWorkGmailEasConfigurationEmailAddressSource = "PrimarySmtpAddress"
	AndroidForWorkGmailEasConfigurationEmailAddressSourceuserPrincipalName  AndroidForWorkGmailEasConfigurationEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationEmailAddressSourceprimarySmtpAddress),
		string(AndroidForWorkGmailEasConfigurationEmailAddressSourceuserPrincipalName),
	}
}

func parseAndroidForWorkGmailEasConfigurationEmailAddressSource(input string) (*AndroidForWorkGmailEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidForWorkGmailEasConfigurationEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  AndroidForWorkGmailEasConfigurationEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationEmailAddressSource(input)
	return &out, nil
}
