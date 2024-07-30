package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_Disabled      Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_Enabled       Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWakeOnPower() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceWakeOnPower(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWakeOnPower(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWakeOnPower(input)
	return &out, nil
}
