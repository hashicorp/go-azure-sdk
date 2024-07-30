package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleDeviceModeRuleType string

const (
	DeviceManagementApplicabilityRuleDeviceModeRuleType_Exclude DeviceManagementApplicabilityRuleDeviceModeRuleType = "exclude"
	DeviceManagementApplicabilityRuleDeviceModeRuleType_Include DeviceManagementApplicabilityRuleDeviceModeRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleDeviceModeRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleDeviceModeRuleType_Exclude),
		string(DeviceManagementApplicabilityRuleDeviceModeRuleType_Include),
	}
}

func (s *DeviceManagementApplicabilityRuleDeviceModeRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleDeviceModeRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleDeviceModeRuleType(input string) (*DeviceManagementApplicabilityRuleDeviceModeRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleDeviceModeRuleType{
		"exclude": DeviceManagementApplicabilityRuleDeviceModeRuleType_Exclude,
		"include": DeviceManagementApplicabilityRuleDeviceModeRuleType_Include,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleDeviceModeRuleType(input)
	return &out, nil
}
