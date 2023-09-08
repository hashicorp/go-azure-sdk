package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable string

const (
	Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTabledisabled      Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTableenabled       Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTablenotConfigured Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTabledisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTableenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTablenotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable(input string) (*Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTabledisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTableenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTablenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceWindowsPlatformBinaryTable(input)
	return &out, nil
}
