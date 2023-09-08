package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort string

const (
	Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortdisabled      Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortenabled       Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortnotConfigured Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort(input string) (*Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPortnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceUsbTypeAPort(input)
	return &out, nil
}
