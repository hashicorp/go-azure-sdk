package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateUserAccountType string

const (
	UserSecurityStateUserAccountTypeadministrator UserSecurityStateUserAccountType = "Administrator"
	UserSecurityStateUserAccountTypepower         UserSecurityStateUserAccountType = "Power"
	UserSecurityStateUserAccountTypestandard      UserSecurityStateUserAccountType = "Standard"
	UserSecurityStateUserAccountTypeunknown       UserSecurityStateUserAccountType = "Unknown"
)

func PossibleValuesForUserSecurityStateUserAccountType() []string {
	return []string{
		string(UserSecurityStateUserAccountTypeadministrator),
		string(UserSecurityStateUserAccountTypepower),
		string(UserSecurityStateUserAccountTypestandard),
		string(UserSecurityStateUserAccountTypeunknown),
	}
}

func parseUserSecurityStateUserAccountType(input string) (*UserSecurityStateUserAccountType, error) {
	vals := map[string]UserSecurityStateUserAccountType{
		"administrator": UserSecurityStateUserAccountTypeadministrator,
		"power":         UserSecurityStateUserAccountTypepower,
		"standard":      UserSecurityStateUserAccountTypestandard,
		"unknown":       UserSecurityStateUserAccountTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSecurityStateUserAccountType(input)
	return &out, nil
}
