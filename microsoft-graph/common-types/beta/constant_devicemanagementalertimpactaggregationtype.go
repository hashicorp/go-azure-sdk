package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertImpactAggregationType string

const (
	DeviceManagementAlertImpactAggregationType_AffectedCloudPCCount      DeviceManagementAlertImpactAggregationType = "affectedCloudPcCount"
	DeviceManagementAlertImpactAggregationType_AffectedCloudPCPercentage DeviceManagementAlertImpactAggregationType = "affectedCloudPcPercentage"
	DeviceManagementAlertImpactAggregationType_Count                     DeviceManagementAlertImpactAggregationType = "count"
	DeviceManagementAlertImpactAggregationType_Percentage                DeviceManagementAlertImpactAggregationType = "percentage"
)

func PossibleValuesForDeviceManagementAlertImpactAggregationType() []string {
	return []string{
		string(DeviceManagementAlertImpactAggregationType_AffectedCloudPCCount),
		string(DeviceManagementAlertImpactAggregationType_AffectedCloudPCPercentage),
		string(DeviceManagementAlertImpactAggregationType_Count),
		string(DeviceManagementAlertImpactAggregationType_Percentage),
	}
}

func (s *DeviceManagementAlertImpactAggregationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementAlertImpactAggregationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementAlertImpactAggregationType(input string) (*DeviceManagementAlertImpactAggregationType, error) {
	vals := map[string]DeviceManagementAlertImpactAggregationType{
		"affectedcloudpccount":      DeviceManagementAlertImpactAggregationType_AffectedCloudPCCount,
		"affectedcloudpcpercentage": DeviceManagementAlertImpactAggregationType_AffectedCloudPCPercentage,
		"count":                     DeviceManagementAlertImpactAggregationType_Count,
		"percentage":                DeviceManagementAlertImpactAggregationType_Percentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementAlertImpactAggregationType(input)
	return &out, nil
}
