package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkGmailEasConfigurationUsernameSource string

const (
	AndroidForWorkGmailEasConfigurationUsernameSourceprimarySmtpAddress AndroidForWorkGmailEasConfigurationUsernameSource = "PrimarySmtpAddress"
	AndroidForWorkGmailEasConfigurationUsernameSourcesamAccountName     AndroidForWorkGmailEasConfigurationUsernameSource = "SamAccountName"
	AndroidForWorkGmailEasConfigurationUsernameSourceuserPrincipalName  AndroidForWorkGmailEasConfigurationUsernameSource = "UserPrincipalName"
	AndroidForWorkGmailEasConfigurationUsernameSourceusername           AndroidForWorkGmailEasConfigurationUsernameSource = "Username"
)

func PossibleValuesForAndroidForWorkGmailEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidForWorkGmailEasConfigurationUsernameSourceprimarySmtpAddress),
		string(AndroidForWorkGmailEasConfigurationUsernameSourcesamAccountName),
		string(AndroidForWorkGmailEasConfigurationUsernameSourceuserPrincipalName),
		string(AndroidForWorkGmailEasConfigurationUsernameSourceusername),
	}
}

func parseAndroidForWorkGmailEasConfigurationUsernameSource(input string) (*AndroidForWorkGmailEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidForWorkGmailEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidForWorkGmailEasConfigurationUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidForWorkGmailEasConfigurationUsernameSourcesamAccountName,
		"userprincipalname":  AndroidForWorkGmailEasConfigurationUsernameSourceuserPrincipalName,
		"username":           AndroidForWorkGmailEasConfigurationUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkGmailEasConfigurationUsernameSource(input)
	return &out, nil
}
