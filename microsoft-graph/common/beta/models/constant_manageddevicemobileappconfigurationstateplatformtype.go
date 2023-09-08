package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfigurationStatePlatformType string

const (
	ManagedDeviceMobileAppConfigurationStatePlatformTypeall                ManagedDeviceMobileAppConfigurationStatePlatformType = "All"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeandroid            ManagedDeviceMobileAppConfigurationStatePlatformType = "Android"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidAOSP        ManagedDeviceMobileAppConfigurationStatePlatformType = "AndroidAOSP"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidForWork     ManagedDeviceMobileAppConfigurationStatePlatformType = "AndroidForWork"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidWorkProfile ManagedDeviceMobileAppConfigurationStatePlatformType = "AndroidWorkProfile"
	ManagedDeviceMobileAppConfigurationStatePlatformTypeiOS                ManagedDeviceMobileAppConfigurationStatePlatformType = "IOS"
	ManagedDeviceMobileAppConfigurationStatePlatformTypemacOS              ManagedDeviceMobileAppConfigurationStatePlatformType = "MacOS"
	ManagedDeviceMobileAppConfigurationStatePlatformTypewindows10AndLater  ManagedDeviceMobileAppConfigurationStatePlatformType = "Windows10AndLater"
	ManagedDeviceMobileAppConfigurationStatePlatformTypewindows10XProfile  ManagedDeviceMobileAppConfigurationStatePlatformType = "Windows10XProfile"
	ManagedDeviceMobileAppConfigurationStatePlatformTypewindows81AndLater  ManagedDeviceMobileAppConfigurationStatePlatformType = "Windows81AndLater"
	ManagedDeviceMobileAppConfigurationStatePlatformTypewindowsPhone81     ManagedDeviceMobileAppConfigurationStatePlatformType = "WindowsPhone81"
)

func PossibleValuesForManagedDeviceMobileAppConfigurationStatePlatformType() []string {
	return []string{
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeall),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeandroid),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidAOSP),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidForWork),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidWorkProfile),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypeiOS),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypemacOS),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypewindows10AndLater),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypewindows10XProfile),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypewindows81AndLater),
		string(ManagedDeviceMobileAppConfigurationStatePlatformTypewindowsPhone81),
	}
}

func parseManagedDeviceMobileAppConfigurationStatePlatformType(input string) (*ManagedDeviceMobileAppConfigurationStatePlatformType, error) {
	vals := map[string]ManagedDeviceMobileAppConfigurationStatePlatformType{
		"all":                ManagedDeviceMobileAppConfigurationStatePlatformTypeall,
		"android":            ManagedDeviceMobileAppConfigurationStatePlatformTypeandroid,
		"androidaosp":        ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidAOSP,
		"androidforwork":     ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidForWork,
		"androidworkprofile": ManagedDeviceMobileAppConfigurationStatePlatformTypeandroidWorkProfile,
		"ios":                ManagedDeviceMobileAppConfigurationStatePlatformTypeiOS,
		"macos":              ManagedDeviceMobileAppConfigurationStatePlatformTypemacOS,
		"windows10andlater":  ManagedDeviceMobileAppConfigurationStatePlatformTypewindows10AndLater,
		"windows10xprofile":  ManagedDeviceMobileAppConfigurationStatePlatformTypewindows10XProfile,
		"windows81andlater":  ManagedDeviceMobileAppConfigurationStatePlatformTypewindows81AndLater,
		"windowsphone81":     ManagedDeviceMobileAppConfigurationStatePlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceMobileAppConfigurationStatePlatformType(input)
	return &out, nil
}
