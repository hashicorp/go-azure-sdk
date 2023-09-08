package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSecurityStateEmailRole string

const (
	UserSecurityStateEmailRolerecipient UserSecurityStateEmailRole = "Recipient"
	UserSecurityStateEmailRolesender    UserSecurityStateEmailRole = "Sender"
	UserSecurityStateEmailRoleunknown   UserSecurityStateEmailRole = "Unknown"
)

func PossibleValuesForUserSecurityStateEmailRole() []string {
	return []string{
		string(UserSecurityStateEmailRolerecipient),
		string(UserSecurityStateEmailRolesender),
		string(UserSecurityStateEmailRoleunknown),
	}
}

func parseUserSecurityStateEmailRole(input string) (*UserSecurityStateEmailRole, error) {
	vals := map[string]UserSecurityStateEmailRole{
		"recipient": UserSecurityStateEmailRolerecipient,
		"sender":    UserSecurityStateEmailRolesender,
		"unknown":   UserSecurityStateEmailRoleunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserSecurityStateEmailRole(input)
	return &out, nil
}
