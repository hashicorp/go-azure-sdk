package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidWorkProfileEasEmailProfileBaseUsernameSource string

const (
	AndroidWorkProfileEasEmailProfileBaseUsernameSourceprimarySmtpAddress AndroidWorkProfileEasEmailProfileBaseUsernameSource = "PrimarySmtpAddress"
	AndroidWorkProfileEasEmailProfileBaseUsernameSourcesamAccountName     AndroidWorkProfileEasEmailProfileBaseUsernameSource = "SamAccountName"
	AndroidWorkProfileEasEmailProfileBaseUsernameSourceuserPrincipalName  AndroidWorkProfileEasEmailProfileBaseUsernameSource = "UserPrincipalName"
	AndroidWorkProfileEasEmailProfileBaseUsernameSourceusername           AndroidWorkProfileEasEmailProfileBaseUsernameSource = "Username"
)

func PossibleValuesForAndroidWorkProfileEasEmailProfileBaseUsernameSource() []string {
	return []string{
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSourceprimarySmtpAddress),
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSourcesamAccountName),
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSourceuserPrincipalName),
		string(AndroidWorkProfileEasEmailProfileBaseUsernameSourceusername),
	}
}

func parseAndroidWorkProfileEasEmailProfileBaseUsernameSource(input string) (*AndroidWorkProfileEasEmailProfileBaseUsernameSource, error) {
	vals := map[string]AndroidWorkProfileEasEmailProfileBaseUsernameSource{
		"primarysmtpaddress": AndroidWorkProfileEasEmailProfileBaseUsernameSourceprimarySmtpAddress,
		"samaccountname":     AndroidWorkProfileEasEmailProfileBaseUsernameSourcesamAccountName,
		"userprincipalname":  AndroidWorkProfileEasEmailProfileBaseUsernameSourceuserPrincipalName,
		"username":           AndroidWorkProfileEasEmailProfileBaseUsernameSourceusername,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AndroidWorkProfileEasEmailProfileBaseUsernameSource(input)
	return &out, nil
}
