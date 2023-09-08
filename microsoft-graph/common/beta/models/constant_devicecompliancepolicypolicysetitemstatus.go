package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyPolicySetItemStatus string

const (
	DeviceCompliancePolicyPolicySetItemStatuserror          DeviceCompliancePolicyPolicySetItemStatus = "Error"
	DeviceCompliancePolicyPolicySetItemStatusnotAssigned    DeviceCompliancePolicyPolicySetItemStatus = "NotAssigned"
	DeviceCompliancePolicyPolicySetItemStatuspartialSuccess DeviceCompliancePolicyPolicySetItemStatus = "PartialSuccess"
	DeviceCompliancePolicyPolicySetItemStatussuccess        DeviceCompliancePolicyPolicySetItemStatus = "Success"
	DeviceCompliancePolicyPolicySetItemStatusunknown        DeviceCompliancePolicyPolicySetItemStatus = "Unknown"
	DeviceCompliancePolicyPolicySetItemStatusvalidating     DeviceCompliancePolicyPolicySetItemStatus = "Validating"
)

func PossibleValuesForDeviceCompliancePolicyPolicySetItemStatus() []string {
	return []string{
		string(DeviceCompliancePolicyPolicySetItemStatuserror),
		string(DeviceCompliancePolicyPolicySetItemStatusnotAssigned),
		string(DeviceCompliancePolicyPolicySetItemStatuspartialSuccess),
		string(DeviceCompliancePolicyPolicySetItemStatussuccess),
		string(DeviceCompliancePolicyPolicySetItemStatusunknown),
		string(DeviceCompliancePolicyPolicySetItemStatusvalidating),
	}
}

func parseDeviceCompliancePolicyPolicySetItemStatus(input string) (*DeviceCompliancePolicyPolicySetItemStatus, error) {
	vals := map[string]DeviceCompliancePolicyPolicySetItemStatus{
		"error":          DeviceCompliancePolicyPolicySetItemStatuserror,
		"notassigned":    DeviceCompliancePolicyPolicySetItemStatusnotAssigned,
		"partialsuccess": DeviceCompliancePolicyPolicySetItemStatuspartialSuccess,
		"success":        DeviceCompliancePolicyPolicySetItemStatussuccess,
		"unknown":        DeviceCompliancePolicyPolicySetItemStatusunknown,
		"validating":     DeviceCompliancePolicyPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyPolicySetItemStatus(input)
	return &out, nil
}
