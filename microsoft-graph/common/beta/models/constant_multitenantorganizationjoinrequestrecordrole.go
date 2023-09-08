package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestRecordRole string

const (
	MultiTenantOrganizationJoinRequestRecordRolemember MultiTenantOrganizationJoinRequestRecordRole = "Member"
	MultiTenantOrganizationJoinRequestRecordRoleowner  MultiTenantOrganizationJoinRequestRecordRole = "Owner"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestRecordRole() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestRecordRolemember),
		string(MultiTenantOrganizationJoinRequestRecordRoleowner),
	}
}

func parseMultiTenantOrganizationJoinRequestRecordRole(input string) (*MultiTenantOrganizationJoinRequestRecordRole, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestRecordRole{
		"member": MultiTenantOrganizationJoinRequestRecordRolemember,
		"owner":  MultiTenantOrganizationJoinRequestRecordRoleowner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestRecordRole(input)
	return &out, nil
}
