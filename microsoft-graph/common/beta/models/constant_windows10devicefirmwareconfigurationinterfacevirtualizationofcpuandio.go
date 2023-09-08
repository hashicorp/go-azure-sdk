package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO string

const (
	Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOdisabled      Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOenabled       Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOnotConfigured Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO(input string) (*Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIOnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceVirtualizationOfCpuAndIO(input)
	return &out, nil
}
