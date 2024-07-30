package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera string

const (
	Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_Disabled      Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_Enabled       Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceInfraredCamera() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceInfraredCamera(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceInfraredCamera(input string) (*Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceInfraredCamera(input)
	return &out, nil
}
