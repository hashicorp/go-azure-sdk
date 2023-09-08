package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationPolicyTemplatePlatforms string

const (
	DeviceManagementConfigurationPolicyTemplatePlatformsandroid    DeviceManagementConfigurationPolicyTemplatePlatforms = "Android"
	DeviceManagementConfigurationPolicyTemplatePlatformsiOS        DeviceManagementConfigurationPolicyTemplatePlatforms = "IOS"
	DeviceManagementConfigurationPolicyTemplatePlatformslinux      DeviceManagementConfigurationPolicyTemplatePlatforms = "Linux"
	DeviceManagementConfigurationPolicyTemplatePlatformsmacOS      DeviceManagementConfigurationPolicyTemplatePlatforms = "MacOS"
	DeviceManagementConfigurationPolicyTemplatePlatformsnone       DeviceManagementConfigurationPolicyTemplatePlatforms = "None"
	DeviceManagementConfigurationPolicyTemplatePlatformswindows10  DeviceManagementConfigurationPolicyTemplatePlatforms = "Windows10"
	DeviceManagementConfigurationPolicyTemplatePlatformswindows10X DeviceManagementConfigurationPolicyTemplatePlatforms = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationPolicyTemplatePlatforms() []string {
	return []string{
		string(DeviceManagementConfigurationPolicyTemplatePlatformsandroid),
		string(DeviceManagementConfigurationPolicyTemplatePlatformsiOS),
		string(DeviceManagementConfigurationPolicyTemplatePlatformslinux),
		string(DeviceManagementConfigurationPolicyTemplatePlatformsmacOS),
		string(DeviceManagementConfigurationPolicyTemplatePlatformsnone),
		string(DeviceManagementConfigurationPolicyTemplatePlatformswindows10),
		string(DeviceManagementConfigurationPolicyTemplatePlatformswindows10X),
	}
}

func parseDeviceManagementConfigurationPolicyTemplatePlatforms(input string) (*DeviceManagementConfigurationPolicyTemplatePlatforms, error) {
	vals := map[string]DeviceManagementConfigurationPolicyTemplatePlatforms{
		"android":    DeviceManagementConfigurationPolicyTemplatePlatformsandroid,
		"ios":        DeviceManagementConfigurationPolicyTemplatePlatformsiOS,
		"linux":      DeviceManagementConfigurationPolicyTemplatePlatformslinux,
		"macos":      DeviceManagementConfigurationPolicyTemplatePlatformsmacOS,
		"none":       DeviceManagementConfigurationPolicyTemplatePlatformsnone,
		"windows10":  DeviceManagementConfigurationPolicyTemplatePlatformswindows10,
		"windows10x": DeviceManagementConfigurationPolicyTemplatePlatformswindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationPolicyTemplatePlatforms(input)
	return &out, nil
}
