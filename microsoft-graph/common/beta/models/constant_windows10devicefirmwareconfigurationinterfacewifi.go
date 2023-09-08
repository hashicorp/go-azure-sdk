package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWiFi string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWiFidisabled      Windows10DeviceFirmwareConfigurationInterfaceWiFi = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWiFienabled       Windows10DeviceFirmwareConfigurationInterfaceWiFi = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWiFinotConfigured Windows10DeviceFirmwareConfigurationInterfaceWiFi = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWiFi() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWiFidisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWiFienabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWiFinotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWiFi(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWiFi, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWiFi{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWiFidisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWiFienabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWiFinotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWiFi(input)
	return &out, nil
}
