package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EasEmailProfileConfigurationBaseUsernameSource string

const (
	EasEmailProfileConfigurationBaseUsernameSourceprimarySmtpAddress EasEmailProfileConfigurationBaseUsernameSource = "PrimarySmtpAddress"
	EasEmailProfileConfigurationBaseUsernameSourceuserPrincipalName  EasEmailProfileConfigurationBaseUsernameSource = "UserPrincipalName"
)

func PossibleValuesForEasEmailProfileConfigurationBaseUsernameSource() []string {
	return []string{
		string(EasEmailProfileConfigurationBaseUsernameSourceprimarySmtpAddress),
		string(EasEmailProfileConfigurationBaseUsernameSourceuserPrincipalName),
	}
}

func parseEasEmailProfileConfigurationBaseUsernameSource(input string) (*EasEmailProfileConfigurationBaseUsernameSource, error) {
	vals := map[string]EasEmailProfileConfigurationBaseUsernameSource{
		"primarysmtpaddress": EasEmailProfileConfigurationBaseUsernameSourceprimarySmtpAddress,
		"userprincipalname":  EasEmailProfileConfigurationBaseUsernameSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EasEmailProfileConfigurationBaseUsernameSource(input)
	return &out, nil
}
