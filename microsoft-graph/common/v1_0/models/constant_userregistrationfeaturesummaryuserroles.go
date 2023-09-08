package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserRegistrationFeatureSummaryUserRoles string

const (
	UserRegistrationFeatureSummaryUserRolesadmin           UserRegistrationFeatureSummaryUserRoles = "Admin"
	UserRegistrationFeatureSummaryUserRolesall             UserRegistrationFeatureSummaryUserRoles = "All"
	UserRegistrationFeatureSummaryUserRolesprivilegedAdmin UserRegistrationFeatureSummaryUserRoles = "PrivilegedAdmin"
	UserRegistrationFeatureSummaryUserRolesuser            UserRegistrationFeatureSummaryUserRoles = "User"
)

func PossibleValuesForUserRegistrationFeatureSummaryUserRoles() []string {
	return []string{
		string(UserRegistrationFeatureSummaryUserRolesadmin),
		string(UserRegistrationFeatureSummaryUserRolesall),
		string(UserRegistrationFeatureSummaryUserRolesprivilegedAdmin),
		string(UserRegistrationFeatureSummaryUserRolesuser),
	}
}

func parseUserRegistrationFeatureSummaryUserRoles(input string) (*UserRegistrationFeatureSummaryUserRoles, error) {
	vals := map[string]UserRegistrationFeatureSummaryUserRoles{
		"admin":           UserRegistrationFeatureSummaryUserRolesadmin,
		"all":             UserRegistrationFeatureSummaryUserRolesall,
		"privilegedadmin": UserRegistrationFeatureSummaryUserRolesprivilegedAdmin,
		"user":            UserRegistrationFeatureSummaryUserRolesuser,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserRegistrationFeatureSummaryUserRoles(input)
	return &out, nil
}
