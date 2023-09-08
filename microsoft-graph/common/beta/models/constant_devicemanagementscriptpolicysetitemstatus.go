package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementScriptPolicySetItemStatus string

const (
	DeviceManagementScriptPolicySetItemStatuserror          DeviceManagementScriptPolicySetItemStatus = "Error"
	DeviceManagementScriptPolicySetItemStatusnotAssigned    DeviceManagementScriptPolicySetItemStatus = "NotAssigned"
	DeviceManagementScriptPolicySetItemStatuspartialSuccess DeviceManagementScriptPolicySetItemStatus = "PartialSuccess"
	DeviceManagementScriptPolicySetItemStatussuccess        DeviceManagementScriptPolicySetItemStatus = "Success"
	DeviceManagementScriptPolicySetItemStatusunknown        DeviceManagementScriptPolicySetItemStatus = "Unknown"
	DeviceManagementScriptPolicySetItemStatusvalidating     DeviceManagementScriptPolicySetItemStatus = "Validating"
)

func PossibleValuesForDeviceManagementScriptPolicySetItemStatus() []string {
	return []string{
		string(DeviceManagementScriptPolicySetItemStatuserror),
		string(DeviceManagementScriptPolicySetItemStatusnotAssigned),
		string(DeviceManagementScriptPolicySetItemStatuspartialSuccess),
		string(DeviceManagementScriptPolicySetItemStatussuccess),
		string(DeviceManagementScriptPolicySetItemStatusunknown),
		string(DeviceManagementScriptPolicySetItemStatusvalidating),
	}
}

func parseDeviceManagementScriptPolicySetItemStatus(input string) (*DeviceManagementScriptPolicySetItemStatus, error) {
	vals := map[string]DeviceManagementScriptPolicySetItemStatus{
		"error":          DeviceManagementScriptPolicySetItemStatuserror,
		"notassigned":    DeviceManagementScriptPolicySetItemStatusnotAssigned,
		"partialsuccess": DeviceManagementScriptPolicySetItemStatuspartialSuccess,
		"success":        DeviceManagementScriptPolicySetItemStatussuccess,
		"unknown":        DeviceManagementScriptPolicySetItemStatusunknown,
		"validating":     DeviceManagementScriptPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementScriptPolicySetItemStatus(input)
	return &out, nil
}
