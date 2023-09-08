package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseEmailAddressSource string

const (
	AndroidForWorkEasEmailProfileBaseEmailAddressSourceprimarySmtpAddress AndroidForWorkEasEmailProfileBaseEmailAddressSource = "PrimarySmtpAddress"
	AndroidForWorkEasEmailProfileBaseEmailAddressSourceuserPrincipalName  AndroidForWorkEasEmailProfileBaseEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseEmailAddressSource() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseEmailAddressSourceprimarySmtpAddress),
		string(AndroidForWorkEasEmailProfileBaseEmailAddressSourceuserPrincipalName),
	}
}

func parseAndroidForWorkEasEmailProfileBaseEmailAddressSource(input string) (*AndroidForWorkEasEmailProfileBaseEmailAddressSource, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseEmailAddressSource{
		"primarysmtpaddress": AndroidForWorkEasEmailProfileBaseEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  AndroidForWorkEasEmailProfileBaseEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseEmailAddressSource(input)
	return &out, nil
}
