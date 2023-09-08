package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyPolicySetItemErrorCode string

const (
	DeviceManagementConfigurationPolicyPolicySetItemErrorCodedeleted      DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "Deleted"
	DeviceManagementConfigurationPolicyPolicySetItemErrorCodenoError      DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "NoError"
	DeviceManagementConfigurationPolicyPolicySetItemErrorCodenotFound     DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "NotFound"
	DeviceManagementConfigurationPolicyPolicySetItemErrorCodeunauthorized DeviceManagementConfigurationPolicyPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForDeviceManagementConfigurationPolicyPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCodedeleted),
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCodenoError),
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCodenotFound),
		string(DeviceManagementConfigurationPolicyPolicySetItemErrorCodeunauthorized),
	}
}

func parseDeviceManagementConfigurationPolicyPolicySetItemErrorCode(input string) (*DeviceManagementConfigurationPolicyPolicySetItemErrorCode, error) {
	vals := map[string]DeviceManagementConfigurationPolicyPolicySetItemErrorCode{
		"deleted":      DeviceManagementConfigurationPolicyPolicySetItemErrorCodedeleted,
		"noerror":      DeviceManagementConfigurationPolicyPolicySetItemErrorCodenoError,
		"notfound":     DeviceManagementConfigurationPolicyPolicySetItemErrorCodenotFound,
		"unauthorized": DeviceManagementConfigurationPolicyPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyPolicySetItemErrorCode(input)
	return &out, nil
}
