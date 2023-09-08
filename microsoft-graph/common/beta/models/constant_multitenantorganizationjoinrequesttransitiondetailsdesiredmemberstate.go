package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState string

const (
	MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStateactive  MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState = "Active"
	MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStatepending MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState = "Pending"
	MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStateremoved MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState = "Removed"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStateactive),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStatepending),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStateremoved),
	}
}

func parseMultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState(input string) (*MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState{
		"active":  MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStateactive,
		"pending": MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStatepending,
		"removed": MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberStateremoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState(input)
	return &out, nil
}
