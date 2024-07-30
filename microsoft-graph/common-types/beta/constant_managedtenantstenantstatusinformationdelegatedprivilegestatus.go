package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus string

const (
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_DelegatedAdminPrivileges                     ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "delegatedAdminPrivileges"
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_DelegatedAndGranularDelegetedAdminPrivileges ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "delegatedAndGranularDelegetedAdminPrivileges"
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_GranularDelegatedAdminPrivileges             ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "granularDelegatedAdminPrivileges"
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_None                                         ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "none"
)

func PossibleValuesForManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus() []string {
	return []string{
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_DelegatedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_DelegatedAndGranularDelegetedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_GranularDelegatedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_None),
	}
}

func (s *ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus(input string) (*ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus, error) {
	vals := map[string]ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus{
		"delegatedadminprivileges":                     ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_DelegatedAdminPrivileges,
		"delegatedandgranulardelegetedadminprivileges": ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_DelegatedAndGranularDelegetedAdminPrivileges,
		"granulardelegatedadminprivileges":             ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_GranularDelegatedAdminPrivileges,
		"none":                                         ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus_None,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus(input)
	return &out, nil
}
