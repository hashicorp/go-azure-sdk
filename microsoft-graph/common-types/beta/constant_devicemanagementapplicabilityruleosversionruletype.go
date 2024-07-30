package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementApplicabilityRuleOsVersionRuleType string

const (
	DeviceManagementApplicabilityRuleOsVersionRuleType_Exclude DeviceManagementApplicabilityRuleOsVersionRuleType = "exclude"
	DeviceManagementApplicabilityRuleOsVersionRuleType_Include DeviceManagementApplicabilityRuleOsVersionRuleType = "include"
)

func PossibleValuesForDeviceManagementApplicabilityRuleOsVersionRuleType() []string {
	return []string{
		string(DeviceManagementApplicabilityRuleOsVersionRuleType_Exclude),
		string(DeviceManagementApplicabilityRuleOsVersionRuleType_Include),
	}
}

func (s *DeviceManagementApplicabilityRuleOsVersionRuleType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementApplicabilityRuleOsVersionRuleType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementApplicabilityRuleOsVersionRuleType(input string) (*DeviceManagementApplicabilityRuleOsVersionRuleType, error) {
	vals := map[string]DeviceManagementApplicabilityRuleOsVersionRuleType{
		"exclude": DeviceManagementApplicabilityRuleOsVersionRuleType_Exclude,
		"include": DeviceManagementApplicabilityRuleOsVersionRuleType_Include,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementApplicabilityRuleOsVersionRuleType(input)
	return &out, nil
}
