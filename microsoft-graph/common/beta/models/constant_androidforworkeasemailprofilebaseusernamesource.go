package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidForWorkEasEmailProfileBaseUsernameSource string

const (
	AndroidForWorkEasEmailProfileBaseUsernameSourceprimarySmtpAddress AndroidForWorkEasEmailProfileBaseUsernameSource = "PrimarySmtpAddress"
	AndroidForWorkEasEmailProfileBaseUsernameSourcesamAccountName     AndroidForWorkEasEmailProfileBaseUsernameSource = "SamAccountName"
	AndroidForWorkEasEmailProfileBaseUsernameSourceuserPrincipalName  AndroidForWorkEasEmailProfileBaseUsernameSource = "UserPrincipalName"
	AndroidForWorkEasEmailProfileBaseUsernameSourceusername           AndroidForWorkEasEmailProfileBaseUsernameSource = "Username"
)

func PossibleValuesForAndroidForWorkEasEmailProfileBaseUsernameSource() []string {
	return []string{
		string(AndroidForWorkEasEmailProfileBaseUsernameSourceprimarySmtpAddress),
		string(AndroidForWorkEasEmailProfileBaseUsernameSourcesamAccountName),
		string(AndroidForWorkEasEmailProfileBaseUsernameSourceuserPrincipalName),
		string(AndroidForWorkEasEmailProfileBaseUsernameSourceusername),
	}
}

func parseAndroidForWorkEasEmailProfileBaseUsernameSource(input string) (*AndroidForWorkEasEmailProfileBaseUsernameSource, error) {
	vals := map[string]AndroidForWorkEasEmailProfileBaseUsernameSource{
		"primarysmtpaddress": AndroidForWorkEasEmailProfileBaseUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidForWorkEasEmailProfileBaseUsernameSourcesamAccountName,
		"userprincipalname":  AndroidForWorkEasEmailProfileBaseUsernameSourceuserPrincipalName,
		"username":           AndroidForWorkEasEmailProfileBaseUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidForWorkEasEmailProfileBaseUsernameSource(input)
	return &out, nil
}
