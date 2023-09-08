package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserAccountStatus string

const (
	UserAccountStatusactive    UserAccountStatus = "Active"
	UserAccountStatusdeleted   UserAccountStatus = "Deleted"
	UserAccountStatusstaged    UserAccountStatus = "Staged"
	UserAccountStatussuspended UserAccountStatus = "Suspended"
	UserAccountStatusunknown   UserAccountStatus = "Unknown"
)

func PossibleValuesForUserAccountStatus() []string {
	return []string{
		string(UserAccountStatusactive),
		string(UserAccountStatusdeleted),
		string(UserAccountStatusstaged),
		string(UserAccountStatussuspended),
		string(UserAccountStatusunknown),
	}
}

func parseUserAccountStatus(input string) (*UserAccountStatus, error) {
	vals := map[string]UserAccountStatus{
		"active":    UserAccountStatusactive,
		"deleted":   UserAccountStatusdeleted,
		"staged":    UserAccountStatusstaged,
		"suspended": UserAccountStatussuspended,
		"unknown":   UserAccountStatusunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserAccountStatus(input)
	return &out, nil
}
