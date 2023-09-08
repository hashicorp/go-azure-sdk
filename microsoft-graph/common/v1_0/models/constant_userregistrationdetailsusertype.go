package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationDetailsUserType string

const (
	UserRegistrationDetailsUserTypeguest  UserRegistrationDetailsUserType = "Guest"
	UserRegistrationDetailsUserTypemember UserRegistrationDetailsUserType = "Member"
)

func PossibleValuesForUserRegistrationDetailsUserType() []string {
	return []string{
		string(UserRegistrationDetailsUserTypeguest),
		string(UserRegistrationDetailsUserTypemember),
	}
}

func parseUserRegistrationDetailsUserType(input string) (*UserRegistrationDetailsUserType, error) {
	vals := map[string]UserRegistrationDetailsUserType{
		"guest":  UserRegistrationDetailsUserTypeguest,
		"member": UserRegistrationDetailsUserTypemember,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationDetailsUserType(input)
	return &out, nil
}
