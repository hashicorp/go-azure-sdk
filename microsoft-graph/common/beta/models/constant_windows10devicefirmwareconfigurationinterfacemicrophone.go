package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceMicrophone string

const (
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonedisabled      Windows10DeviceFirmwareConfigurationInterfaceMicrophone = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophoneenabled       Windows10DeviceFirmwareConfigurationInterfaceMicrophone = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceMicrophonenotConfigured Windows10DeviceFirmwareConfigurationInterfaceMicrophone = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceMicrophone() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonedisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophoneenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceMicrophonenotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceMicrophone(input string) (*Windows10DeviceFirmwareConfigurationInterfaceMicrophone, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceMicrophone{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceMicrophonedisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceMicrophoneenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceMicrophonenotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceMicrophone(input)
	return &out, nil
}
