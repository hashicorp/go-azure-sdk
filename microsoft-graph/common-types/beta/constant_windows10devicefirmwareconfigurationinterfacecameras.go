package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceCameras string

const (
	Windows10DeviceFirmwareConfigurationInterfaceCameras_Disabled      Windows10DeviceFirmwareConfigurationInterfaceCameras = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceCameras_Enabled       Windows10DeviceFirmwareConfigurationInterfaceCameras = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceCameras_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceCameras = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceCameras() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceCameras_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceCameras_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceCameras_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceCameras) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceCameras(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceCameras(input string) (*Windows10DeviceFirmwareConfigurationInterfaceCameras, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceCameras{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceCameras_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceCameras_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceCameras_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceCameras(input)
	return &out, nil
}
