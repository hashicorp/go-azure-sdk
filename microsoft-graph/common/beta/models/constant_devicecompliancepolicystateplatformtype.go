package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicyStatePlatformType string

const (
	DeviceCompliancePolicyStatePlatformTypeall                DeviceCompliancePolicyStatePlatformType = "All"
	DeviceCompliancePolicyStatePlatformTypeandroid            DeviceCompliancePolicyStatePlatformType = "Android"
	DeviceCompliancePolicyStatePlatformTypeandroidAOSP        DeviceCompliancePolicyStatePlatformType = "AndroidAOSP"
	DeviceCompliancePolicyStatePlatformTypeandroidForWork     DeviceCompliancePolicyStatePlatformType = "AndroidForWork"
	DeviceCompliancePolicyStatePlatformTypeandroidWorkProfile DeviceCompliancePolicyStatePlatformType = "AndroidWorkProfile"
	DeviceCompliancePolicyStatePlatformTypeiOS                DeviceCompliancePolicyStatePlatformType = "IOS"
	DeviceCompliancePolicyStatePlatformTypemacOS              DeviceCompliancePolicyStatePlatformType = "MacOS"
	DeviceCompliancePolicyStatePlatformTypewindows10AndLater  DeviceCompliancePolicyStatePlatformType = "Windows10AndLater"
	DeviceCompliancePolicyStatePlatformTypewindows10XProfile  DeviceCompliancePolicyStatePlatformType = "Windows10XProfile"
	DeviceCompliancePolicyStatePlatformTypewindows81AndLater  DeviceCompliancePolicyStatePlatformType = "Windows81AndLater"
	DeviceCompliancePolicyStatePlatformTypewindowsPhone81     DeviceCompliancePolicyStatePlatformType = "WindowsPhone81"
)

func PossibleValuesForDeviceCompliancePolicyStatePlatformType() []string {
	return []string{
		string(DeviceCompliancePolicyStatePlatformTypeall),
		string(DeviceCompliancePolicyStatePlatformTypeandroid),
		string(DeviceCompliancePolicyStatePlatformTypeandroidAOSP),
		string(DeviceCompliancePolicyStatePlatformTypeandroidForWork),
		string(DeviceCompliancePolicyStatePlatformTypeandroidWorkProfile),
		string(DeviceCompliancePolicyStatePlatformTypeiOS),
		string(DeviceCompliancePolicyStatePlatformTypemacOS),
		string(DeviceCompliancePolicyStatePlatformTypewindows10AndLater),
		string(DeviceCompliancePolicyStatePlatformTypewindows10XProfile),
		string(DeviceCompliancePolicyStatePlatformTypewindows81AndLater),
		string(DeviceCompliancePolicyStatePlatformTypewindowsPhone81),
	}
}

func parseDeviceCompliancePolicyStatePlatformType(input string) (*DeviceCompliancePolicyStatePlatformType, error) {
	vals := map[string]DeviceCompliancePolicyStatePlatformType{
		"all":                DeviceCompliancePolicyStatePlatformTypeall,
		"android":            DeviceCompliancePolicyStatePlatformTypeandroid,
		"androidaosp":        DeviceCompliancePolicyStatePlatformTypeandroidAOSP,
		"androidforwork":     DeviceCompliancePolicyStatePlatformTypeandroidForWork,
		"androidworkprofile": DeviceCompliancePolicyStatePlatformTypeandroidWorkProfile,
		"ios":                DeviceCompliancePolicyStatePlatformTypeiOS,
		"macos":              DeviceCompliancePolicyStatePlatformTypemacOS,
		"windows10andlater":  DeviceCompliancePolicyStatePlatformTypewindows10AndLater,
		"windows10xprofile":  DeviceCompliancePolicyStatePlatformTypewindows10XProfile,
		"windows81andlater":  DeviceCompliancePolicyStatePlatformTypewindows81AndLater,
		"windowsphone81":     DeviceCompliancePolicyStatePlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicyStatePlatformType(input)
	return &out, nil
}
