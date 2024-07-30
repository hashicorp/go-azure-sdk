package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication string

const (
	Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_Disabled      Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_Enabled       Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication(input string) (*Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication(input)
	return &out, nil
}
