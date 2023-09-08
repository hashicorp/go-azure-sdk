package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleDeviceModeRuleType string

const (
	DeviceManagementApplicabilityRuleDeviceModeRuleTypeexclude DeviceManagementApplicabilityRuleDeviceModeRuleType = "Exclude"
	DeviceManagementApplicabilityRuleDeviceModeRuleTypeinclude DeviceManagementApplicabilityRuleDeviceModeRuleType = "Include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleDeviceModeRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleDeviceModeRuleTypeexclude),
		string(DeviceManagementApplicabilityRuleDeviceModeRuleTypeinclude),
	}
}

func parseDeviceManagementApplicabilityRuleDeviceModeRuleType(input string) (*DeviceManagementApplicabilityRuleDeviceModeRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleDeviceModeRuleType{
		"exclude": DeviceManagementApplicabilityRuleDeviceModeRuleTypeexclude,
		"include": DeviceManagementApplicabilityRuleDeviceModeRuleTypeinclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleDeviceModeRuleType(input)
	return &out, nil
}
