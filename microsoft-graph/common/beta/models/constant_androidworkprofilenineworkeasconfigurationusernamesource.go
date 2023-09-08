package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationUsernameSource string

const (
	AndroidWorkProfileNineWorkEasConfigurationUsernameSourceprimarySmtpAddress AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "PrimarySmtpAddress"
	AndroidWorkProfileNineWorkEasConfigurationUsernameSourcesamAccountName     AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "SamAccountName"
	AndroidWorkProfileNineWorkEasConfigurationUsernameSourceuserPrincipalName  AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "UserPrincipalName"
	AndroidWorkProfileNineWorkEasConfigurationUsernameSourceusername           AndroidWorkProfileNineWorkEasConfigurationUsernameSource = "Username"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationUsernameSource() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSourceprimarySmtpAddress),
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSourcesamAccountName),
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSourceuserPrincipalName),
		string(AndroidWorkProfileNineWorkEasConfigurationUsernameSourceusername),
	}
}

func parseAndroidWorkProfileNineWorkEasConfigurationUsernameSource(input string) (*AndroidWorkProfileNineWorkEasConfigurationUsernameSource, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationUsernameSource{
		"primarysmtpaddress": AndroidWorkProfileNineWorkEasConfigurationUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidWorkProfileNineWorkEasConfigurationUsernameSourcesamAccountName,
		"userprincipalname":  AndroidWorkProfileNineWorkEasConfigurationUsernameSourceuserPrincipalName,
		"username":           AndroidWorkProfileNineWorkEasConfigurationUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationUsernameSource(input)
	return &out, nil
}
