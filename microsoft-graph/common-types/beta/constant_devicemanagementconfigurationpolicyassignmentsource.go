package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyAssignmentSource string

const (
	DeviceManagementConfigurationPolicyAssignmentSource_Direct     DeviceManagementConfigurationPolicyAssignmentSource = "direct"
	DeviceManagementConfigurationPolicyAssignmentSource_PolicySets DeviceManagementConfigurationPolicyAssignmentSource = "policySets"
)

func PossibleValuesForDeviceManagementConfigurationPolicyAssignmentSource() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyAssignmentSource_Direct),
		string(DeviceManagementConfigurationPolicyAssignmentSource_PolicySets),
	}
}

func (s *DeviceManagementConfigurationPolicyAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceManagementConfigurationPolicyAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceManagementConfigurationPolicyAssignmentSource(input string) (*DeviceManagementConfigurationPolicyAssignmentSource, error) {
	vals := map[string]DeviceManagementConfigurationPolicyAssignmentSource{
		"direct":     DeviceManagementConfigurationPolicyAssignmentSource_Direct,
		"policysets": DeviceManagementConfigurationPolicyAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyAssignmentSource(input)
	return &out, nil
}
