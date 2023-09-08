package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestRecordMemberState string

const (
	MultiTenantOrganizationJoinRequestRecordMemberStateactive  MultiTenantOrganizationJoinRequestRecordMemberState = "Active"
	MultiTenantOrganizationJoinRequestRecordMemberStatepending MultiTenantOrganizationJoinRequestRecordMemberState = "Pending"
	MultiTenantOrganizationJoinRequestRecordMemberStateremoved MultiTenantOrganizationJoinRequestRecordMemberState = "Removed"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestRecordMemberState() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestRecordMemberStateactive),
		string(MultiTenantOrganizationJoinRequestRecordMemberStatepending),
		string(MultiTenantOrganizationJoinRequestRecordMemberStateremoved),
	}
}

func parseMultiTenantOrganizationJoinRequestRecordMemberState(input string) (*MultiTenantOrganizationJoinRequestRecordMemberState, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestRecordMemberState{
		"active":  MultiTenantOrganizationJoinRequestRecordMemberStateactive,
		"pending": MultiTenantOrganizationJoinRequestRecordMemberStatepending,
		"removed": MultiTenantOrganizationJoinRequestRecordMemberStateremoved,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestRecordMemberState(input)
	return &out, nil
}
