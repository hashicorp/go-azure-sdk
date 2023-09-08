package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationSettingApplicabilityPlatformandroid    DeviceManagementConfigurationSettingApplicabilityPlatform = "Android"
	DeviceManagementConfigurationSettingApplicabilityPlatformiOS        DeviceManagementConfigurationSettingApplicabilityPlatform = "IOS"
	DeviceManagementConfigurationSettingApplicabilityPlatformlinux      DeviceManagementConfigurationSettingApplicabilityPlatform = "Linux"
	DeviceManagementConfigurationSettingApplicabilityPlatformmacOS      DeviceManagementConfigurationSettingApplicabilityPlatform = "MacOS"
	DeviceManagementConfigurationSettingApplicabilityPlatformnone       DeviceManagementConfigurationSettingApplicabilityPlatform = "None"
	DeviceManagementConfigurationSettingApplicabilityPlatformwindows10  DeviceManagementConfigurationSettingApplicabilityPlatform = "Windows10"
	DeviceManagementConfigurationSettingApplicabilityPlatformwindows10X DeviceManagementConfigurationSettingApplicabilityPlatform = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationSettingApplicabilityPlatformandroid),
		string(DeviceManagementConfigurationSettingApplicabilityPlatformiOS),
		string(DeviceManagementConfigurationSettingApplicabilityPlatformlinux),
		string(DeviceManagementConfigurationSettingApplicabilityPlatformmacOS),
		string(DeviceManagementConfigurationSettingApplicabilityPlatformnone),
		string(DeviceManagementConfigurationSettingApplicabilityPlatformwindows10),
		string(DeviceManagementConfigurationSettingApplicabilityPlatformwindows10X),
	}
}

func parseDeviceManagementConfigurationSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationSettingApplicabilityPlatformandroid,
		"ios":        DeviceManagementConfigurationSettingApplicabilityPlatformiOS,
		"linux":      DeviceManagementConfigurationSettingApplicabilityPlatformlinux,
		"macos":      DeviceManagementConfigurationSettingApplicabilityPlatformmacOS,
		"none":       DeviceManagementConfigurationSettingApplicabilityPlatformnone,
		"windows10":  DeviceManagementConfigurationSettingApplicabilityPlatformwindows10,
		"windows10x": DeviceManagementConfigurationSettingApplicabilityPlatformwindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationSettingApplicabilityPlatform(input)
	return &out, nil
}
