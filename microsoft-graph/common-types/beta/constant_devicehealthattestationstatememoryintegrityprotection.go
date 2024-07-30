package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateMemoryIntegrityProtection string

const (
	DeviceHealthAttestationStateMemoryIntegrityProtection_Disabled      DeviceHealthAttestationStateMemoryIntegrityProtection = "disabled"
	DeviceHealthAttestationStateMemoryIntegrityProtection_Enabled       DeviceHealthAttestationStateMemoryIntegrityProtection = "enabled"
	DeviceHealthAttestationStateMemoryIntegrityProtection_NotApplicable DeviceHealthAttestationStateMemoryIntegrityProtection = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateMemoryIntegrityProtection() []string {
	return []string{
		string(DeviceHealthAttestationStateMemoryIntegrityProtection_Disabled),
		string(DeviceHealthAttestationStateMemoryIntegrityProtection_Enabled),
		string(DeviceHealthAttestationStateMemoryIntegrityProtection_NotApplicable),
	}
}

func (s *DeviceHealthAttestationStateMemoryIntegrityProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateMemoryIntegrityProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateMemoryIntegrityProtection(input string) (*DeviceHealthAttestationStateMemoryIntegrityProtection, error) {
	vals := map[string]DeviceHealthAttestationStateMemoryIntegrityProtection{
		"disabled":      DeviceHealthAttestationStateMemoryIntegrityProtection_Disabled,
		"enabled":       DeviceHealthAttestationStateMemoryIntegrityProtection_Enabled,
		"notapplicable": DeviceHealthAttestationStateMemoryIntegrityProtection_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateMemoryIntegrityProtection(input)
	return &out, nil
}
