package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetailsDefaultMfaMethod string

const (
	UserRegistrationDetailsDefaultMfaMethodalternateMobilePhone       UserRegistrationDetailsDefaultMfaMethod = "AlternateMobilePhone"
	UserRegistrationDetailsDefaultMfaMethodmicrosoftAuthenticatorPush UserRegistrationDetailsDefaultMfaMethod = "MicrosoftAuthenticatorPush"
	UserRegistrationDetailsDefaultMfaMethodmobilePhone                UserRegistrationDetailsDefaultMfaMethod = "MobilePhone"
	UserRegistrationDetailsDefaultMfaMethodnone                       UserRegistrationDetailsDefaultMfaMethod = "None"
	UserRegistrationDetailsDefaultMfaMethodofficePhone                UserRegistrationDetailsDefaultMfaMethod = "OfficePhone"
	UserRegistrationDetailsDefaultMfaMethodsoftwareOneTimePasscode    UserRegistrationDetailsDefaultMfaMethod = "SoftwareOneTimePasscode"
)

func PossibleValuesForUserRegistrationDetailsDefaultMfaMethod() []string {
	return []string{
		string(UserRegistrationDetailsDefaultMfaMethodalternateMobilePhone),
		string(UserRegistrationDetailsDefaultMfaMethodmicrosoftAuthenticatorPush),
		string(UserRegistrationDetailsDefaultMfaMethodmobilePhone),
		string(UserRegistrationDetailsDefaultMfaMethodnone),
		string(UserRegistrationDetailsDefaultMfaMethodofficePhone),
		string(UserRegistrationDetailsDefaultMfaMethodsoftwareOneTimePasscode),
	}
}

func parseUserRegistrationDetailsDefaultMfaMethod(input string) (*UserRegistrationDetailsDefaultMfaMethod, error) {
	vals := map[string]UserRegistrationDetailsDefaultMfaMethod{
		"alternatemobilephone":       UserRegistrationDetailsDefaultMfaMethodalternateMobilePhone,
		"microsoftauthenticatorpush": UserRegistrationDetailsDefaultMfaMethodmicrosoftAuthenticatorPush,
		"mobilephone":                UserRegistrationDetailsDefaultMfaMethodmobilePhone,
		"none":                       UserRegistrationDetailsDefaultMfaMethodnone,
		"officephone":                UserRegistrationDetailsDefaultMfaMethodofficePhone,
		"softwareonetimepasscode":    UserRegistrationDetailsDefaultMfaMethodsoftwareOneTimePasscode,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationDetailsDefaultMfaMethod(input)
	return &out, nil
}
