package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters string

const (
	Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_Disabled      Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_Enabled       Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters(input string) (*Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters(input)
	return &out, nil
}
