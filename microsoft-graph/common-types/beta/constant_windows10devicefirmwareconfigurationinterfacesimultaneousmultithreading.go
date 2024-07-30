package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading string

const (
	Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_Disabled      Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_Enabled       Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading(input string) (*Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading(input)
	return &out, nil
}
