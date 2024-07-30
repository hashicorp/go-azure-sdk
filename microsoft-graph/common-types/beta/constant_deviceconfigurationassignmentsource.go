package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationAssignmentSource string

const (
	DeviceConfigurationAssignmentSource_Direct     DeviceConfigurationAssignmentSource = "direct"
	DeviceConfigurationAssignmentSource_PolicySets DeviceConfigurationAssignmentSource = "policySets"
)

func PossibleValuesForDeviceConfigurationAssignmentSource() []string {
	return []string{
		string(DeviceConfigurationAssignmentSource_Direct),
		string(DeviceConfigurationAssignmentSource_PolicySets),
	}
}

func (s *DeviceConfigurationAssignmentSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceConfigurationAssignmentSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceConfigurationAssignmentSource(input string) (*DeviceConfigurationAssignmentSource, error) {
	vals := map[string]DeviceConfigurationAssignmentSource{
		"direct":     DeviceConfigurationAssignmentSource_Direct,
		"policysets": DeviceConfigurationAssignmentSource_PolicySets,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationAssignmentSource(input)
	return &out, nil
}
