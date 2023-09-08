package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleThresholdOperator string

const (
	DeviceManagementRuleThresholdOperatorequal          DeviceManagementRuleThresholdOperator = "Equal"
	DeviceManagementRuleThresholdOperatorgreater        DeviceManagementRuleThresholdOperator = "Greater"
	DeviceManagementRuleThresholdOperatorgreaterOrEqual DeviceManagementRuleThresholdOperator = "GreaterOrEqual"
	DeviceManagementRuleThresholdOperatorless           DeviceManagementRuleThresholdOperator = "Less"
	DeviceManagementRuleThresholdOperatorlessOrEqual    DeviceManagementRuleThresholdOperator = "LessOrEqual"
	DeviceManagementRuleThresholdOperatornotEqual       DeviceManagementRuleThresholdOperator = "NotEqual"
)

func PossibleValuesForDeviceManagementRuleThresholdOperator() []string {
	return []string{
		string(DeviceManagementRuleThresholdOperatorequal),
		string(DeviceManagementRuleThresholdOperatorgreater),
		string(DeviceManagementRuleThresholdOperatorgreaterOrEqual),
		string(DeviceManagementRuleThresholdOperatorless),
		string(DeviceManagementRuleThresholdOperatorlessOrEqual),
		string(DeviceManagementRuleThresholdOperatornotEqual),
	}
}

func parseDeviceManagementRuleThresholdOperator(input string) (*DeviceManagementRuleThresholdOperator, error) {
	vals := map[string]DeviceManagementRuleThresholdOperator{
		"equal":          DeviceManagementRuleThresholdOperatorequal,
		"greater":        DeviceManagementRuleThresholdOperatorgreater,
		"greaterorequal": DeviceManagementRuleThresholdOperatorgreaterOrEqual,
		"less":           DeviceManagementRuleThresholdOperatorless,
		"lessorequal":    DeviceManagementRuleThresholdOperatorlessOrEqual,
		"notequal":       DeviceManagementRuleThresholdOperatornotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementRuleThresholdOperator(input)
	return &out, nil
}
