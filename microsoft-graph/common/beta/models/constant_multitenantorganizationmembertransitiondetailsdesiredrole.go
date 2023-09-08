package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetailsDesiredRole string

const (
	MultiTenantOrganizationMemberTransitionDetailsDesiredRolemember MultiTenantOrganizationMemberTransitionDetailsDesiredRole = "Member"
	MultiTenantOrganizationMemberTransitionDetailsDesiredRoleowner  MultiTenantOrganizationMemberTransitionDetailsDesiredRole = "Owner"
)

func PossibleValuesForMultiTenantOrganizationMemberTransitionDetailsDesiredRole() []string {
	return []string{
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredRolemember),
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredRoleowner),
	}
}

func parseMultiTenantOrganizationMemberTransitionDetailsDesiredRole(input string) (*MultiTenantOrganizationMemberTransitionDetailsDesiredRole, error) {
	vals := map[string]MultiTenantOrganizationMemberTransitionDetailsDesiredRole{
		"member": MultiTenantOrganizationMemberTransitionDetailsDesiredRolemember,
		"owner":  MultiTenantOrganizationMemberTransitionDetailsDesiredRoleowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberTransitionDetailsDesiredRole(input)
	return &out, nil
}
