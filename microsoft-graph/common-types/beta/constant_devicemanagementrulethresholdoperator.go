package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementRuleThresholdOperator string

const (
	DeviceManagementRuleThresholdOperator_Equal          DeviceManagementRuleThresholdOperator = "equal"
	DeviceManagementRuleThresholdOperator_Greater        DeviceManagementRuleThresholdOperator = "greater"
	DeviceManagementRuleThresholdOperator_GreaterOrEqual DeviceManagementRuleThresholdOperator = "greaterOrEqual"
	DeviceManagementRuleThresholdOperator_Less           DeviceManagementRuleThresholdOperator = "less"
	DeviceManagementRuleThresholdOperator_LessOrEqual    DeviceManagementRuleThresholdOperator = "lessOrEqual"
	DeviceManagementRuleThresholdOperator_NotEqual       DeviceManagementRuleThresholdOperator = "notEqual"
)

func PossibleValuesForDeviceManagementRuleThresholdOperator() []string {
	return []string{
		string(DeviceManagementRuleThresholdOperator_Equal),
		string(DeviceManagementRuleThresholdOperator_Greater),
		string(DeviceManagementRuleThresholdOperator_GreaterOrEqual),
		string(DeviceManagementRuleThresholdOperator_Less),
		string(DeviceManagementRuleThresholdOperator_LessOrEqual),
		string(DeviceManagementRuleThresholdOperator_NotEqual),
	}
}

func (s *DeviceManagementRuleThresholdOperator) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementRuleThresholdOperator(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementRuleThresholdOperator(input string) (*DeviceManagementRuleThresholdOperator, error) {
	vals := map[string]DeviceManagementRuleThresholdOperator{
		"equal":          DeviceManagementRuleThresholdOperator_Equal,
		"greater":        DeviceManagementRuleThresholdOperator_Greater,
		"greaterorequal": DeviceManagementRuleThresholdOperator_GreaterOrEqual,
		"less":           DeviceManagementRuleThresholdOperator_Less,
		"lessorequal":    DeviceManagementRuleThresholdOperator_LessOrEqual,
		"notequal":       DeviceManagementRuleThresholdOperator_NotEqual,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementRuleThresholdOperator(input)
	return &out, nil
}
