package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationPolicySetItemStatus string

const (
	DeviceConfigurationPolicySetItemStatuserror          DeviceConfigurationPolicySetItemStatus = "Error"
	DeviceConfigurationPolicySetItemStatusnotAssigned    DeviceConfigurationPolicySetItemStatus = "NotAssigned"
	DeviceConfigurationPolicySetItemStatuspartialSuccess DeviceConfigurationPolicySetItemStatus = "PartialSuccess"
	DeviceConfigurationPolicySetItemStatussuccess        DeviceConfigurationPolicySetItemStatus = "Success"
	DeviceConfigurationPolicySetItemStatusunknown        DeviceConfigurationPolicySetItemStatus = "Unknown"
	DeviceConfigurationPolicySetItemStatusvalidating     DeviceConfigurationPolicySetItemStatus = "Validating"
)

func PossibleValuesForDeviceConfigurationPolicySetItemStatus() []string {
	return []string{
		string(DeviceConfigurationPolicySetItemStatuserror),
		string(DeviceConfigurationPolicySetItemStatusnotAssigned),
		string(DeviceConfigurationPolicySetItemStatuspartialSuccess),
		string(DeviceConfigurationPolicySetItemStatussuccess),
		string(DeviceConfigurationPolicySetItemStatusunknown),
		string(DeviceConfigurationPolicySetItemStatusvalidating),
	}
}

func parseDeviceConfigurationPolicySetItemStatus(input string) (*DeviceConfigurationPolicySetItemStatus, error) {
	vals := map[string]DeviceConfigurationPolicySetItemStatus{
		"error":          DeviceConfigurationPolicySetItemStatuserror,
		"notassigned":    DeviceConfigurationPolicySetItemStatusnotAssigned,
		"partialsuccess": DeviceConfigurationPolicySetItemStatuspartialSuccess,
		"success":        DeviceConfigurationPolicySetItemStatussuccess,
		"unknown":        DeviceConfigurationPolicySetItemStatusunknown,
		"validating":     DeviceConfigurationPolicySetItemStatusvalidating,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationPolicySetItemStatus(input)
	return &out, nil
}
