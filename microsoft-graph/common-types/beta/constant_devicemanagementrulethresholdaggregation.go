package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleThresholdAggregation string

const (
	DeviceManagementRuleThresholdAggregation_AffectedCloudPCCount      DeviceManagementRuleThresholdAggregation = "affectedCloudPcCount"
	DeviceManagementRuleThresholdAggregation_AffectedCloudPCPercentage DeviceManagementRuleThresholdAggregation = "affectedCloudPcPercentage"
	DeviceManagementRuleThresholdAggregation_Count                     DeviceManagementRuleThresholdAggregation = "count"
	DeviceManagementRuleThresholdAggregation_Percentage                DeviceManagementRuleThresholdAggregation = "percentage"
)

func PossibleValuesForDeviceManagementRuleThresholdAggregation() []string {
	return []string{
		string(DeviceManagementRuleThresholdAggregation_AffectedCloudPCCount),
		string(DeviceManagementRuleThresholdAggregation_AffectedCloudPCPercentage),
		string(DeviceManagementRuleThresholdAggregation_Count),
		string(DeviceManagementRuleThresholdAggregation_Percentage),
	}
}

func (s *DeviceManagementRuleThresholdAggregation) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementRuleThresholdAggregation(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementRuleThresholdAggregation(input string) (*DeviceManagementRuleThresholdAggregation, error) {
	vals := map[string]DeviceManagementRuleThresholdAggregation{
		"affectedcloudpccount":      DeviceManagementRuleThresholdAggregation_AffectedCloudPCCount,
		"affectedcloudpcpercentage": DeviceManagementRuleThresholdAggregation_AffectedCloudPCPercentage,
		"count":                     DeviceManagementRuleThresholdAggregation_Count,
		"percentage":                DeviceManagementRuleThresholdAggregation_Percentage,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementRuleThresholdAggregation(input)
	return &out, nil
}
