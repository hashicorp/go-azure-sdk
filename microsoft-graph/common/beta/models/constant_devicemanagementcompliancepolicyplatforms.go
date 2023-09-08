package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementCompliancePolicyPlatforms string

const (
	DeviceManagementCompliancePolicyPlatformsandroid    DeviceManagementCompliancePolicyPlatforms = "Android"
	DeviceManagementCompliancePolicyPlatformsiOS        DeviceManagementCompliancePolicyPlatforms = "IOS"
	DeviceManagementCompliancePolicyPlatformslinux      DeviceManagementCompliancePolicyPlatforms = "Linux"
	DeviceManagementCompliancePolicyPlatformsmacOS      DeviceManagementCompliancePolicyPlatforms = "MacOS"
	DeviceManagementCompliancePolicyPlatformsnone       DeviceManagementCompliancePolicyPlatforms = "None"
	DeviceManagementCompliancePolicyPlatformswindows10  DeviceManagementCompliancePolicyPlatforms = "Windows10"
	DeviceManagementCompliancePolicyPlatformswindows10X DeviceManagementCompliancePolicyPlatforms = "Windows10X"
)

func PossibleValuesForDeviceManagementCompliancePolicyPlatforms() []string {
	return []string{
		string(DeviceManagementCompliancePolicyPlatformsandroid),
		string(DeviceManagementCompliancePolicyPlatformsiOS),
		string(DeviceManagementCompliancePolicyPlatformslinux),
		string(DeviceManagementCompliancePolicyPlatformsmacOS),
		string(DeviceManagementCompliancePolicyPlatformsnone),
		string(DeviceManagementCompliancePolicyPlatformswindows10),
		string(DeviceManagementCompliancePolicyPlatformswindows10X),
	}
}

func parseDeviceManagementCompliancePolicyPlatforms(input string) (*DeviceManagementCompliancePolicyPlatforms, error) {
	vals := map[string]DeviceManagementCompliancePolicyPlatforms{
		"android":    DeviceManagementCompliancePolicyPlatformsandroid,
		"ios":        DeviceManagementCompliancePolicyPlatformsiOS,
		"linux":      DeviceManagementCompliancePolicyPlatformslinux,
		"macos":      DeviceManagementCompliancePolicyPlatformsmacOS,
		"none":       DeviceManagementCompliancePolicyPlatformsnone,
		"windows10":  DeviceManagementCompliancePolicyPlatformswindows10,
		"windows10x": DeviceManagementCompliancePolicyPlatformswindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementCompliancePolicyPlatforms(input)
	return &out, nil
}
