package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceRadios string

const (
	Windows10DeviceFirmwareConfigurationInterfaceRadios_Disabled      Windows10DeviceFirmwareConfigurationInterfaceRadios = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceRadios_Enabled       Windows10DeviceFirmwareConfigurationInterfaceRadios = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceRadios_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceRadios = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceRadios() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceRadios_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceRadios_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceRadios_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceRadios) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceRadios(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceRadios(input string) (*Windows10DeviceFirmwareConfigurationInterfaceRadios, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceRadios{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceRadios_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceRadios_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceRadios_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceRadios(input)
	return &out, nil
}
