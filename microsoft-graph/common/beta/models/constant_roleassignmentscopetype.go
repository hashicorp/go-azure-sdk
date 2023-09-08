package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoleAssignmentScopeType string

const (
	RoleAssignmentScopeTypeallDevices                 RoleAssignmentScopeType = "AllDevices"
	RoleAssignmentScopeTypeallDevicesAndLicensedUsers RoleAssignmentScopeType = "AllDevicesAndLicensedUsers"
	RoleAssignmentScopeTypeallLicensedUsers           RoleAssignmentScopeType = "AllLicensedUsers"
	RoleAssignmentScopeTyperesourceScope              RoleAssignmentScopeType = "ResourceScope"
)

func PossibleValuesForRoleAssignmentScopeType() []string {
	return []string{
		string(RoleAssignmentScopeTypeallDevices),
		string(RoleAssignmentScopeTypeallDevicesAndLicensedUsers),
		string(RoleAssignmentScopeTypeallLicensedUsers),
		string(RoleAssignmentScopeTyperesourceScope),
	}
}

func parseRoleAssignmentScopeType(input string) (*RoleAssignmentScopeType, error) {
	vals := map[string]RoleAssignmentScopeType{
		"alldevices":                 RoleAssignmentScopeTypeallDevices,
		"alldevicesandlicensedusers": RoleAssignmentScopeTypeallDevicesAndLicensedUsers,
		"alllicensedusers":           RoleAssignmentScopeTypeallLicensedUsers,
		"resourcescope":              RoleAssignmentScopeTyperesourceScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RoleAssignmentScopeType(input)
	return &out, nil
}
