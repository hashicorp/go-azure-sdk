package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceHealthAttestationStateVirtualizationBasedSecurity string

const (
	DeviceHealthAttestationStateVirtualizationBasedSecurity_Disabled      DeviceHealthAttestationStateVirtualizationBasedSecurity = "disabled"
	DeviceHealthAttestationStateVirtualizationBasedSecurity_Enabled       DeviceHealthAttestationStateVirtualizationBasedSecurity = "enabled"
	DeviceHealthAttestationStateVirtualizationBasedSecurity_NotApplicable DeviceHealthAttestationStateVirtualizationBasedSecurity = "notApplicable"
)

func PossibleValuesForDeviceHealthAttestationStateVirtualizationBasedSecurity() []string {
	return []string{
		string(DeviceHealthAttestationStateVirtualizationBasedSecurity_Disabled),
		string(DeviceHealthAttestationStateVirtualizationBasedSecurity_Enabled),
		string(DeviceHealthAttestationStateVirtualizationBasedSecurity_NotApplicable),
	}
}

func (s *DeviceHealthAttestationStateVirtualizationBasedSecurity) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceHealthAttestationStateVirtualizationBasedSecurity(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceHealthAttestationStateVirtualizationBasedSecurity(input string) (*DeviceHealthAttestationStateVirtualizationBasedSecurity, error) {
	vals := map[string]DeviceHealthAttestationStateVirtualizationBasedSecurity{
		"disabled":      DeviceHealthAttestationStateVirtualizationBasedSecurity_Disabled,
		"enabled":       DeviceHealthAttestationStateVirtualizationBasedSecurity_Enabled,
		"notapplicable": DeviceHealthAttestationStateVirtualizationBasedSecurity_NotApplicable,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceHealthAttestationStateVirtualizationBasedSecurity(input)
	return &out, nil
}
