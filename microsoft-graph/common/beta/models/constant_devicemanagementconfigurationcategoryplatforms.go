package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationCategoryPlatforms string

const (
	DeviceManagementConfigurationCategoryPlatformsandroid    DeviceManagementConfigurationCategoryPlatforms = "Android"
	DeviceManagementConfigurationCategoryPlatformsiOS        DeviceManagementConfigurationCategoryPlatforms = "IOS"
	DeviceManagementConfigurationCategoryPlatformslinux      DeviceManagementConfigurationCategoryPlatforms = "Linux"
	DeviceManagementConfigurationCategoryPlatformsmacOS      DeviceManagementConfigurationCategoryPlatforms = "MacOS"
	DeviceManagementConfigurationCategoryPlatformsnone       DeviceManagementConfigurationCategoryPlatforms = "None"
	DeviceManagementConfigurationCategoryPlatformswindows10  DeviceManagementConfigurationCategoryPlatforms = "Windows10"
	DeviceManagementConfigurationCategoryPlatformswindows10X DeviceManagementConfigurationCategoryPlatforms = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationCategoryPlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationCategoryPlatformsandroid),
		string(DeviceManagementConfigurationCategoryPlatformsiOS),
		string(DeviceManagementConfigurationCategoryPlatformslinux),
		string(DeviceManagementConfigurationCategoryPlatformsmacOS),
		string(DeviceManagementConfigurationCategoryPlatformsnone),
		string(DeviceManagementConfigurationCategoryPlatformswindows10),
		string(DeviceManagementConfigurationCategoryPlatformswindows10X),
	}
}

func parseDeviceManagementConfigurationCategoryPlatforms(input string) (*DeviceManagementConfigurationCategoryPlatforms, error) {
	vals := map[string]DeviceManagementConfigurationCategoryPlatforms{
		"android":    DeviceManagementConfigurationCategoryPlatformsandroid,
		"ios":        DeviceManagementConfigurationCategoryPlatformsiOS,
		"linux":      DeviceManagementConfigurationCategoryPlatformslinux,
		"macos":      DeviceManagementConfigurationCategoryPlatformsmacOS,
		"none":       DeviceManagementConfigurationCategoryPlatformsnone,
		"windows10":  DeviceManagementConfigurationCategoryPlatformswindows10,
		"windows10x": DeviceManagementConfigurationCategoryPlatformswindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationCategoryPlatforms(input)
	return &out, nil
}
