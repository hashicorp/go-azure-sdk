package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkNineWorkEasConfigurationEmailAddressSource string

const (
	AndroidForWorkNineWorkEasConfigurationEmailAddressSourceprimarySmtpAddress AndroidForWorkNineWorkEasConfigurationEmailAddressSource = "PrimarySmtpAddress"
	AndroidForWorkNineWorkEasConfigurationEmailAddressSourceuserPrincipalName  AndroidForWorkNineWorkEasConfigurationEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkNineWorkEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidForWorkNineWorkEasConfigurationEmailAddressSourceprimarySmtpAddress),
		string(AndroidForWorkNineWorkEasConfigurationEmailAddressSourceuserPrincipalName),
	}
}

func parseAndroidForWorkNineWorkEasConfigurationEmailAddressSource(input string) (*AndroidForWorkNineWorkEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidForWorkNineWorkEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidForWorkNineWorkEasConfigurationEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  AndroidForWorkNineWorkEasConfigurationEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkNineWorkEasConfigurationEmailAddressSource(input)
	return &out, nil
}
