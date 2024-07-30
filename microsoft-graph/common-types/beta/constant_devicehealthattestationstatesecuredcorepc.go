package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateSecuredCorePC string

const (
	DeviceHealthAttestationStateSecuredCorePC_Disabled      DeviceHealthAttestationStateSecuredCorePC = "disabled"
	DeviceHealthAttestationStateSecuredCorePC_Enabled       DeviceHealthAttestationStateSecuredCorePC = "enabled"
	DeviceHealthAttestationStateSecuredCorePC_NotApplicable DeviceHealthAttestationStateSecuredCorePC = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateSecuredCorePC() []string {
	return []string{
		string(DeviceHealthAttestationStateSecuredCorePC_Disabled),
		string(DeviceHealthAttestationStateSecuredCorePC_Enabled),
		string(DeviceHealthAttestationStateSecuredCorePC_NotApplicable),
	}
}

func (s *DeviceHealthAttestationStateSecuredCorePC) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateSecuredCorePC(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateSecuredCorePC(input string) (*DeviceHealthAttestationStateSecuredCorePC, error) {
	vals := map[string]DeviceHealthAttestationStateSecuredCorePC{
		"disabled":      DeviceHealthAttestationStateSecuredCorePC_Disabled,
		"enabled":       DeviceHealthAttestationStateSecuredCorePC_Enabled,
		"notapplicable": DeviceHealthAttestationStateSecuredCorePC_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateSecuredCorePC(input)
	return &out, nil
}
