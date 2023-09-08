package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementRoleAssignmentScopeType string

const (
	DeviceAndAppManagementRoleAssignmentScopeTypeallDevices                 DeviceAndAppManagementRoleAssignmentScopeType = "AllDevices"
	DeviceAndAppManagementRoleAssignmentScopeTypeallDevicesAndLicensedUsers DeviceAndAppManagementRoleAssignmentScopeType = "AllDevicesAndLicensedUsers"
	DeviceAndAppManagementRoleAssignmentScopeTypeallLicensedUsers           DeviceAndAppManagementRoleAssignmentScopeType = "AllLicensedUsers"
	DeviceAndAppManagementRoleAssignmentScopeTyperesourceScope              DeviceAndAppManagementRoleAssignmentScopeType = "ResourceScope"
)

func PossibleValuesForDeviceAndAppManagementRoleAssignmentScopeType() []string {
	return []string{
		string(DeviceAndAppManagementRoleAssignmentScopeTypeallDevices),
		string(DeviceAndAppManagementRoleAssignmentScopeTypeallDevicesAndLicensedUsers),
		string(DeviceAndAppManagementRoleAssignmentScopeTypeallLicensedUsers),
		string(DeviceAndAppManagementRoleAssignmentScopeTyperesourceScope),
	}
}

func parseDeviceAndAppManagementRoleAssignmentScopeType(input string) (*DeviceAndAppManagementRoleAssignmentScopeType, error) {
	vals := map[string]DeviceAndAppManagementRoleAssignmentScopeType{
		"alldevices":                 DeviceAndAppManagementRoleAssignmentScopeTypeallDevices,
		"alldevicesandlicensedusers": DeviceAndAppManagementRoleAssignmentScopeTypeallDevicesAndLicensedUsers,
		"alllicensedusers":           DeviceAndAppManagementRoleAssignmentScopeTypeallLicensedUsers,
		"resourcescope":              DeviceAndAppManagementRoleAssignmentScopeTyperesourceScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementRoleAssignmentScopeType(input)
	return &out, nil
}
