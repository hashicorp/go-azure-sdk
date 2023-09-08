package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters string

const (
	Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersdisabled      Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersenabled       Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersnotConfigured Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters(input string) (*Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdaptersnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceBootFromBuiltInNetworkAdapters(input)
	return &out, nil
}
