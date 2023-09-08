package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationUsernameAADSource string

const (
	IosEasEmailProfileConfigurationUsernameAADSourceprimarySmtpAddress IosEasEmailProfileConfigurationUsernameAADSource = "PrimarySmtpAddress"
	IosEasEmailProfileConfigurationUsernameAADSourcesamAccountName     IosEasEmailProfileConfigurationUsernameAADSource = "SamAccountName"
	IosEasEmailProfileConfigurationUsernameAADSourceuserPrincipalName  IosEasEmailProfileConfigurationUsernameAADSource = "UserPrincipalName"
)

func PossibleValuesForIosEasEmailProfileConfigurationUsernameAADSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationUsernameAADSourceprimarySmtpAddress),
		string(IosEasEmailProfileConfigurationUsernameAADSourcesamAccountName),
		string(IosEasEmailProfileConfigurationUsernameAADSourceuserPrincipalName),
	}
}

func parseIosEasEmailProfileConfigurationUsernameAADSource(input string) (*IosEasEmailProfileConfigurationUsernameAADSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationUsernameAADSource{
		"primarysmtpaddress": IosEasEmailProfileConfigurationUsernameAADSourceprimarySmtpAddress,
		"samaccountname":     IosEasEmailProfileConfigurationUsernameAADSourcesamAccountName,
		"userprincipalname":  IosEasEmailProfileConfigurationUsernameAADSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationUsernameAADSource(input)
	return &out, nil
}
