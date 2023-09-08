package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationUsernameSource string

const (
	AndroidForWorkNineWorkEasConfigurationUsernameSourceprimarySmtpAddress AndroidForWorkNineWorkEasConfigurationUsernameSource = "PrimarySmtpAddress"
	AndroidForWorkNineWorkEasConfigurationUsernameSourcesamAccountName     AndroidForWorkNineWorkEasConfigurationUsernameSource = "SamAccountName"
	AndroidForWorkNineWorkEasConfigurationUsernameSourceuserPrincipalName  AndroidForWorkNineWorkEasConfigurationUsernameSource = "UserPrincipalName"
	AndroidForWorkNineWorkEasConfigurationUsernameSourceusername           AndroidForWorkNineWorkEasConfigurationUsernameSource = "Username"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationUsernameSourceprimarySmtpAddress),
		string(AndroidForWorkNineWorkEasConfigurationUsernameSourcesamAccountName),
		string(AndroidForWorkNineWorkEasConfigurationUsernameSourceuserPrincipalName),
		string(AndroidForWorkNineWorkEasConfigurationUsernameSourceusername),
	}
}

func parseAndroidForWorkNineWorkEasConfigurationUsernameSource(input string) (*AndroidForWorkNineWorkEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidForWorkNineWorkEasConfigurationUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidForWorkNineWorkEasConfigurationUsernameSourcesamAccountName,
		"userprincipalname":  AndroidForWorkNineWorkEasConfigurationUsernameSourceuserPrincipalName,
		"username":           AndroidForWorkNineWorkEasConfigurationUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationUsernameSource(input)
	return &out, nil
}
