package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationState string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationState_Disabled      DeviceEnrollmentWindowsHelloForBusinessConfigurationState = "disabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationState_Enabled       DeviceEnrollmentWindowsHelloForBusinessConfigurationState = "enabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationState_NotConfigured DeviceEnrollmentWindowsHelloForBusinessConfigurationState = "notConfigured"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationState() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationState_Disabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationState_Enabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationState_NotConfigured),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationState(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationState, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationState{
		"disabled":      DeviceEnrollmentWindowsHelloForBusinessConfigurationState_Disabled,
		"enabled":       DeviceEnrollmentWindowsHelloForBusinessConfigurationState_Enabled,
		"notconfigured": DeviceEnrollmentWindowsHelloForBusinessConfigurationState_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationState(input)
	return &out, nil
}
