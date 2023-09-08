package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBookmarkPlatforms string

const (
	SearchBookmarkPlatformsandroid                            SearchBookmarkPlatforms = "Android"
	SearchBookmarkPlatformsandroidAOSP                        SearchBookmarkPlatforms = "AndroidAOSP"
	SearchBookmarkPlatformsandroidForWork                     SearchBookmarkPlatforms = "AndroidForWork"
	SearchBookmarkPlatformsandroidMobileApplicationManagement SearchBookmarkPlatforms = "AndroidMobileApplicationManagement"
	SearchBookmarkPlatformsandroidWorkProfile                 SearchBookmarkPlatforms = "AndroidWorkProfile"
	SearchBookmarkPlatformsiOS                                SearchBookmarkPlatforms = "IOS"
	SearchBookmarkPlatformsiOSMobileApplicationManagement     SearchBookmarkPlatforms = "IOSMobileApplicationManagement"
	SearchBookmarkPlatformsmacOS                              SearchBookmarkPlatforms = "MacOS"
	SearchBookmarkPlatformsunknown                            SearchBookmarkPlatforms = "Unknown"
	SearchBookmarkPlatformswindows10AndLater                  SearchBookmarkPlatforms = "Windows10AndLater"
	SearchBookmarkPlatformswindows81AndLater                  SearchBookmarkPlatforms = "Windows81AndLater"
	SearchBookmarkPlatformswindowsPhone81                     SearchBookmarkPlatforms = "WindowsPhone81"
)

func PossibleValuesForSearchBookmarkPlatforms() []string {
	return []string{
		string(SearchBookmarkPlatformsandroid),
		string(SearchBookmarkPlatformsandroidAOSP),
		string(SearchBookmarkPlatformsandroidForWork),
		string(SearchBookmarkPlatformsandroidMobileApplicationManagement),
		string(SearchBookmarkPlatformsandroidWorkProfile),
		string(SearchBookmarkPlatformsiOS),
		string(SearchBookmarkPlatformsiOSMobileApplicationManagement),
		string(SearchBookmarkPlatformsmacOS),
		string(SearchBookmarkPlatformsunknown),
		string(SearchBookmarkPlatformswindows10AndLater),
		string(SearchBookmarkPlatformswindows81AndLater),
		string(SearchBookmarkPlatformswindowsPhone81),
	}
}

func parseSearchBookmarkPlatforms(input string) (*SearchBookmarkPlatforms, error) {
	vals := map[string]SearchBookmarkPlatforms{
		"android":                            SearchBookmarkPlatformsandroid,
		"androidaosp":                        SearchBookmarkPlatformsandroidAOSP,
		"androidforwork":                     SearchBookmarkPlatformsandroidForWork,
		"androidmobileapplicationmanagement": SearchBookmarkPlatformsandroidMobileApplicationManagement,
		"androidworkprofile":                 SearchBookmarkPlatformsandroidWorkProfile,
		"ios":                                SearchBookmarkPlatformsiOS,
		"iosmobileapplicationmanagement":     SearchBookmarkPlatformsiOSMobileApplicationManagement,
		"macos":                              SearchBookmarkPlatformsmacOS,
		"unknown":                            SearchBookmarkPlatformsunknown,
		"windows10andlater":                  SearchBookmarkPlatformswindows10AndLater,
		"windows81andlater":                  SearchBookmarkPlatformswindows81AndLater,
		"windowsphone81":                     SearchBookmarkPlatformswindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchBookmarkPlatforms(input)
	return &out, nil
}
