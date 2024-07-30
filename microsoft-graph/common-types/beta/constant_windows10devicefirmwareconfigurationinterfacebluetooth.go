package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceBluetooth string

const (
	Windows10DeviceFirmwareConfigurationInterfaceBluetooth_Disabled      Windows10DeviceFirmwareConfigurationInterfaceBluetooth = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceBluetooth_Enabled       Windows10DeviceFirmwareConfigurationInterfaceBluetooth = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceBluetooth_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceBluetooth = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceBluetooth() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceBluetooth_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBluetooth_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBluetooth_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceBluetooth) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceBluetooth(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceBluetooth(input string) (*Windows10DeviceFirmwareConfigurationInterfaceBluetooth, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceBluetooth{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceBluetooth_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceBluetooth_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceBluetooth_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceBluetooth(input)
	return &out, nil
}
