package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceFrontCamera string

const (
	Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_Disabled      Windows10DeviceFirmwareConfigurationInterfaceFrontCamera = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_Enabled       Windows10DeviceFirmwareConfigurationInterfaceFrontCamera = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceFrontCamera = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceFrontCamera() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceFrontCamera) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceFrontCamera(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceFrontCamera(input string) (*Windows10DeviceFirmwareConfigurationInterfaceFrontCamera, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceFrontCamera{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceFrontCamera_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceFrontCamera(input)
	return &out, nil
}
