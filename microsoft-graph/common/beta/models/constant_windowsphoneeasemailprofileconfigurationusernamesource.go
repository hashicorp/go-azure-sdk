package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsPhoneEASEmailProfileConfigurationUsernameSource string

const (
	WindowsPhoneEASEmailProfileConfigurationUsernameSourceprimarySmtpAddress WindowsPhoneEASEmailProfileConfigurationUsernameSource = "PrimarySmtpAddress"
	WindowsPhoneEASEmailProfileConfigurationUsernameSourceuserPrincipalName  WindowsPhoneEASEmailProfileConfigurationUsernameSource = "UserPrincipalName"
)

func PossibleValuesForWindowsPhoneEASEmailProfileConfigurationUsernameSource() []string {
	return []string{
		string(WindowsPhoneEASEmailProfileConfigurationUsernameSourceprimarySmtpAddress),
		string(WindowsPhoneEASEmailProfileConfigurationUsernameSourceuserPrincipalName),
	}
}

func parseWindowsPhoneEASEmailProfileConfigurationUsernameSource(input string) (*WindowsPhoneEASEmailProfileConfigurationUsernameSource, error) {
	vals := map[string]WindowsPhoneEASEmailProfileConfigurationUsernameSource{
		"primarysmtpaddress": WindowsPhoneEASEmailProfileConfigurationUsernameSourceprimarySmtpAddress,
		"userprincipalname":  WindowsPhoneEASEmailProfileConfigurationUsernameSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := WindowsPhoneEASEmailProfileConfigurationUsernameSource(input)
	return &out, nil
}
