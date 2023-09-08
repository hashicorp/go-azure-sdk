package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationMethodSummaryUserTypes string

const (
	UserRegistrationMethodSummaryUserTypesall    UserRegistrationMethodSummaryUserTypes = "All"
	UserRegistrationMethodSummaryUserTypesguest  UserRegistrationMethodSummaryUserTypes = "Guest"
	UserRegistrationMethodSummaryUserTypesmember UserRegistrationMethodSummaryUserTypes = "Member"
)

func PossibleValuesForUserRegistrationMethodSummaryUserTypes() []string {
	return []string{
		string(UserRegistrationMethodSummaryUserTypesall),
		string(UserRegistrationMethodSummaryUserTypesguest),
		string(UserRegistrationMethodSummaryUserTypesmember),
	}
}

func parseUserRegistrationMethodSummaryUserTypes(input string) (*UserRegistrationMethodSummaryUserTypes, error) {
	vals := map[string]UserRegistrationMethodSummaryUserTypes{
		"all":    UserRegistrationMethodSummaryUserTypesall,
		"guest":  UserRegistrationMethodSummaryUserTypesguest,
		"member": UserRegistrationMethodSummaryUserTypesmember,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationMethodSummaryUserTypes(input)
	return &out, nil
}
