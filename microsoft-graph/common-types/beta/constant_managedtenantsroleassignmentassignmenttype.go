package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsRoleAssignmentAssignmentType string

const (
	ManagedTenantsRoleAssignmentAssignmentType_DelegatedAdminPrivileges                     ManagedTenantsRoleAssignmentAssignmentType = "delegatedAdminPrivileges"
	ManagedTenantsRoleAssignmentAssignmentType_DelegatedAndGranularDelegetedAdminPrivileges ManagedTenantsRoleAssignmentAssignmentType = "delegatedAndGranularDelegetedAdminPrivileges"
	ManagedTenantsRoleAssignmentAssignmentType_GranularDelegatedAdminPrivileges             ManagedTenantsRoleAssignmentAssignmentType = "granularDelegatedAdminPrivileges"
	ManagedTenantsRoleAssignmentAssignmentType_None                                         ManagedTenantsRoleAssignmentAssignmentType = "none"
)

func PossibleValuesForManagedTenantsRoleAssignmentAssignmentType() []string {
	return []string{
		string(ManagedTenantsRoleAssignmentAssignmentType_DelegatedAdminPrivileges),
		string(ManagedTenantsRoleAssignmentAssignmentType_DelegatedAndGranularDelegetedAdminPrivileges),
		string(ManagedTenantsRoleAssignmentAssignmentType_GranularDelegatedAdminPrivileges),
		string(ManagedTenantsRoleAssignmentAssignmentType_None),
	}
}

func (s *ManagedTenantsRoleAssignmentAssignmentType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsRoleAssignmentAssignmentType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsRoleAssignmentAssignmentType(input string) (*ManagedTenantsRoleAssignmentAssignmentType, error) {
	vals := map[string]ManagedTenantsRoleAssignmentAssignmentType{
		"delegatedadminprivileges":                     ManagedTenantsRoleAssignmentAssignmentType_DelegatedAdminPrivileges,
		"delegatedandgranulardelegetedadminprivileges": ManagedTenantsRoleAssignmentAssignmentType_DelegatedAndGranularDelegetedAdminPrivileges,
		"granulardelegatedadminprivileges":             ManagedTenantsRoleAssignmentAssignmentType_GranularDelegatedAdminPrivileges,
		"none":                                         ManagedTenantsRoleAssignmentAssignmentType_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsRoleAssignmentAssignmentType(input)
	return &out, nil
}
