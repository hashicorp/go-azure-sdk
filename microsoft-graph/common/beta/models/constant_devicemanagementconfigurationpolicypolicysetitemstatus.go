package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyPolicySetItemStatus string

const (
	DeviceManagementConfigurationPolicyPolicySetItemStatuserror          DeviceManagementConfigurationPolicyPolicySetItemStatus = "Error"
	DeviceManagementConfigurationPolicyPolicySetItemStatusnotAssigned    DeviceManagementConfigurationPolicyPolicySetItemStatus = "NotAssigned"
	DeviceManagementConfigurationPolicyPolicySetItemStatuspartialSuccess DeviceManagementConfigurationPolicyPolicySetItemStatus = "PartialSuccess"
	DeviceManagementConfigurationPolicyPolicySetItemStatussuccess        DeviceManagementConfigurationPolicyPolicySetItemStatus = "Success"
	DeviceManagementConfigurationPolicyPolicySetItemStatusunknown        DeviceManagementConfigurationPolicyPolicySetItemStatus = "Unknown"
	DeviceManagementConfigurationPolicyPolicySetItemStatusvalidating     DeviceManagementConfigurationPolicyPolicySetItemStatus = "Validating"
)

func PossibleValuesForDeviceManagementConfigurationPolicyPolicySetItemStatus() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyPolicySetItemStatuserror),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatusnotAssigned),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatuspartialSuccess),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatussuccess),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatusunknown),
		string(DeviceManagementConfigurationPolicyPolicySetItemStatusvalidating),
	}
}

func parseDeviceManagementConfigurationPolicyPolicySetItemStatus(input string) (*DeviceManagementConfigurationPolicyPolicySetItemStatus, error) {
	vals := map[string]DeviceManagementConfigurationPolicyPolicySetItemStatus{
		"error":          DeviceManagementConfigurationPolicyPolicySetItemStatuserror,
		"notassigned":    DeviceManagementConfigurationPolicyPolicySetItemStatusnotAssigned,
		"partialsuccess": DeviceManagementConfigurationPolicyPolicySetItemStatuspartialSuccess,
		"success":        DeviceManagementConfigurationPolicyPolicySetItemStatussuccess,
		"unknown":        DeviceManagementConfigurationPolicyPolicySetItemStatusunknown,
		"validating":     DeviceManagementConfigurationPolicyPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyPolicySetItemStatus(input)
	return &out, nil
}
