package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState string

const (
	DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_Disabled      DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState = "disabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_Enabled       DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState = "enabled"
	DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_NotConfigured DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState = "notConfigured"
)

func PossibleValuesForDeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState() []string {
	return []string{
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_Disabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_Enabled),
		string(DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_NotConfigured),
	}
}

func (s *DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState(input string) (*DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState, error) {
	vals := map[string]DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState{
		"disabled":      DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_Disabled,
		"enabled":       DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_Enabled,
		"notconfigured": DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceEnrollmentWindowsHelloForBusinessConfigurationEnhancedBiometricsState(input)
	return &out, nil
}
