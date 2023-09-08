package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertImpactAggregationType string

const (
	DeviceManagementAlertImpactAggregationTypeaffectedCloudPcCount      DeviceManagementAlertImpactAggregationType = "AffectedCloudPCCount"
	DeviceManagementAlertImpactAggregationTypeaffectedCloudPcPercentage DeviceManagementAlertImpactAggregationType = "AffectedCloudPCPercentage"
	DeviceManagementAlertImpactAggregationTypecount                     DeviceManagementAlertImpactAggregationType = "Count"
	DeviceManagementAlertImpactAggregationTypepercentage                DeviceManagementAlertImpactAggregationType = "Percentage"
)

func PossibleValuesForDeviceManagementAlertImpactAggregationType() []string {
	return []string{
		string(DeviceManagementAlertImpactAggregationTypeaffectedCloudPcCount),
		string(DeviceManagementAlertImpactAggregationTypeaffectedCloudPcPercentage),
		string(DeviceManagementAlertImpactAggregationTypecount),
		string(DeviceManagementAlertImpactAggregationTypepercentage),
	}
}

func parseDeviceManagementAlertImpactAggregationType(input string) (*DeviceManagementAlertImpactAggregationType, error) {
	vals := map[string]DeviceManagementAlertImpactAggregationType{
		"affectedcloudpccount":      DeviceManagementAlertImpactAggregationTypeaffectedCloudPcCount,
		"affectedcloudpcpercentage": DeviceManagementAlertImpactAggregationTypeaffectedCloudPcPercentage,
		"count":                     DeviceManagementAlertImpactAggregationTypecount,
		"percentage":                DeviceManagementAlertImpactAggregationTypepercentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertImpactAggregationType(input)
	return &out, nil
}
