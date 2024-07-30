package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_Disabled      Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_Enabled       Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWirelessWideAreaNetwork(input)
	return &out, nil
}
