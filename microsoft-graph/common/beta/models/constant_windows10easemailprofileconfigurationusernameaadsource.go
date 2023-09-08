package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10EasEmailProfileConfigurationUsernameAADSource string

const (
	Windows10EasEmailProfileConfigurationUsernameAADSourceprimarySmtpAddress Windows10EasEmailProfileConfigurationUsernameAADSource = "PrimarySmtpAddress"
	Windows10EasEmailProfileConfigurationUsernameAADSourcesamAccountName     Windows10EasEmailProfileConfigurationUsernameAADSource = "SamAccountName"
	Windows10EasEmailProfileConfigurationUsernameAADSourceuserPrincipalName  Windows10EasEmailProfileConfigurationUsernameAADSource = "UserPrincipalName"
)

func PossibleValuesForWindows10EasEmailProfileConfigurationUsernameAADSource() []string {
	return []string{
		string(Windows10EasEmailProfileConfigurationUsernameAADSourceprimarySmtpAddress),
		string(Windows10EasEmailProfileConfigurationUsernameAADSourcesamAccountName),
		string(Windows10EasEmailProfileConfigurationUsernameAADSourceuserPrincipalName),
	}
}

func parseWindows10EasEmailProfileConfigurationUsernameAADSource(input string) (*Windows10EasEmailProfileConfigurationUsernameAADSource, error) {
	vals := map[string]Windows10EasEmailProfileConfigurationUsernameAADSource{
		"primarysmtpaddress": Windows10EasEmailProfileConfigurationUsernameAADSourceprimarySmtpAddress,
		"samaccountname":     Windows10EasEmailProfileConfigurationUsernameAADSourcesamAccountName,
		"userprincipalname":  Windows10EasEmailProfileConfigurationUsernameAADSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10EasEmailProfileConfigurationUsernameAADSource(input)
	return &out, nil
}
