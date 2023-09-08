package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementTemplatePlatformType string

const (
	DeviceManagementTemplatePlatformTypeall                DeviceManagementTemplatePlatformType = "All"
	DeviceManagementTemplatePlatformTypeandroid            DeviceManagementTemplatePlatformType = "Android"
	DeviceManagementTemplatePlatformTypeandroidAOSP        DeviceManagementTemplatePlatformType = "AndroidAOSP"
	DeviceManagementTemplatePlatformTypeandroidForWork     DeviceManagementTemplatePlatformType = "AndroidForWork"
	DeviceManagementTemplatePlatformTypeandroidWorkProfile DeviceManagementTemplatePlatformType = "AndroidWorkProfile"
	DeviceManagementTemplatePlatformTypeiOS                DeviceManagementTemplatePlatformType = "IOS"
	DeviceManagementTemplatePlatformTypemacOS              DeviceManagementTemplatePlatformType = "MacOS"
	DeviceManagementTemplatePlatformTypewindows10AndLater  DeviceManagementTemplatePlatformType = "Windows10AndLater"
	DeviceManagementTemplatePlatformTypewindows10XProfile  DeviceManagementTemplatePlatformType = "Windows10XProfile"
	DeviceManagementTemplatePlatformTypewindows81AndLater  DeviceManagementTemplatePlatformType = "Windows81AndLater"
	DeviceManagementTemplatePlatformTypewindowsPhone81     DeviceManagementTemplatePlatformType = "WindowsPhone81"
)

func PossibleValuesForDeviceManagementTemplatePlatformType() []string {
	return []string{
		string(DeviceManagementTemplatePlatformTypeall),
		string(DeviceManagementTemplatePlatformTypeandroid),
		string(DeviceManagementTemplatePlatformTypeandroidAOSP),
		string(DeviceManagementTemplatePlatformTypeandroidForWork),
		string(DeviceManagementTemplatePlatformTypeandroidWorkProfile),
		string(DeviceManagementTemplatePlatformTypeiOS),
		string(DeviceManagementTemplatePlatformTypemacOS),
		string(DeviceManagementTemplatePlatformTypewindows10AndLater),
		string(DeviceManagementTemplatePlatformTypewindows10XProfile),
		string(DeviceManagementTemplatePlatformTypewindows81AndLater),
		string(DeviceManagementTemplatePlatformTypewindowsPhone81),
	}
}

func parseDeviceManagementTemplatePlatformType(input string) (*DeviceManagementTemplatePlatformType, error) {
	vals := map[string]DeviceManagementTemplatePlatformType{
		"all":                DeviceManagementTemplatePlatformTypeall,
		"android":            DeviceManagementTemplatePlatformTypeandroid,
		"androidaosp":        DeviceManagementTemplatePlatformTypeandroidAOSP,
		"androidforwork":     DeviceManagementTemplatePlatformTypeandroidForWork,
		"androidworkprofile": DeviceManagementTemplatePlatformTypeandroidWorkProfile,
		"ios":                DeviceManagementTemplatePlatformTypeiOS,
		"macos":              DeviceManagementTemplatePlatformTypemacOS,
		"windows10andlater":  DeviceManagementTemplatePlatformTypewindows10AndLater,
		"windows10xprofile":  DeviceManagementTemplatePlatformTypewindows10XProfile,
		"windows81andlater":  DeviceManagementTemplatePlatformTypewindows81AndLater,
		"windowsphone81":     DeviceManagementTemplatePlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceManagementTemplatePlatformType(input)
	return &out, nil
}
