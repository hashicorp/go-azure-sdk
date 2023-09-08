package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia string

const (
	Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMediadisabled      Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMediaenabled       Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedianotConfigured Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMediadisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMediaenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedianotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia(input string) (*Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMediadisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMediaenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedianotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceBootFromExternalMedia(input)
	return &out, nil
}
