package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceRearCamera string

const (
	Windows10DeviceFirmwareConfigurationInterfaceRearCamera_Disabled      Windows10DeviceFirmwareConfigurationInterfaceRearCamera = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceRearCamera_Enabled       Windows10DeviceFirmwareConfigurationInterfaceRearCamera = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceRearCamera_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceRearCamera = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceRearCamera() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceRearCamera_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceRearCamera_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceRearCamera_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceRearCamera) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceRearCamera(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceRearCamera(input string) (*Windows10DeviceFirmwareConfigurationInterfaceRearCamera, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceRearCamera{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceRearCamera_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceRearCamera_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceRearCamera_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceRearCamera(input)
	return &out, nil
}
