package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceRadios string

const (
	Windows10DeviceFirmwareConfigurationInterfaceRadiosdisabled      Windows10DeviceFirmwareConfigurationInterfaceRadios = "Disabled"
	Windows10DeviceFirmwareConfigurationInterfaceRadiosenabled       Windows10DeviceFirmwareConfigurationInterfaceRadios = "Enabled"
	Windows10DeviceFirmwareConfigurationInterfaceRadiosnotConfigured Windows10DeviceFirmwareConfigurationInterfaceRadios = "NotConfigured"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceRadios() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceRadiosdisabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceRadiosenabled),
		string(Windows10DeviceFirmwareConfigurationInterfaceRadiosnotConfigured),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceRadios(input string) (*Windows10DeviceFirmwareConfigurationInterfaceRadios, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceRadios{
		"disabled":      Windows10DeviceFirmwareConfigurationInterfaceRadiosdisabled,
		"enabled":       Windows10DeviceFirmwareConfigurationInterfaceRadiosenabled,
		"notconfigured": Windows10DeviceFirmwareConfigurationInterfaceRadiosnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceRadios(input)
	return &out, nil
}
