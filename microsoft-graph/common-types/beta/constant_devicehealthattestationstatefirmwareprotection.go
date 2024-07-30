package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateFirmwareProtection string

const (
	DeviceHealthAttestationStateFirmwareProtection_Disabled                       DeviceHealthAttestationStateFirmwareProtection = "disabled"
	DeviceHealthAttestationStateFirmwareProtection_FirmwareAttackSurfaceReduction DeviceHealthAttestationStateFirmwareProtection = "firmwareAttackSurfaceReduction"
	DeviceHealthAttestationStateFirmwareProtection_NotApplicable                  DeviceHealthAttestationStateFirmwareProtection = "notApplicable"
	DeviceHealthAttestationStateFirmwareProtection_SystemGuardSecureLaunch        DeviceHealthAttestationStateFirmwareProtection = "systemGuardSecureLaunch"
)

func PossibleValuesForDeviceHealthAttestationStateFirmwareProtection() []string {
	return []string{
		string(DeviceHealthAttestationStateFirmwareProtection_Disabled),
		string(DeviceHealthAttestationStateFirmwareProtection_FirmwareAttackSurfaceReduction),
		string(DeviceHealthAttestationStateFirmwareProtection_NotApplicable),
		string(DeviceHealthAttestationStateFirmwareProtection_SystemGuardSecureLaunch),
	}
}

func (s *DeviceHealthAttestationStateFirmwareProtection) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateFirmwareProtection(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateFirmwareProtection(input string) (*DeviceHealthAttestationStateFirmwareProtection, error) {
	vals := map[string]DeviceHealthAttestationStateFirmwareProtection{
		"disabled":                       DeviceHealthAttestationStateFirmwareProtection_Disabled,
		"firmwareattacksurfacereduction": DeviceHealthAttestationStateFirmwareProtection_FirmwareAttackSurfaceReduction,
		"notapplicable":                  DeviceHealthAttestationStateFirmwareProtection_NotApplicable,
		"systemguardsecurelaunch":        DeviceHealthAttestationStateFirmwareProtection_SystemGuardSecureLaunch,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateFirmwareProtection(input)
	return &out, nil
}
