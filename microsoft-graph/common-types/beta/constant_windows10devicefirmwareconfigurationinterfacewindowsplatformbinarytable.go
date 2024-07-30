package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_Disabled      Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_Enabled       Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable(input)
	return &out, nil
}
