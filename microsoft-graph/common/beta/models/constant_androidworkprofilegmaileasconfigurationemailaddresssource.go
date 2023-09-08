package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationEmailAddressSource string

const (
	AndroidWorkProfileGmailEasConfigurationEmailAddressSourceprimarySmtpAddress AndroidWorkProfileGmailEasConfigurationEmailAddressSource = "PrimarySmtpAddress"
	AndroidWorkProfileGmailEasConfigurationEmailAddressSourceuserPrincipalName  AndroidWorkProfileGmailEasConfigurationEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationEmailAddressSourceprimarySmtpAddress),
		string(AndroidWorkProfileGmailEasConfigurationEmailAddressSourceuserPrincipalName),
	}
}

func parseAndroidWorkProfileGmailEasConfigurationEmailAddressSource(input string) (*AndroidWorkProfileGmailEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidWorkProfileGmailEasConfigurationEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  AndroidWorkProfileGmailEasConfigurationEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationEmailAddressSource(input)
	return &out, nil
}
