package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationApplicationSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformandroid    DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "Android"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformiOS        DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "IOS"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformlinux      DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "Linux"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformmacOS      DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "MacOS"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformnone       DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "None"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformwindows10  DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "Windows10"
	DeviceManagementConfigurationApplicationSettingApplicabilityPlatformwindows10X DeviceManagementConfigurationApplicationSettingApplicabilityPlatform = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationApplicationSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformandroid),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformiOS),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformlinux),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformmacOS),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformnone),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformwindows10),
		string(DeviceManagementConfigurationApplicationSettingApplicabilityPlatformwindows10X),
	}
}

func parseDeviceManagementConfigurationApplicationSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationApplicationSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationApplicationSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationApplicationSettingApplicabilityPlatformandroid,
		"ios":        DeviceManagementConfigurationApplicationSettingApplicabilityPlatformiOS,
		"linux":      DeviceManagementConfigurationApplicationSettingApplicabilityPlatformlinux,
		"macos":      DeviceManagementConfigurationApplicationSettingApplicabilityPlatformmacOS,
		"none":       DeviceManagementConfigurationApplicationSettingApplicabilityPlatformnone,
		"windows10":  DeviceManagementConfigurationApplicationSettingApplicabilityPlatformwindows10,
		"windows10x": DeviceManagementConfigurationApplicationSettingApplicabilityPlatformwindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationApplicationSettingApplicabilityPlatform(input)
	return &out, nil
}
