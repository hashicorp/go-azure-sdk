package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_Disabled      Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_Enabled       Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWakeOnLAN() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceWakeOnLAN(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWakeOnLAN(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN(input)
	return &out, nil
}
