package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication string

const (
	Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationdisabled      Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationenabled       Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationnotConfigured Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication(input string) (*Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunicationnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceNearFieldCommunication(input)
	return &out, nil
}
