package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PhoneAuthenticationMethodPhoneType string

const (
	PhoneAuthenticationMethodPhoneTypealternateMobile PhoneAuthenticationMethodPhoneType = "AlternateMobile"
	PhoneAuthenticationMethodPhoneTypemobile          PhoneAuthenticationMethodPhoneType = "Mobile"
	PhoneAuthenticationMethodPhoneTypeoffice          PhoneAuthenticationMethodPhoneType = "Office"
)

func PossibleValuesForPhoneAuthenticationMethodPhoneType() []string {
	return []string{
		string(PhoneAuthenticationMethodPhoneTypealternateMobile),
		string(PhoneAuthenticationMethodPhoneTypemobile),
		string(PhoneAuthenticationMethodPhoneTypeoffice),
	}
}

func parsePhoneAuthenticationMethodPhoneType(input string) (*PhoneAuthenticationMethodPhoneType, error) {
	vals := map[string]PhoneAuthenticationMethodPhoneType{
		"alternatemobile": PhoneAuthenticationMethodPhoneTypealternateMobile,
		"mobile":          PhoneAuthenticationMethodPhoneTypemobile,
		"office":          PhoneAuthenticationMethodPhoneTypeoffice,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PhoneAuthenticationMethodPhoneType(input)
	return &out, nil
}
