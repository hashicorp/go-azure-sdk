package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState string

const (
	MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Active  MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState = "active"
	MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Pending MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState = "pending"
	MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Removed MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState = "removed"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Active),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Pending),
		string(MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Removed),
	}
}

func (s *MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState(input string) (*MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState{
		"active":  MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Active,
		"pending": MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Pending,
		"removed": MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState_Removed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestTransitionDetailsDesiredMemberState(input)
	return &out, nil
}
