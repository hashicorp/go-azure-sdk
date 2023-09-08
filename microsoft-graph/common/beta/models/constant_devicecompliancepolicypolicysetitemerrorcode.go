package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyPolicySetItemErrorCode string

const (
	DeviceCompliancePolicyPolicySetItemErrorCodedeleted      DeviceCompliancePolicyPolicySetItemErrorCode = "Deleted"
	DeviceCompliancePolicyPolicySetItemErrorCodenoError      DeviceCompliancePolicyPolicySetItemErrorCode = "NoError"
	DeviceCompliancePolicyPolicySetItemErrorCodenotFound     DeviceCompliancePolicyPolicySetItemErrorCode = "NotFound"
	DeviceCompliancePolicyPolicySetItemErrorCodeunauthorized DeviceCompliancePolicyPolicySetItemErrorCode = "Unauthorized"
)

func PossibleValuesForDeviceCompliancePolicyPolicySetItemErrorCode() []string {
	return []string{
		string(DeviceCompliancePolicyPolicySetItemErrorCodedeleted),
		string(DeviceCompliancePolicyPolicySetItemErrorCodenoError),
		string(DeviceCompliancePolicyPolicySetItemErrorCodenotFound),
		string(DeviceCompliancePolicyPolicySetItemErrorCodeunauthorized),
	}
}

func parseDeviceCompliancePolicyPolicySetItemErrorCode(input string) (*DeviceCompliancePolicyPolicySetItemErrorCode, error) {
	vals := map[string]DeviceCompliancePolicyPolicySetItemErrorCode{
		"deleted":      DeviceCompliancePolicyPolicySetItemErrorCodedeleted,
		"noerror":      DeviceCompliancePolicyPolicySetItemErrorCodenoError,
		"notfound":     DeviceCompliancePolicyPolicySetItemErrorCodenotFound,
		"unauthorized": DeviceCompliancePolicyPolicySetItemErrorCodeunauthorized,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyPolicySetItemErrorCode(input)
	return &out, nil
}
