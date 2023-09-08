package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceCameras string

const (
	Windows10DeviceFirmwareConfigurationInterfaceCamerasdisabled      Windows10DeviceFirmwareConfigurationInterfaceCameras = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceCamerasenabled       Windows10DeviceFirmwareConfigurationInterfaceCameras = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceCamerasnotConfigured Windows10DeviceFirmwareConfigurationInterfaceCameras = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceCameras() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceCamerasdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceCamerasenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceCamerasnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceCameras(input string) (*Windows10DeviceFirmwareConfigurationInterfaceCameras, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceCameras{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceCamerasdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceCamerasenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceCamerasnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceCameras(input)
	return &out, nil
}
