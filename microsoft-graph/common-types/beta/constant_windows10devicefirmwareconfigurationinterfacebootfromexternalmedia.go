package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia string

const (
	Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_Disabled      Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_Enabled       Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia(input string) (*Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia(input)
	return &out, nil
}
