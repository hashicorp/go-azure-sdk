package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummaryUserTypes string

const (
	UserRegistrationFeatureSummaryUserTypesall    UserRegistrationFeatureSummaryUserTypes = "All"
	UserRegistrationFeatureSummaryUserTypesguest  UserRegistrationFeatureSummaryUserTypes = "Guest"
	UserRegistrationFeatureSummaryUserTypesmember UserRegistrationFeatureSummaryUserTypes = "Member"
)

func PossibleValuesForUserRegistrationFeatureSummaryUserTypes() []string {
	return []string{
		string(UserRegistrationFeatureSummaryUserTypesall),
		string(UserRegistrationFeatureSummaryUserTypesguest),
		string(UserRegistrationFeatureSummaryUserTypesmember),
	}
}

func parseUserRegistrationFeatureSummaryUserTypes(input string) (*UserRegistrationFeatureSummaryUserTypes, error) {
	vals := map[string]UserRegistrationFeatureSummaryUserTypes{
		"all":    UserRegistrationFeatureSummaryUserTypesall,
		"guest":  UserRegistrationFeatureSummaryUserTypesguest,
		"member": UserRegistrationFeatureSummaryUserTypesmember,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationFeatureSummaryUserTypes(input)
	return &out, nil
}
