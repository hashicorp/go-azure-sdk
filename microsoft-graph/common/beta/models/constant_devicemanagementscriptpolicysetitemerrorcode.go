package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptPolicySetItemErrorCode string

const (
	DeviceManagementScriptPolicySetItemErrorCodedeleted      DeviceManagementScriptPolicySetItemErrorCode = "Deleted"
	DeviceManagementScriptPolicySetItemErrorCodenoError      DeviceManagementScriptPolicySetItemErrorCode = "NoError"
	DeviceManagementScriptPolicySetItemErrorCodenotFound     DeviceManagementScriptPolicySetItemErrorCode = "NotFound"
	DeviceManagementScriptPolicySetItemErrorCodeunauthorized DeviceManagementScriptPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForDeviceManagementScriptPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceManagementScriptPolicySetItemErrorCodedeleted),
		string(DeviceManagementScriptPolicySetItemErrorCodenoError),
		string(DeviceManagementScriptPolicySetItemErrorCodenotFound),
		string(DeviceManagementScriptPolicySetItemErrorCodeunauthorized),
	}
}

func parseDeviceManagementScriptPolicySetItemErrorCode(input string) (*DeviceManagementScriptPolicySetItemErrorCode, error) {
	vals := map[string]DeviceManagementScriptPolicySetItemErrorCode{
		"deleted":      DeviceManagementScriptPolicySetItemErrorCodedeleted,
		"noerror":      DeviceManagementScriptPolicySetItemErrorCodenoError,
		"notfound":     DeviceManagementScriptPolicySetItemErrorCodenotFound,
		"unauthorized": DeviceManagementScriptPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptPolicySetItemErrorCode(input)
	return &out, nil
}
