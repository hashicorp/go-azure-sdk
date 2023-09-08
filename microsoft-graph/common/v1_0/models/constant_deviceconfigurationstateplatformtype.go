package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceConfigurationStatePlatformType string

const (
	DeviceConfigurationStatePlatformTypeall               DeviceConfigurationStatePlatformType = "All"
	DeviceConfigurationStatePlatformTypeandroid           DeviceConfigurationStatePlatformType = "Android"
	DeviceConfigurationStatePlatformTypeandroidForWork    DeviceConfigurationStatePlatformType = "AndroidForWork"
	DeviceConfigurationStatePlatformTypeiOS               DeviceConfigurationStatePlatformType = "IOS"
	DeviceConfigurationStatePlatformTypemacOS             DeviceConfigurationStatePlatformType = "MacOS"
	DeviceConfigurationStatePlatformTypewindows10AndLater DeviceConfigurationStatePlatformType = "Windows10AndLater"
	DeviceConfigurationStatePlatformTypewindows81AndLater DeviceConfigurationStatePlatformType = "Windows81AndLater"
	DeviceConfigurationStatePlatformTypewindowsPhone81    DeviceConfigurationStatePlatformType = "WindowsPhone81"
)

func PossibleValuesForDeviceConfigurationStatePlatformType() []string {
	return []string{
		string(DeviceConfigurationStatePlatformTypeall),
		string(DeviceConfigurationStatePlatformTypeandroid),
		string(DeviceConfigurationStatePlatformTypeandroidForWork),
		string(DeviceConfigurationStatePlatformTypeiOS),
		string(DeviceConfigurationStatePlatformTypemacOS),
		string(DeviceConfigurationStatePlatformTypewindows10AndLater),
		string(DeviceConfigurationStatePlatformTypewindows81AndLater),
		string(DeviceConfigurationStatePlatformTypewindowsPhone81),
	}
}

func parseDeviceConfigurationStatePlatformType(input string) (*DeviceConfigurationStatePlatformType, error) {
	vals := map[string]DeviceConfigurationStatePlatformType{
		"all":               DeviceConfigurationStatePlatformTypeall,
		"android":           DeviceConfigurationStatePlatformTypeandroid,
		"androidforwork":    DeviceConfigurationStatePlatformTypeandroidForWork,
		"ios":               DeviceConfigurationStatePlatformTypeiOS,
		"macos":             DeviceConfigurationStatePlatformTypemacOS,
		"windows10andlater": DeviceConfigurationStatePlatformTypewindows10AndLater,
		"windows81andlater": DeviceConfigurationStatePlatformTypewindows81AndLater,
		"windowsphone81":    DeviceConfigurationStatePlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceConfigurationStatePlatformType(input)
	return &out, nil
}
