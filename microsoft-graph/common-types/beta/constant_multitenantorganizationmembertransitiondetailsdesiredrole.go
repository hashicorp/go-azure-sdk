package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationMemberTransitionDetailsDesiredRole string

const (
	MultiTenantOrganizationMemberTransitionDetailsDesiredRole_Member MultiTenantOrganizationMemberTransitionDetailsDesiredRole = "member"
	MultiTenantOrganizationMemberTransitionDetailsDesiredRole_Owner  MultiTenantOrganizationMemberTransitionDetailsDesiredRole = "owner"
)

func PossibleValuesForMultiTenantOrganizationMemberTransitionDetailsDesiredRole() []string {
	return []string{
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredRole_Member),
		string(MultiTenantOrganizationMemberTransitionDetailsDesiredRole_Owner),
	}
}

func (s *MultiTenantOrganizationMemberTransitionDetailsDesiredRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationMemberTransitionDetailsDesiredRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationMemberTransitionDetailsDesiredRole(input string) (*MultiTenantOrganizationMemberTransitionDetailsDesiredRole, error) {
	vals := map[string]MultiTenantOrganizationMemberTransitionDetailsDesiredRole{
		"member": MultiTenantOrganizationMemberTransitionDetailsDesiredRole_Member,
		"owner":  MultiTenantOrganizationMemberTransitionDetailsDesiredRole_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationMemberTransitionDetailsDesiredRole(input)
	return &out, nil
}
