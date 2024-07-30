package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetailsDesiredState string

const (
	MultiTenantOrganizationMemberTransitionDetailsDesiredState_Active  MultiTenantOrganizationMemberTransitionDetailsDesiredState = "active"
	MultiTenantOrganizationMemberTransitionDetailsDesiredState_Pending MultiTenantOrganizationMemberTransitionDetailsDesiredState = "pending"
	MultiTenantOrganizationMemberTransitionDetailsDesiredState_Removed MultiTenantOrganizationMemberTransitionDetailsDesiredState = "removed"
)

func PossibleValuesForMultiTenantOrganizationMemberTransitionDetailsDesiredState() []string {
	return []string{
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredState_Active),
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredState_Pending),
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredState_Removed),
	}
}

func (s *MultiTenantOrganizationMemberTransitionDetailsDesiredState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationMemberTransitionDetailsDesiredState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationMemberTransitionDetailsDesiredState(input string) (*MultiTenantOrganizationMemberTransitionDetailsDesiredState, error) {
	vals := map[string]MultiTenantOrganizationMemberTransitionDetailsDesiredState{
		"active":  MultiTenantOrganizationMemberTransitionDetailsDesiredState_Active,
		"pending": MultiTenantOrganizationMemberTransitionDetailsDesiredState_Pending,
		"removed": MultiTenantOrganizationMemberTransitionDetailsDesiredState_Removed,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberTransitionDetailsDesiredState(input)
	return &out, nil
}
