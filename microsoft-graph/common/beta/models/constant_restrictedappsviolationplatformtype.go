package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsViolationPlatformType string

const (
	RestrictedAppsViolationPlatformTypeall                RestrictedAppsViolationPlatformType = "All"
	RestrictedAppsViolationPlatformTypeandroid            RestrictedAppsViolationPlatformType = "Android"
	RestrictedAppsViolationPlatformTypeandroidAOSP        RestrictedAppsViolationPlatformType = "AndroidAOSP"
	RestrictedAppsViolationPlatformTypeandroidForWork     RestrictedAppsViolationPlatformType = "AndroidForWork"
	RestrictedAppsViolationPlatformTypeandroidWorkProfile RestrictedAppsViolationPlatformType = "AndroidWorkProfile"
	RestrictedAppsViolationPlatformTypeiOS                RestrictedAppsViolationPlatformType = "IOS"
	RestrictedAppsViolationPlatformTypemacOS              RestrictedAppsViolationPlatformType = "MacOS"
	RestrictedAppsViolationPlatformTypewindows10AndLater  RestrictedAppsViolationPlatformType = "Windows10AndLater"
	RestrictedAppsViolationPlatformTypewindows10XProfile  RestrictedAppsViolationPlatformType = "Windows10XProfile"
	RestrictedAppsViolationPlatformTypewindows81AndLater  RestrictedAppsViolationPlatformType = "Windows81AndLater"
	RestrictedAppsViolationPlatformTypewindowsPhone81     RestrictedAppsViolationPlatformType = "WindowsPhone81"
)

func PossibleValuesForRestrictedAppsViolationPlatformType() []string {
	return []string{
		string(RestrictedAppsViolationPlatformTypeall),
		string(RestrictedAppsViolationPlatformTypeandroid),
		string(RestrictedAppsViolationPlatformTypeandroidAOSP),
		string(RestrictedAppsViolationPlatformTypeandroidForWork),
		string(RestrictedAppsViolationPlatformTypeandroidWorkProfile),
		string(RestrictedAppsViolationPlatformTypeiOS),
		string(RestrictedAppsViolationPlatformTypemacOS),
		string(RestrictedAppsViolationPlatformTypewindows10AndLater),
		string(RestrictedAppsViolationPlatformTypewindows10XProfile),
		string(RestrictedAppsViolationPlatformTypewindows81AndLater),
		string(RestrictedAppsViolationPlatformTypewindowsPhone81),
	}
}

func parseRestrictedAppsViolationPlatformType(input string) (*RestrictedAppsViolationPlatformType, error) {
	vals := map[string]RestrictedAppsViolationPlatformType{
		"all":                RestrictedAppsViolationPlatformTypeall,
		"android":            RestrictedAppsViolationPlatformTypeandroid,
		"androidaosp":        RestrictedAppsViolationPlatformTypeandroidAOSP,
		"androidforwork":     RestrictedAppsViolationPlatformTypeandroidForWork,
		"androidworkprofile": RestrictedAppsViolationPlatformTypeandroidWorkProfile,
		"ios":                RestrictedAppsViolationPlatformTypeiOS,
		"macos":              RestrictedAppsViolationPlatformTypemacOS,
		"windows10andlater":  RestrictedAppsViolationPlatformTypewindows10AndLater,
		"windows10xprofile":  RestrictedAppsViolationPlatformTypewindows10XProfile,
		"windows81andlater":  RestrictedAppsViolationPlatformTypewindows81AndLater,
		"windowsphone81":     RestrictedAppsViolationPlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictedAppsViolationPlatformType(input)
	return &out, nil
}
