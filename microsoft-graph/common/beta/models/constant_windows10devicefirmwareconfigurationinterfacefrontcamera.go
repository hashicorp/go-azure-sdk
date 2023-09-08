package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceFrontCamera string

const (
	Windows10DeviceFirmwareConfigurationInterfaceFrontCameradisabled      Windows10DeviceFirmwareConfigurationInterfaceFrontCamera = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceFrontCameraenabled       Windows10DeviceFirmwareConfigurationInterfaceFrontCamera = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceFrontCameranotConfigured Windows10DeviceFirmwareConfigurationInterfaceFrontCamera = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceFrontCamera() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceFrontCameradisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceFrontCameraenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceFrontCameranotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceFrontCamera(input string) (*Windows10DeviceFirmwareConfigurationInterfaceFrontCamera, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceFrontCamera{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceFrontCameradisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceFrontCameraenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceFrontCameranotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceFrontCamera(input)
	return &out, nil
}
