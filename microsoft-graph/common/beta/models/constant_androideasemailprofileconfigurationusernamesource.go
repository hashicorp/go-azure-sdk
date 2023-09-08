package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidEasEmailProfileConfigurationUsernameSource string

const (
	AndroidEasEmailProfileConfigurationUsernameSourceprimarySmtpAddress AndroidEasEmailProfileConfigurationUsernameSource = "PrimarySmtpAddress"
	AndroidEasEmailProfileConfigurationUsernameSourcesamAccountName     AndroidEasEmailProfileConfigurationUsernameSource = "SamAccountName"
	AndroidEasEmailProfileConfigurationUsernameSourceuserPrincipalName  AndroidEasEmailProfileConfigurationUsernameSource = "UserPrincipalName"
	AndroidEasEmailProfileConfigurationUsernameSourceusername           AndroidEasEmailProfileConfigurationUsernameSource = "Username"
)

func PossibleValuesForAndroidEasEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(AndroidEasEmailProfileConfigurationUsernameSourceprimarySmtpAddress),
		string(AndroidEasEmailProfileConfigurationUsernameSourcesamAccountName),
		string(AndroidEasEmailProfileConfigurationUsernameSourceuserPrincipalName),
		string(AndroidEasEmailProfileConfigurationUsernameSourceusername),
	}
}

func parseAndroidEasEmailProfileConfigurationUsernameSource(input string) (*AndroidEasEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]AndroidEasEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": AndroidEasEmailProfileConfigurationUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidEasEmailProfileConfigurationUsernameSourcesamAccountName,
		"userprincipalname":  AndroidEasEmailProfileConfigurationUsernameSourceuserPrincipalName,
		"username":           AndroidEasEmailProfileConfigurationUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidEasEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
