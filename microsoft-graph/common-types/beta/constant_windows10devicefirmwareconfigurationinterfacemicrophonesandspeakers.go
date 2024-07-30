package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers string

const (
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_Disabled      Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_Enabled       Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers(input string) (*Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceMicrophonesAndSpeakers(input)
	return &out, nil
}
