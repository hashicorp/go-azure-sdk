package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileNineWorkEasConfigurationEmailAddressSource string

const (
	AndroidWorkProfileNineWorkEasConfigurationEmailAddressSourceprimarySmtpAddress AndroidWorkProfileNineWorkEasConfigurationEmailAddressSource = "PrimarySmtpAddress"
	AndroidWorkProfileNineWorkEasConfigurationEmailAddressSourceuserPrincipalName  AndroidWorkProfileNineWorkEasConfigurationEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForAndroidWorkProfileNineWorkEasConfigurationEmailAddressSource() []string {
	return []string{
		string(AndroidWorkProfileNineWorkEasConfigurationEmailAddressSourceprimarySmtpAddress),
		string(AndroidWorkProfileNineWorkEasConfigurationEmailAddressSourceuserPrincipalName),
	}
}

func parseAndroidWorkProfileNineWorkEasConfigurationEmailAddressSource(input string) (*AndroidWorkProfileNineWorkEasConfigurationEmailAddressSource, error) {
	vals := map[string]AndroidWorkProfileNineWorkEasConfigurationEmailAddressSource{
		"primarysmtpaddress": AndroidWorkProfileNineWorkEasConfigurationEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  AndroidWorkProfileNineWorkEasConfigurationEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileNineWorkEasConfigurationEmailAddressSource(input)
	return &out, nil
}
