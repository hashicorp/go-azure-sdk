package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MultiTenantOrganizationJoinRequestRecordRole string

const (
	MultiTenantOrganizationJoinRequestRecordRole_Member MultiTenantOrganizationJoinRequestRecordRole = "member"
	MultiTenantOrganizationJoinRequestRecordRole_Owner  MultiTenantOrganizationJoinRequestRecordRole = "owner"
)

func PossibleValuesForMultiTenantOrganizationJoinRequestRecordRole() []string {
	return []string{
		string(MultiTenantOrganizationJoinRequestRecordRole_Member),
		string(MultiTenantOrganizationJoinRequestRecordRole_Owner),
	}
}

func (s *MultiTenantOrganizationJoinRequestRecordRole) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMultiTenantOrganizationJoinRequestRecordRole(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMultiTenantOrganizationJoinRequestRecordRole(input string) (*MultiTenantOrganizationJoinRequestRecordRole, error) {
	vals := map[string]MultiTenantOrganizationJoinRequestRecordRole{
		"member": MultiTenantOrganizationJoinRequestRecordRole_Member,
		"owner":  MultiTenantOrganizationJoinRequestRecordRole_Owner,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MultiTenantOrganizationJoinRequestRecordRole(input)
	return &out, nil
}
