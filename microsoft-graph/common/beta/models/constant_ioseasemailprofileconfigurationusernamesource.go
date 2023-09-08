package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationUsernameSource string

const (
	IosEasEmailProfileConfigurationUsernameSourceprimarySmtpAddress IosEasEmailProfileConfigurationUsernameSource = "PrimarySmtpAddress"
	IosEasEmailProfileConfigurationUsernameSourceuserPrincipalName  IosEasEmailProfileConfigurationUsernameSource = "UserPrincipalName"
)

func PossibleValuesForIosEasEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationUsernameSourceprimarySmtpAddress),
		string(IosEasEmailProfileConfigurationUsernameSourceuserPrincipalName),
	}
}

func parseIosEasEmailProfileConfigurationUsernameSource(input string) (*IosEasEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": IosEasEmailProfileConfigurationUsernameSourceprimarySmtpAddress,
		"userprincipalname":  IosEasEmailProfileConfigurationUsernameSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
