package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileGmailEasConfigurationUsernameSource string

const (
	AndroidWorkProfileGmailEasConfigurationUsernameSourceprimarySmtpAddress AndroidWorkProfileGmailEasConfigurationUsernameSource = "PrimarySmtpAddress"
	AndroidWorkProfileGmailEasConfigurationUsernameSourcesamAccountName     AndroidWorkProfileGmailEasConfigurationUsernameSource = "SamAccountName"
	AndroidWorkProfileGmailEasConfigurationUsernameSourceuserPrincipalName  AndroidWorkProfileGmailEasConfigurationUsernameSource = "UserPrincipalName"
	AndroidWorkProfileGmailEasConfigurationUsernameSourceusername           AndroidWorkProfileGmailEasConfigurationUsernameSource = "Username"
)

func PossibleValuesForAndroidWorkProfileGmailEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidWorkProfileGmailEasConfigurationUsernameSourceprimarySmtpAddress),
		string(AndroidWorkProfileGmailEasConfigurationUsernameSourcesamAccountName),
		string(AndroidWorkProfileGmailEasConfigurationUsernameSourceuserPrincipalName),
		string(AndroidWorkProfileGmailEasConfigurationUsernameSourceusername),
	}
}

func parseAndroidWorkProfileGmailEasConfigurationUsernameSource(input string) (*AndroidWorkProfileGmailEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidWorkProfileGmailEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidWorkProfileGmailEasConfigurationUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidWorkProfileGmailEasConfigurationUsernameSourcesamAccountName,
		"userprincipalname":  AndroidWorkProfileGmailEasConfigurationUsernameSourceuserPrincipalName,
		"username":           AndroidWorkProfileGmailEasConfigurationUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileGmailEasConfigurationUsernameSource(input)
	return &out, nil
}
