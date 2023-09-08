package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleThresholdAggregation string

const (
	DeviceManagementRuleThresholdAggregationaffectedCloudPcCount      DeviceManagementRuleThresholdAggregation = "AffectedCloudPCCount"
	DeviceManagementRuleThresholdAggregationaffectedCloudPcPercentage DeviceManagementRuleThresholdAggregation = "AffectedCloudPCPercentage"
	DeviceManagementRuleThresholdAggregationcount                     DeviceManagementRuleThresholdAggregation = "Count"
	DeviceManagementRuleThresholdAggregationpercentage                DeviceManagementRuleThresholdAggregation = "Percentage"
)

func PossibleValuesForDeviceManagementRuleThresholdAggregation() []string {
	return []string{
		string(DeviceManagementRuleThresholdAggregationaffectedCloudPcCount),
		string(DeviceManagementRuleThresholdAggregationaffectedCloudPcPercentage),
		string(DeviceManagementRuleThresholdAggregationcount),
		string(DeviceManagementRuleThresholdAggregationpercentage),
	}
}

func parseDeviceManagementRuleThresholdAggregation(input string) (*DeviceManagementRuleThresholdAggregation, error) {
	vals := map[string]DeviceManagementRuleThresholdAggregation{
		"affectedcloudpccount":      DeviceManagementRuleThresholdAggregationaffectedCloudPcCount,
		"affectedcloudpcpercentage": DeviceManagementRuleThresholdAggregationaffectedCloudPcPercentage,
		"count":                     DeviceManagementRuleThresholdAggregationcount,
		"percentage":                DeviceManagementRuleThresholdAggregationpercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementRuleThresholdAggregation(input)
	return &out, nil
}
