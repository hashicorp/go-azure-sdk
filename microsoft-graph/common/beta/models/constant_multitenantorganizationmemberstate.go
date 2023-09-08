package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberState string

const (
	MultiTenantOrganizationMemberStateactive  MultiTenantOrganizationMemberState = "Active"
	MultiTenantOrganizationMemberStatepending MultiTenantOrganizationMemberState = "Pending"
	MultiTenantOrganizationMemberStateremoved MultiTenantOrganizationMemberState = "Removed"
)

func PossibleValuesForMultiTenantOrganizationMemberState() []string {
	return []string{
		string(MultiTenantOrganizationMemberStateactive),
		string(MultiTenantOrganizationMemberStatepending),
		string(MultiTenantOrganizationMemberStateremoved),
	}
}

func parseMultiTenantOrganizationMemberState(input string) (*MultiTenantOrganizationMemberState, error) {
	vals := map[string]MultiTenantOrganizationMemberState{
		"active":  MultiTenantOrganizationMemberStateactive,
		"pending": MultiTenantOrganizationMemberStatepending,
		"removed": MultiTenantOrganizationMemberStateremoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberState(input)
	return &out, nil
}
