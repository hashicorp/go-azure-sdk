package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseEmailAddressSource string

const (
	AndroidWorkProfileEasEmailProfileBaseEmailAddressSourceprimarySmtpAddress AndroidWorkProfileEasEmailProfileBaseEmailAddressSource = "PrimarySmtpAddress"
	AndroidWorkProfileEasEmailProfileBaseEmailAddressSourceuserPrincipalName  AndroidWorkProfileEasEmailProfileBaseEmailAddressSource = "UserPrincipalName"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseEmailAddressSource() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseEmailAddressSourceprimarySmtpAddress),
		string(AndroidWorkProfileEasEmailProfileBaseEmailAddressSourceuserPrincipalName),
	}
}

func parseAndroidWorkProfileEasEmailProfileBaseEmailAddressSource(input string) (*AndroidWorkProfileEasEmailProfileBaseEmailAddressSource, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseEmailAddressSource{
		"primarysmtpaddress": AndroidWorkProfileEasEmailProfileBaseEmailAddressSourceprimarySmtpAddress,
		"userprincipalname":  AndroidWorkProfileEasEmailProfileBaseEmailAddressSourceuserPrincipalName,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseEmailAddressSource(input)
	return &out, nil
}
