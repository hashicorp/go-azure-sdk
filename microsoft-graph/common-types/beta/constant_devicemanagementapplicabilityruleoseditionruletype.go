package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsEditionRuleType string

const (
	DeviceManagementApplicabilityRuleOsEditionRuleType_Exclude DeviceManagementApplicabilityRuleOsEditionRuleType = "exclude"
	DeviceManagementApplicabilityRuleOsEditionRuleType_Include DeviceManagementApplicabilityRuleOsEditionRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsEditionRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsEditionRuleType_Exclude),
		string(DeviceManagementApplicabilityRuleOsEditionRuleType_Include),
	}
}

func (s *DeviceManagementApplicabilityRuleOsEditionRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleOsEditionRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleOsEditionRuleType(input string) (*DeviceManagementApplicabilityRuleOsEditionRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsEditionRuleType{
		"exclude": DeviceManagementApplicabilityRuleOsEditionRuleType_Exclude,
		"include": DeviceManagementApplicabilityRuleOsEditionRuleType_Include,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsEditionRuleType(input)
	return &out, nil
}
