package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform string

const (
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformandroid    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "Android"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformiOS        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "IOS"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformlinux      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "Linux"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformmacOS      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "MacOS"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformnone       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "None"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformwindows10  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "Windows10"
	DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformwindows10X DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform = "Windows10X"
)

func PossibleValuesForDeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform() []string {
	return []string{
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformandroid),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformiOS),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformlinux),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformmacOS),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformnone),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformwindows10),
		string(DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformwindows10X),
	}
}

func parseDeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform(input string) (*DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform, error) {
	vals := map[string]DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform{
		"android":    DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformandroid,
		"ios":        DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformiOS,
		"linux":      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformlinux,
		"macos":      DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformmacOS,
		"none":       DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformnone,
		"windows10":  DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformwindows10,
		"windows10x": DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatformwindows10X,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementConfigurationExchangeOnlineSettingApplicabilityPlatform(input)
	return &out, nil
}
