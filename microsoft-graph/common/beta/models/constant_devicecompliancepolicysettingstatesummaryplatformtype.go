package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceCompliancePolicySettingStateSummaryPlatformType string

const (
	DeviceCompliancePolicySettingStateSummaryPlatformTypeall                DeviceCompliancePolicySettingStateSummaryPlatformType = "All"
	DeviceCompliancePolicySettingStateSummaryPlatformTypeandroid            DeviceCompliancePolicySettingStateSummaryPlatformType = "Android"
	DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidAOSP        DeviceCompliancePolicySettingStateSummaryPlatformType = "AndroidAOSP"
	DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidForWork     DeviceCompliancePolicySettingStateSummaryPlatformType = "AndroidForWork"
	DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidWorkProfile DeviceCompliancePolicySettingStateSummaryPlatformType = "AndroidWorkProfile"
	DeviceCompliancePolicySettingStateSummaryPlatformTypeiOS                DeviceCompliancePolicySettingStateSummaryPlatformType = "IOS"
	DeviceCompliancePolicySettingStateSummaryPlatformTypemacOS              DeviceCompliancePolicySettingStateSummaryPlatformType = "MacOS"
	DeviceCompliancePolicySettingStateSummaryPlatformTypewindows10AndLater  DeviceCompliancePolicySettingStateSummaryPlatformType = "Windows10AndLater"
	DeviceCompliancePolicySettingStateSummaryPlatformTypewindows10XProfile  DeviceCompliancePolicySettingStateSummaryPlatformType = "Windows10XProfile"
	DeviceCompliancePolicySettingStateSummaryPlatformTypewindows81AndLater  DeviceCompliancePolicySettingStateSummaryPlatformType = "Windows81AndLater"
	DeviceCompliancePolicySettingStateSummaryPlatformTypewindowsPhone81     DeviceCompliancePolicySettingStateSummaryPlatformType = "WindowsPhone81"
)

func PossibleValuesForDeviceCompliancePolicySettingStateSummaryPlatformType() []string {
	return []string{
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypeall),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypeandroid),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidAOSP),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidForWork),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidWorkProfile),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypeiOS),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypemacOS),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypewindows10AndLater),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypewindows10XProfile),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypewindows81AndLater),
		string(DeviceCompliancePolicySettingStateSummaryPlatformTypewindowsPhone81),
	}
}

func parseDeviceCompliancePolicySettingStateSummaryPlatformType(input string) (*DeviceCompliancePolicySettingStateSummaryPlatformType, error) {
	vals := map[string]DeviceCompliancePolicySettingStateSummaryPlatformType{
		"all":                DeviceCompliancePolicySettingStateSummaryPlatformTypeall,
		"android":            DeviceCompliancePolicySettingStateSummaryPlatformTypeandroid,
		"androidaosp":        DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidAOSP,
		"androidforwork":     DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidForWork,
		"androidworkprofile": DeviceCompliancePolicySettingStateSummaryPlatformTypeandroidWorkProfile,
		"ios":                DeviceCompliancePolicySettingStateSummaryPlatformTypeiOS,
		"macos":              DeviceCompliancePolicySettingStateSummaryPlatformTypemacOS,
		"windows10andlater":  DeviceCompliancePolicySettingStateSummaryPlatformTypewindows10AndLater,
		"windows10xprofile":  DeviceCompliancePolicySettingStateSummaryPlatformTypewindows10XProfile,
		"windows81andlater":  DeviceCompliancePolicySettingStateSummaryPlatformTypewindows81AndLater,
		"windowsphone81":     DeviceCompliancePolicySettingStateSummaryPlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DeviceCompliancePolicySettingStateSummaryPlatformType(input)
	return &out, nil
}
