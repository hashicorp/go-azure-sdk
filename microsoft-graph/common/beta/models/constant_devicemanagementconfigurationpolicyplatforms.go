package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyPlatforms string

const (
	DeviceManagementConfigurationPolicyPlatformsandroid    DeviceManagementConfigurationPolicyPlatforms = "Android"
	DeviceManagementConfigurationPolicyPlatformsiOS        DeviceManagementConfigurationPolicyPlatforms = "IOS"
	DeviceManagementConfigurationPolicyPlatformslinux      DeviceManagementConfigurationPolicyPlatforms = "Linux"
	DeviceManagementConfigurationPolicyPlatformsmacOS      DeviceManagementConfigurationPolicyPlatforms = "MacOS"
	DeviceManagementConfigurationPolicyPlatformsnone       DeviceManagementConfigurationPolicyPlatforms = "None"
	DeviceManagementConfigurationPolicyPlatformswindows10  DeviceManagementConfigurationPolicyPlatforms = "Windows10"
	DeviceManagementConfigurationPolicyPlatformswindows10X DeviceManagementConfigurationPolicyPlatforms = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationPolicyPlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyPlatformsandroid),
		string(DeviceManagementConfigurationPolicyPlatformsiOS),
		string(DeviceManagementConfigurationPolicyPlatformslinux),
		string(DeviceManagementConfigurationPolicyPlatformsmacOS),
		string(DeviceManagementConfigurationPolicyPlatformsnone),
		string(DeviceManagementConfigurationPolicyPlatformswindows10),
		string(DeviceManagementConfigurationPolicyPlatformswindows10X),
	}
}

func parseDeviceManagementConfigurationPolicyPlatforms(input string) (*DeviceManagementConfigurationPolicyPlatforms, error) {
	vals := map[string]DeviceManagementConfigurationPolicyPlatforms{
		"android":    DeviceManagementConfigurationPolicyPlatformsandroid,
		"ios":        DeviceManagementConfigurationPolicyPlatformsiOS,
		"linux":      DeviceManagementConfigurationPolicyPlatformslinux,
		"macos":      DeviceManagementConfigurationPolicyPlatformsmacOS,
		"none":       DeviceManagementConfigurationPolicyPlatformsnone,
		"windows10":  DeviceManagementConfigurationPolicyPlatformswindows10,
		"windows10x": DeviceManagementConfigurationPolicyPlatformswindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyPlatforms(input)
	return &out, nil
}
