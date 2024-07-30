package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO string

const (
	Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_Disabled      Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO = "disabled"
	Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_Enabled       Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO = "enabled"
	Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_NotConfigured Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO = "notConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_Disabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_Enabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_NotConfigured),
	}
}

func (s *Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseWindows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseWindows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO(input string) (*Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_Disabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_Enabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO(input)
	return &out, nil
}
