package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryAssignmentFilterPlatform string

const (
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroid                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "Android"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidAOSP                        AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "AndroidAOSP"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidForWork                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "AndroidForWork"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidMobileApplicationManagement AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "AndroidMobileApplicationManagement"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidWorkProfile                 AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "AndroidWorkProfile"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformiOS                                AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "IOS"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformiOSMobileApplicationManagement     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "IOSMobileApplicationManagement"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformmacOS                              AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "MacOS"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformunknown                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "Unknown"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindows10AndLater                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "Windows10AndLater"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindows81AndLater                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "Windows81AndLater"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindowsPhone81                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "WindowsPhone81"
)

func PossibleValuesForAssignmentFilterEvaluationSummaryAssignmentFilterPlatform() []string {
	return []string{
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroid),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidAOSP),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidForWork),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidMobileApplicationManagement),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidWorkProfile),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformiOS),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformiOSMobileApplicationManagement),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformmacOS),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformunknown),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindows10AndLater),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindows81AndLater),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindowsPhone81),
	}
}

func parseAssignmentFilterEvaluationSummaryAssignmentFilterPlatform(input string) (*AssignmentFilterEvaluationSummaryAssignmentFilterPlatform, error) {
	vals := map[string]AssignmentFilterEvaluationSummaryAssignmentFilterPlatform{
		"android":                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroid,
		"androidaosp":                        AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidAOSP,
		"androidforwork":                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidForWork,
		"androidmobileapplicationmanagement": AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 AssignmentFilterEvaluationSummaryAssignmentFilterPlatformandroidWorkProfile,
		"ios":                                AssignmentFilterEvaluationSummaryAssignmentFilterPlatformiOS,
		"iosmobileapplicationmanagement":     AssignmentFilterEvaluationSummaryAssignmentFilterPlatformiOSMobileApplicationManagement,
		"macos":                              AssignmentFilterEvaluationSummaryAssignmentFilterPlatformmacOS,
		"unknown":                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatformunknown,
		"windows10andlater":                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindows10AndLater,
		"windows81andlater":                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindows81AndLater,
		"windowsphone81":                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationSummaryAssignmentFilterPlatform(input)
	return &out, nil
}
