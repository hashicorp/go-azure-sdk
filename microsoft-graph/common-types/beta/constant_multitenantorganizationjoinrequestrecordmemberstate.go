package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestRecordMemberState string

const (
	MultiTenantOrganizationJoinRequestRecordMemberState_Active  MultiTenantOrganizationJoinRequestRecordMemberState = "active"
	MultiTenantOrganizationJoinRequestRecordMemberState_Pending MultiTenantOrganizationJoinRequestRecordMemberState = "pending"
	MultiTenantOrganizationJoinRequestRecordMemberState_Removed MultiTenantOrganizationJoinRequestRecordMemberState = "removed"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestRecordMemberState() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestRecordMemberState_Active),
		string(MultiTenantOrganizationJoinRequestRecordMemberState_Pending),
		string(MultiTenantOrganizationJoinRequestRecordMemberState_Removed),
	}
}

func (s *MultiTenantOrganizationJoinRequestRecordMemberState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationJoinRequestRecordMemberState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationJoinRequestRecordMemberState(input string) (*MultiTenantOrganizationJoinRequestRecordMemberState, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestRecordMemberState{
		"active":  MultiTenantOrganizationJoinRequestRecordMemberState_Active,
		"pending": MultiTenantOrganizationJoinRequestRecordMemberState_Pending,
		"removed": MultiTenantOrganizationJoinRequestRecordMemberState_Removed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestRecordMemberState(input)
	return &out, nil
}
