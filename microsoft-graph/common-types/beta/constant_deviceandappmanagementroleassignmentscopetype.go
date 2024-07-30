package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceAndAppManagementRoleAssignmentScopeType string

const (
	DeviceAndAppManagementRoleAssignmentScopeType_AllDevices                 DeviceAndAppManagementRoleAssignmentScopeType = "allDevices"
	DeviceAndAppManagementRoleAssignmentScopeType_AllDevicesAndLicensedUsers DeviceAndAppManagementRoleAssignmentScopeType = "allDevicesAndLicensedUsers"
	DeviceAndAppManagementRoleAssignmentScopeType_AllLicensedUsers           DeviceAndAppManagementRoleAssignmentScopeType = "allLicensedUsers"
	DeviceAndAppManagementRoleAssignmentScopeType_ResourceScope              DeviceAndAppManagementRoleAssignmentScopeType = "resourceScope"
)

func PossibleValuesForDeviceAndAppManagementRoleAssignmentScopeType() []string {
	return []string{
		string(DeviceAndAppManagementRoleAssignmentScopeType_AllDevices),
		string(DeviceAndAppManagementRoleAssignmentScopeType_AllDevicesAndLicensedUsers),
		string(DeviceAndAppManagementRoleAssignmentScopeType_AllLicensedUsers),
		string(DeviceAndAppManagementRoleAssignmentScopeType_ResourceScope),
	}
}

func (s *DeviceAndAppManagementRoleAssignmentScopeType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceAndAppManagementRoleAssignmentScopeType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceAndAppManagementRoleAssignmentScopeType(input string) (*DeviceAndAppManagementRoleAssignmentScopeType, error) {
	vals := map[string]DeviceAndAppManagementRoleAssignmentScopeType{
		"alldevices":                 DeviceAndAppManagementRoleAssignmentScopeType_AllDevices,
		"alldevicesandlicensedusers": DeviceAndAppManagementRoleAssignmentScopeType_AllDevicesAndLicensedUsers,
		"alllicensedusers":           DeviceAndAppManagementRoleAssignmentScopeType_AllLicensedUsers,
		"resourcescope":              DeviceAndAppManagementRoleAssignmentScopeType_ResourceScope,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceAndAppManagementRoleAssignmentScopeType(input)
	return &out, nil
}
