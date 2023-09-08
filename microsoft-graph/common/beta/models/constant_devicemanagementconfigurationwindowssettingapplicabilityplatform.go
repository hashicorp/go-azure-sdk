package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationWindowsSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformandroid    DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "Android"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformiOS        DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "IOS"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformlinux      DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "Linux"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformmacOS      DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "MacOS"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformnone       DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "None"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformwindows10  DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "Windows10"
	DeviceManagementConfigurationWindowsSettingApplicabilityPlatformwindows10X DeviceManagementConfigurationWindowsSettingApplicabilityPlatform = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationWindowsSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformandroid),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformiOS),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformlinux),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformmacOS),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformnone),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformwindows10),
		string(DeviceManagementConfigurationWindowsSettingApplicabilityPlatformwindows10X),
	}
}

func parseDeviceManagementConfigurationWindowsSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationWindowsSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationWindowsSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationWindowsSettingApplicabilityPlatformandroid,
		"ios":        DeviceManagementConfigurationWindowsSettingApplicabilityPlatformiOS,
		"linux":      DeviceManagementConfigurationWindowsSettingApplicabilityPlatformlinux,
		"macos":      DeviceManagementConfigurationWindowsSettingApplicabilityPlatformmacOS,
		"none":       DeviceManagementConfigurationWindowsSettingApplicabilityPlatformnone,
		"windows10":  DeviceManagementConfigurationWindowsSettingApplicabilityPlatformwindows10,
		"windows10x": DeviceManagementConfigurationWindowsSettingApplicabilityPlatformwindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationWindowsSettingApplicabilityPlatform(input)
	return &out, nil
}
