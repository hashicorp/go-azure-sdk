package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceSdCard string

const (
	Windows10DeviceFirmwareConfigurationInterfaceSdCarddisabled      Windows10DeviceFirmwareConfigurationInterfaceSdCard = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceSdCardenabled       Windows10DeviceFirmwareConfigurationInterfaceSdCard = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceSdCardnotConfigured Windows10DeviceFirmwareConfigurationInterfaceSdCard = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceSdCard() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceSdCarddisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSdCardenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSdCardnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceSdCard(input string) (*Windows10DeviceFirmwareConfigurationInterfaceSdCard, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceSdCard{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceSdCarddisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceSdCardenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceSdCardnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceSdCard(input)
	return &out, nil
}
