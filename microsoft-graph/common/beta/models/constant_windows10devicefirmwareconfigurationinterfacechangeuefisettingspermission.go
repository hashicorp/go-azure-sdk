package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission string

const (
	Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermissionnone              Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission = "None"
	Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermissionnotConfiguredOnly Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission = "NotConfiguredOnly"
)

func PossibleValuesForWindows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission() []string {
	return []string{
		string(Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermissionnone),
		string(Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermissionnotConfiguredOnly),
	}
}

func parseWindows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission(input string) (*Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission, error) {
	vals := map[string]Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission{
		"none":              Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermissionnone,
		"notconfiguredonly": Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermissionnotConfiguredOnly,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Windows10DeviceFirmwareConfigurationInterfaceChangeUefiSettingsPermission(input)
	return &out, nil
}
