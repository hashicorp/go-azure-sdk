package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBaseUsernameAADSource string

const (
	EasEmailProfileConfigurationBaseUsernameAADSourceprimarySmtpAddress EasEmailProfileConfigurationBaseUsernameAADSource = "PrimarySmtpAddress"
	EasEmailProfileConfigurationBaseUsernameAADSourcesamAccountName     EasEmailProfileConfigurationBaseUsernameAADSource = "SamAccountName"
	EasEmailProfileConfigurationBaseUsernameAADSourceuserPrincipalName  EasEmailProfileConfigurationBaseUsernameAADSource = "UserPrincipalName"
)

func PossibleValuesForEasEmailProfileConfigurationBaseUsernameAADSource() []string {
	return []string{
		string(EasEmailProfileConfigurationBaseUsernameAADSourceprimarySmtpAddress),
		string(EasEmailProfileConfigurationBaseUsernameAADSourcesamAccountName),
		string(EasEmailProfileConfigurationBaseUsernameAADSourceuserPrincipalName),
	}
}

func parseEasEmailProfileConfigurationBaseUsernameAADSource(input string) (*EasEmailProfileConfigurationBaseUsernameAADSource, error) {
	vals := map[string]EasEmailProfileConfigurationBaseUsernameAADSource{
		"primarysmtpaddress": EasEmailProfileConfigurationBaseUsernameAADSourceprimarySmtpAddress,
		"samaccountname":     EasEmailProfileConfigurationBaseUsernameAADSourcesamAccountName,
		"userprincipalname":  EasEmailProfileConfigurationBaseUsernameAADSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasEmailProfileConfigurationBaseUsernameAADSource(input)
	return &out, nil
}
