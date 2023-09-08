package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityBaselineTemplatePlatformType string

const (
	SecurityBaselineTemplatePlatformTypeall                SecurityBaselineTemplatePlatformType = "All"
	SecurityBaselineTemplatePlatformTypeandroid            SecurityBaselineTemplatePlatformType = "Android"
	SecurityBaselineTemplatePlatformTypeandroidAOSP        SecurityBaselineTemplatePlatformType = "AndroidAOSP"
	SecurityBaselineTemplatePlatformTypeandroidForWork     SecurityBaselineTemplatePlatformType = "AndroidForWork"
	SecurityBaselineTemplatePlatformTypeandroidWorkProfile SecurityBaselineTemplatePlatformType = "AndroidWorkProfile"
	SecurityBaselineTemplatePlatformTypeiOS                SecurityBaselineTemplatePlatformType = "IOS"
	SecurityBaselineTemplatePlatformTypemacOS              SecurityBaselineTemplatePlatformType = "MacOS"
	SecurityBaselineTemplatePlatformTypewindows10AndLater  SecurityBaselineTemplatePlatformType = "Windows10AndLater"
	SecurityBaselineTemplatePlatformTypewindows10XProfile  SecurityBaselineTemplatePlatformType = "Windows10XProfile"
	SecurityBaselineTemplatePlatformTypewindows81AndLater  SecurityBaselineTemplatePlatformType = "Windows81AndLater"
	SecurityBaselineTemplatePlatformTypewindowsPhone81     SecurityBaselineTemplatePlatformType = "WindowsPhone81"
)

func PossibleValuesForSecurityBaselineTemplatePlatformType() []string {
	return []string{
		string(SecurityBaselineTemplatePlatformTypeall),
		string(SecurityBaselineTemplatePlatformTypeandroid),
		string(SecurityBaselineTemplatePlatformTypeandroidAOSP),
		string(SecurityBaselineTemplatePlatformTypeandroidForWork),
		string(SecurityBaselineTemplatePlatformTypeandroidWorkProfile),
		string(SecurityBaselineTemplatePlatformTypeiOS),
		string(SecurityBaselineTemplatePlatformTypemacOS),
		string(SecurityBaselineTemplatePlatformTypewindows10AndLater),
		string(SecurityBaselineTemplatePlatformTypewindows10XProfile),
		string(SecurityBaselineTemplatePlatformTypewindows81AndLater),
		string(SecurityBaselineTemplatePlatformTypewindowsPhone81),
	}
}

func parseSecurityBaselineTemplatePlatformType(input string) (*SecurityBaselineTemplatePlatformType, error) {
	vals := map[string]SecurityBaselineTemplatePlatformType{
		"all":                SecurityBaselineTemplatePlatformTypeall,
		"android":            SecurityBaselineTemplatePlatformTypeandroid,
		"androidaosp":        SecurityBaselineTemplatePlatformTypeandroidAOSP,
		"androidforwork":     SecurityBaselineTemplatePlatformTypeandroidForWork,
		"androidworkprofile": SecurityBaselineTemplatePlatformTypeandroidWorkProfile,
		"ios":                SecurityBaselineTemplatePlatformTypeiOS,
		"macos":              SecurityBaselineTemplatePlatformTypemacOS,
		"windows10andlater":  SecurityBaselineTemplatePlatformTypewindows10AndLater,
		"windows10xprofile":  SecurityBaselineTemplatePlatformTypewindows10XProfile,
		"windows81andlater":  SecurityBaselineTemplatePlatformTypewindows81AndLater,
		"windowsphone81":     SecurityBaselineTemplatePlatformTypewindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityBaselineTemplatePlatformType(input)
	return &out, nil
}
