package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceBluetooth string

const (
	Windows10DeviceFirmwareConfigurationInterfaceBluetoothdisabled      Windows10DeviceFirmwareConfigurationInterfaceBluetooth = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceBluetoothenabled       Windows10DeviceFirmwareConfigurationInterfaceBluetooth = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceBluetoothnotConfigured Windows10DeviceFirmwareConfigurationInterfaceBluetooth = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceBluetooth() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceBluetoothdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBluetoothenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceBluetoothnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceBluetooth(input string) (*Windows10DeviceFirmwareConfigurationInterfaceBluetooth, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceBluetooth{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceBluetoothdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceBluetoothenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceBluetoothnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceBluetooth(input)
	return &out, nil
}
