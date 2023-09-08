package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberRole string

const (
	MultiTenantOrganizationMemberRolemember MultiTenantOrganizationMemberRole = "Member"
	MultiTenantOrganizationMemberRoleowner  MultiTenantOrganizationMemberRole = "Owner"
)

func PossibleValuesForMultiTenantOrganizationMemberRole() []string {
	return []string{
		string(MultiTenantOrganizationMemberRolemember),
		string(MultiTenantOrganizationMemberRoleowner),
	}
}

func parseMultiTenantOrganizationMemberRole(input string) (*MultiTenantOrganizationMemberRole, error) {
	vals := map[string]MultiTenantOrganizationMemberRole{
		"member": MultiTenantOrganizationMemberRolemember,
		"owner":  MultiTenantOrganizationMemberRoleowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberRole(input)
	return &out, nil
}
