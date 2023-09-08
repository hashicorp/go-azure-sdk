package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluateRequestPlatform string

const (
	AssignmentFilterEvaluateRequestPlatformandroid                            AssignmentFilterEvaluateRequestPlatform = "Android"
	AssignmentFilterEvaluateRequestPlatformandroidAOSP                        AssignmentFilterEvaluateRequestPlatform = "AndroidAOSP"
	AssignmentFilterEvaluateRequestPlatformandroidForWork                     AssignmentFilterEvaluateRequestPlatform = "AndroidForWork"
	AssignmentFilterEvaluateRequestPlatformandroidMobileApplicationManagement AssignmentFilterEvaluateRequestPlatform = "AndroidMobileApplicationManagement"
	AssignmentFilterEvaluateRequestPlatformandroidWorkProfile                 AssignmentFilterEvaluateRequestPlatform = "AndroidWorkProfile"
	AssignmentFilterEvaluateRequestPlatformiOS                                AssignmentFilterEvaluateRequestPlatform = "IOS"
	AssignmentFilterEvaluateRequestPlatformiOSMobileApplicationManagement     AssignmentFilterEvaluateRequestPlatform = "IOSMobileApplicationManagement"
	AssignmentFilterEvaluateRequestPlatformmacOS                              AssignmentFilterEvaluateRequestPlatform = "MacOS"
	AssignmentFilterEvaluateRequestPlatformunknown                            AssignmentFilterEvaluateRequestPlatform = "Unknown"
	AssignmentFilterEvaluateRequestPlatformwindows10AndLater                  AssignmentFilterEvaluateRequestPlatform = "Windows10AndLater"
	AssignmentFilterEvaluateRequestPlatformwindows81AndLater                  AssignmentFilterEvaluateRequestPlatform = "Windows81AndLater"
	AssignmentFilterEvaluateRequestPlatformwindowsPhone81                     AssignmentFilterEvaluateRequestPlatform = "WindowsPhone81"
)

func PossibleValuesForAssignmentFilterEvaluateRequestPlatform() []string {
	return []string{
		string(AssignmentFilterEvaluateRequestPlatformandroid),
		string(AssignmentFilterEvaluateRequestPlatformandroidAOSP),
		string(AssignmentFilterEvaluateRequestPlatformandroidForWork),
		string(AssignmentFilterEvaluateRequestPlatformandroidMobileApplicationManagement),
		string(AssignmentFilterEvaluateRequestPlatformandroidWorkProfile),
		string(AssignmentFilterEvaluateRequestPlatformiOS),
		string(AssignmentFilterEvaluateRequestPlatformiOSMobileApplicationManagement),
		string(AssignmentFilterEvaluateRequestPlatformmacOS),
		string(AssignmentFilterEvaluateRequestPlatformunknown),
		string(AssignmentFilterEvaluateRequestPlatformwindows10AndLater),
		string(AssignmentFilterEvaluateRequestPlatformwindows81AndLater),
		string(AssignmentFilterEvaluateRequestPlatformwindowsPhone81),
	}
}

func parseAssignmentFilterEvaluateRequestPlatform(input string) (*AssignmentFilterEvaluateRequestPlatform, error) {
	vals := map[string]AssignmentFilterEvaluateRequestPlatform{
		"android":                            AssignmentFilterEvaluateRequestPlatformandroid,
		"androidaosp":                        AssignmentFilterEvaluateRequestPlatformandroidAOSP,
		"androidforwork":                     AssignmentFilterEvaluateRequestPlatformandroidForWork,
		"androidmobileapplicationmanagement": AssignmentFilterEvaluateRequestPlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 AssignmentFilterEvaluateRequestPlatformandroidWorkProfile,
		"ios":                                AssignmentFilterEvaluateRequestPlatformiOS,
		"iosmobileapplicationmanagement":     AssignmentFilterEvaluateRequestPlatformiOSMobileApplicationManagement,
		"macos":                              AssignmentFilterEvaluateRequestPlatformmacOS,
		"unknown":                            AssignmentFilterEvaluateRequestPlatformunknown,
		"windows10andlater":                  AssignmentFilterEvaluateRequestPlatformwindows10AndLater,
		"windows81andlater":                  AssignmentFilterEvaluateRequestPlatformwindows81AndLater,
		"windowsphone81":                     AssignmentFilterEvaluateRequestPlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluateRequestPlatform(input)
	return &out, nil
}
