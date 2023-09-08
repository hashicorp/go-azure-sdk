package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus string

const (
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusdelegatedAdminPrivileges                     ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "DelegatedAdminPrivileges"
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusdelegatedAndGranularDelegetedAdminPrivileges ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "DelegatedAndGranularDelegetedAdminPrivileges"
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusgranularDelegatedAdminPrivileges             ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "GranularDelegatedAdminPrivileges"
	ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusnone                                         ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus = "None"
)

func PossibleValuesForManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus() []string {
	return []string{
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusdelegatedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusdelegatedAndGranularDelegetedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusgranularDelegatedAdminPrivileges),
		string(ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusnone),
	}
}

func parseManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus(input string) (*ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus, error) {
	vals := map[string]ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus{
		"delegatedadminprivileges":                     ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusdelegatedAdminPrivileges,
		"delegatedandgranulardelegetedadminprivileges": ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusdelegatedAndGranularDelegetedAdminPrivileges,
		"granulardelegatedadminprivileges":             ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusgranularDelegatedAdminPrivileges,
		"none":                                         ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatusnone,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedTenantsTenantStatusInformationDelegatedPrivilegeStatus(input)
	return &out, nil
}
