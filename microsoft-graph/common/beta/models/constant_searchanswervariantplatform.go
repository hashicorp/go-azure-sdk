package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerVariantPlatform string

const (
	SearchAnswerVariantPlatformandroid                            SearchAnswerVariantPlatform = "Android"
	SearchAnswerVariantPlatformandroidAOSP                        SearchAnswerVariantPlatform = "AndroidAOSP"
	SearchAnswerVariantPlatformandroidForWork                     SearchAnswerVariantPlatform = "AndroidForWork"
	SearchAnswerVariantPlatformandroidMobileApplicationManagement SearchAnswerVariantPlatform = "AndroidMobileApplicationManagement"
	SearchAnswerVariantPlatformandroidWorkProfile                 SearchAnswerVariantPlatform = "AndroidWorkProfile"
	SearchAnswerVariantPlatformiOS                                SearchAnswerVariantPlatform = "IOS"
	SearchAnswerVariantPlatformiOSMobileApplicationManagement     SearchAnswerVariantPlatform = "IOSMobileApplicationManagement"
	SearchAnswerVariantPlatformmacOS                              SearchAnswerVariantPlatform = "MacOS"
	SearchAnswerVariantPlatformunknown                            SearchAnswerVariantPlatform = "Unknown"
	SearchAnswerVariantPlatformwindows10AndLater                  SearchAnswerVariantPlatform = "Windows10AndLater"
	SearchAnswerVariantPlatformwindows81AndLater                  SearchAnswerVariantPlatform = "Windows81AndLater"
	SearchAnswerVariantPlatformwindowsPhone81                     SearchAnswerVariantPlatform = "WindowsPhone81"
)

func PossibleValuesForSearchAnswerVariantPlatform() []string {
	return []string{
		string(SearchAnswerVariantPlatformandroid),
		string(SearchAnswerVariantPlatformandroidAOSP),
		string(SearchAnswerVariantPlatformandroidForWork),
		string(SearchAnswerVariantPlatformandroidMobileApplicationManagement),
		string(SearchAnswerVariantPlatformandroidWorkProfile),
		string(SearchAnswerVariantPlatformiOS),
		string(SearchAnswerVariantPlatformiOSMobileApplicationManagement),
		string(SearchAnswerVariantPlatformmacOS),
		string(SearchAnswerVariantPlatformunknown),
		string(SearchAnswerVariantPlatformwindows10AndLater),
		string(SearchAnswerVariantPlatformwindows81AndLater),
		string(SearchAnswerVariantPlatformwindowsPhone81),
	}
}

func parseSearchAnswerVariantPlatform(input string) (*SearchAnswerVariantPlatform, error) {
	vals := map[string]SearchAnswerVariantPlatform{
		"android":                            SearchAnswerVariantPlatformandroid,
		"androidaosp":                        SearchAnswerVariantPlatformandroidAOSP,
		"androidforwork":                     SearchAnswerVariantPlatformandroidForWork,
		"androidmobileapplicationmanagement": SearchAnswerVariantPlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 SearchAnswerVariantPlatformandroidWorkProfile,
		"ios":                                SearchAnswerVariantPlatformiOS,
		"iosmobileapplicationmanagement":     SearchAnswerVariantPlatformiOSMobileApplicationManagement,
		"macos":                              SearchAnswerVariantPlatformmacOS,
		"unknown":                            SearchAnswerVariantPlatformunknown,
		"windows10andlater":                  SearchAnswerVariantPlatformwindows10AndLater,
		"windows81andlater":                  SearchAnswerVariantPlatformwindows81AndLater,
		"windowsphone81":                     SearchAnswerVariantPlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchAnswerVariantPlatform(input)
	return &out, nil
}
