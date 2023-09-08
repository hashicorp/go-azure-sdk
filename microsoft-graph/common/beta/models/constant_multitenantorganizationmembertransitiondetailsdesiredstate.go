package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetailsDesiredState string

const (
	MultiTenantOrganizationMemberTransitionDetailsDesiredStateactive  MultiTenantOrganizationMemberTransitionDetailsDesiredState = "Active"
	MultiTenantOrganizationMemberTransitionDetailsDesiredStatepending MultiTenantOrganizationMemberTransitionDetailsDesiredState = "Pending"
	MultiTenantOrganizationMemberTransitionDetailsDesiredStateremoved MultiTenantOrganizationMemberTransitionDetailsDesiredState = "Removed"
)

func PossibleValuesForMultiTenantOrganizationMemberTransitionDetailsDesiredState() []string {
	return []string{
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredStateactive),
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredStatepending),
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredStateremoved),
	}
}

func parseMultiTenantOrganizationMemberTransitionDetailsDesiredState(input string) (*MultiTenantOrganizationMemberTransitionDetailsDesiredState, error) {
	vals := map[string]MultiTenantOrganizationMemberTransitionDetailsDesiredState{
		"active":  MultiTenantOrganizationMemberTransitionDetailsDesiredStateactive,
		"pending": MultiTenantOrganizationMemberTransitionDetailsDesiredStatepending,
		"removed": MultiTenantOrganizationMemberTransitionDetailsDesiredStateremoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberTransitionDetailsDesiredState(input)
	return &out, nil
}
