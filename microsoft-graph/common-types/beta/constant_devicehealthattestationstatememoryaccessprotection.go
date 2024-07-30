package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateMemoryAccessProtection string

const (
	DeviceHealthAttestationStateMemoryAccessProtection_Disabled      DeviceHealthAttestationStateMemoryAccessProtection = "disabled"
	DeviceHealthAttestationStateMemoryAccessProtection_Enabled       DeviceHealthAttestationStateMemoryAccessProtection = "enabled"
	DeviceHealthAttestationStateMemoryAccessProtection_NotApplicable DeviceHealthAttestationStateMemoryAccessProtection = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateMemoryAccessProtection() []string {
	return []string{
		string(DeviceHealthAttestationStateMemoryAccessProtection_Disabled),
		string(DeviceHealthAttestationStateMemoryAccessProtection_Enabled),
		string(DeviceHealthAttestationStateMemoryAccessProtection_NotApplicable),
	}
}

func (s *DeviceHealthAttestationStateMemoryAccessProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateMemoryAccessProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateMemoryAccessProtection(input string) (*DeviceHealthAttestationStateMemoryAccessProtection, error) {
	vals := map[string]DeviceHealthAttestationStateMemoryAccessProtection{
		"disabled":      DeviceHealthAttestationStateMemoryAccessProtection_Disabled,
		"enabled":       DeviceHealthAttestationStateMemoryAccessProtection_Enabled,
		"notapplicable": DeviceHealthAttestationStateMemoryAccessProtection_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateMemoryAccessProtection(input)
	return &out, nil
}
