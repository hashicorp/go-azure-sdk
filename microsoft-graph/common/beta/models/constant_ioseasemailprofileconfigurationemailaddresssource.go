package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosEasEmailProfileConfigurationEmailAddressSource string

const (
	IosEasEmailProfileConfigurationEmailAddressSourceprimarySmtpAddress IosEasEmailProfileConfigurationEmailAddressSource = "PrimarySmtpAddress"
	IosEasEmailProfileConfigurationEmailAddressSourceuserPrincipalName  IosEasEmailProfileConfigurationEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForIosEasEmailProfileConfigurationEmailAddressSource() []string {
	return []string{
		string(IosEasEmailProfileConfigurationEmailAddressSourceprimarySmtpAddress),
		string(IosEasEmailProfileConfigurationEmailAddressSourceuserPrincipalName),
	}
}

func parseIosEasEmailProfileConfigurationEmailAddressSource(input string) (*IosEasEmailProfileConfigurationEmailAddressSource, error) {
	vals := map[string]IosEasEmailProfileConfigurationEmailAddressSource{
		"primarysmtpaddress": IosEasEmailProfileConfigurationEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  IosEasEmailProfileConfigurationEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := IosEasEmailProfileConfigurationEmailAddressSource(input)
	return &out, nil
}
