package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading string

const (
	Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingdisabled      Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingenabled       Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingnotConfigured Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading(input string) (*Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreadingnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceSimultaneousMultiThreading(input)
	return &out, nil
}
