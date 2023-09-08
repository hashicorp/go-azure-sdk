package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQnaPlatforms string

const (
	SearchQnaPlatformsandroid                            SearchQnaPlatforms = "Android"
	SearchQnaPlatformsandroidAOSP                        SearchQnaPlatforms = "AndroidAOSP"
	SearchQnaPlatformsandroidForWork                     SearchQnaPlatforms = "AndroidForWork"
	SearchQnaPlatformsandroidMobileApplicationManagement SearchQnaPlatforms = "AndroidMobileApplicationManagement"
	SearchQnaPlatformsandroidWorkProfile                 SearchQnaPlatforms = "AndroidWorkProfile"
	SearchQnaPlatformsiOS                                SearchQnaPlatforms = "IOS"
	SearchQnaPlatformsiOSMobileApplicationManagement     SearchQnaPlatforms = "IOSMobileApplicationManagement"
	SearchQnaPlatformsmacOS                              SearchQnaPlatforms = "MacOS"
	SearchQnaPlatformsunknown                            SearchQnaPlatforms = "Unknown"
	SearchQnaPlatformswindows10AndLater                  SearchQnaPlatforms = "Windows10AndLater"
	SearchQnaPlatformswindows81AndLater                  SearchQnaPlatforms = "Windows81AndLater"
	SearchQnaPlatformswindowsPhone81                     SearchQnaPlatforms = "WindowsPhone81"
)

func PossibleValuesForSearchQnaPlatforms() []string {
	return []string{
		string(SearchQnaPlatformsandroid),
		string(SearchQnaPlatformsandroidAOSP),
		string(SearchQnaPlatformsandroidForWork),
		string(SearchQnaPlatformsandroidMobileApplicationManagement),
		string(SearchQnaPlatformsandroidWorkProfile),
		string(SearchQnaPlatformsiOS),
		string(SearchQnaPlatformsiOSMobileApplicationManagement),
		string(SearchQnaPlatformsmacOS),
		string(SearchQnaPlatformsunknown),
		string(SearchQnaPlatformswindows10AndLater),
		string(SearchQnaPlatformswindows81AndLater),
		string(SearchQnaPlatformswindowsPhone81),
	}
}

func parseSearchQnaPlatforms(input string) (*SearchQnaPlatforms, error) {
	vals := map[string]SearchQnaPlatforms{
		"android":                            SearchQnaPlatformsandroid,
		"androidaosp":                        SearchQnaPlatformsandroidAOSP,
		"androidforwork":                     SearchQnaPlatformsandroidForWork,
		"androidmobileapplicationmanagement": SearchQnaPlatformsandroidMobileApplicationManagement,
		"androidworkprofile":                 SearchQnaPlatformsandroidWorkProfile,
		"ios":                                SearchQnaPlatformsiOS,
		"iosmobileapplicationmanagement":     SearchQnaPlatformsiOSMobileApplicationManagement,
		"macos":                              SearchQnaPlatformsmacOS,
		"unknown":                            SearchQnaPlatformsunknown,
		"windows10andlater":                  SearchQnaPlatformswindows10AndLater,
		"windows81andlater":                  SearchQnaPlatformswindows81AndLater,
		"windowsphone81":                     SearchQnaPlatformswindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchQnaPlatforms(input)
	return &out, nil
}
