package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateSystemManagementMode string

const (
	DeviceHealthAttestationStateSystemManagementMode_Level1        DeviceHealthAttestationStateSystemManagementMode = "level1"
	DeviceHealthAttestationStateSystemManagementMode_Level2        DeviceHealthAttestationStateSystemManagementMode = "level2"
	DeviceHealthAttestationStateSystemManagementMode_Level3        DeviceHealthAttestationStateSystemManagementMode = "level3"
	DeviceHealthAttestationStateSystemManagementMode_NotApplicable DeviceHealthAttestationStateSystemManagementMode = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateSystemManagementMode() []string {
	return []string{
		string(DeviceHealthAttestationStateSystemManagementMode_Level1),
		string(DeviceHealthAttestationStateSystemManagementMode_Level2),
		string(DeviceHealthAttestationStateSystemManagementMode_Level3),
		string(DeviceHealthAttestationStateSystemManagementMode_NotApplicable),
	}
}

func (s *DeviceHealthAttestationStateSystemManagementMode) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateSystemManagementMode(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateSystemManagementMode(input string) (*DeviceHealthAttestationStateSystemManagementMode, error) {
	vals := map[string]DeviceHealthAttestationStateSystemManagementMode{
		"level1":        DeviceHealthAttestationStateSystemManagementMode_Level1,
		"level2":        DeviceHealthAttestationStateSystemManagementMode_Level2,
		"level3":        DeviceHealthAttestationStateSystemManagementMode_Level3,
		"notapplicable": DeviceHealthAttestationStateSystemManagementMode_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateSystemManagementMode(input)
	return &out, nil
}
