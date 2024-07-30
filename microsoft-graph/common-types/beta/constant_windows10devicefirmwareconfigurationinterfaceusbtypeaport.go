package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort string

const (
	Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_Disabled      Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_Enabled       Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort(input string) (*Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort(input)
	return &out, nil
}
