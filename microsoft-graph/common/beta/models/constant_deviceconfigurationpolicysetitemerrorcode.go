package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationPolicySetItemErrorCode string

const (
	DeviceConfigurationPolicySetItemErrorCodedeleted      DeviceConfigurationPolicySetItemErrorCode = "Deleted"
	DeviceConfigurationPolicySetItemErrorCodenoError      DeviceConfigurationPolicySetItemErrorCode = "NoError"
	DeviceConfigurationPolicySetItemErrorCodenotFound     DeviceConfigurationPolicySetItemErrorCode = "NotFound"
	DeviceConfigurationPolicySetItemErrorCodeunauthorized DeviceConfigurationPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForDeviceConfigurationPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceConfigurationPolicySetItemErrorCodedeleted),
		string(DeviceConfigurationPolicySetItemErrorCodenoError),
		string(DeviceConfigurationPolicySetItemErrorCodenotFound),
		string(DeviceConfigurationPolicySetItemErrorCodeunauthorized),
	}
}

func parseDeviceConfigurationPolicySetItemErrorCode(input string) (*DeviceConfigurationPolicySetItemErrorCode, error) {
	vals := map[string]DeviceConfigurationPolicySetItemErrorCode{
		"deleted":      DeviceConfigurationPolicySetItemErrorCodedeleted,
		"noerror":      DeviceConfigurationPolicySetItemErrorCodenoError,
		"notfound":     DeviceConfigurationPolicySetItemErrorCodenotFound,
		"unauthorized": DeviceConfigurationPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationPolicySetItemErrorCode(input)
	return &out, nil
}
