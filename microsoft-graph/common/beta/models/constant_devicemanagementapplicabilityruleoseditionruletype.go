package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsEditionRuleType string

const (
	DeviceManagementApplicabilityRuleOsEditionRuleTypeexclude DeviceManagementApplicabilityRuleOsEditionRuleType = "Exclude"
	DeviceManagementApplicabilityRuleOsEditionRuleTypeinclude DeviceManagementApplicabilityRuleOsEditionRuleType = "Include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsEditionRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsEditionRuleTypeexclude),
		string(DeviceManagementApplicabilityRuleOsEditionRuleTypeinclude),
	}
}

func parseDeviceManagementApplicabilityRuleOsEditionRuleType(input string) (*DeviceManagementApplicabilityRuleOsEditionRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsEditionRuleType{
		"exclude": DeviceManagementApplicabilityRuleOsEditionRuleTypeexclude,
		"include": DeviceManagementApplicabilityRuleOsEditionRuleTypeinclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsEditionRuleType(input)
	return &out, nil
}
