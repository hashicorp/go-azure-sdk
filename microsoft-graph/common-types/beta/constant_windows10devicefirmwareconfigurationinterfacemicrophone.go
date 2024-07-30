package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceMicrophone string

const (
	Windows10DeviceFirmwareConfigurationInterfaceMicrophone_Disabled      Windows10DeviceFirmwareConfigurationInterfaceMicrophone = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophone_Enabled       Windows10DeviceFirmwareConfigurationInterfaceMicrophone = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophone_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceMicrophone = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceMicrophone() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophone_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophone_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophone_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceMicrophone) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceMicrophone(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceMicrophone(input string) (*Windows10DeviceFirmwareConfigurationInterfaceMicrophone, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceMicrophone{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceMicrophone_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceMicrophone_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceMicrophone_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceMicrophone(input)
	return &out, nil
}
