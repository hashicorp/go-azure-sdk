package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsVersionRuleType string

const (
	DeviceManagementApplicabilityRuleOsVersionRuleTypeexclude DeviceManagementApplicabilityRuleOsVersionRuleType = "Exclude"
	DeviceManagementApplicabilityRuleOsVersionRuleTypeinclude DeviceManagementApplicabilityRuleOsVersionRuleType = "Include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsVersionRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsVersionRuleTypeexclude),
		string(DeviceManagementApplicabilityRuleOsVersionRuleTypeinclude),
	}
}

func parseDeviceManagementApplicabilityRuleOsVersionRuleType(input string) (*DeviceManagementApplicabilityRuleOsVersionRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsVersionRuleType{
		"exclude": DeviceManagementApplicabilityRuleOsVersionRuleTypeexclude,
		"include": DeviceManagementApplicabilityRuleOsVersionRuleTypeinclude,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsVersionRuleType(input)
	return &out, nil
}
