package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANdisabled      Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANenabled       Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANnotConfigured Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWakeOnLAN() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWakeOnLAN(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWakeOnLANnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWakeOnLAN(input)
	return &out, nil
}
