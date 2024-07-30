package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBookmarkPlatforms string

const (
	SearchBookmarkPlatforms_Android                            SearchBookmarkPlatforms = "android"
	SearchBookmarkPlatforms_AndroidAOSP                        SearchBookmarkPlatforms = "androidAOSP"
	SearchBookmarkPlatforms_AndroidForWork                     SearchBookmarkPlatforms = "androidForWork"
	SearchBookmarkPlatforms_AndroidMobileApplicationManagement SearchBookmarkPlatforms = "androidMobileApplicationManagement"
	SearchBookmarkPlatforms_AndroidWorkProfile                 SearchBookmarkPlatforms = "androidWorkProfile"
	SearchBookmarkPlatforms_IOS                                SearchBookmarkPlatforms = "iOS"
	SearchBookmarkPlatforms_IOSMobileApplicationManagement     SearchBookmarkPlatforms = "iOSMobileApplicationManagement"
	SearchBookmarkPlatforms_MacOS                              SearchBookmarkPlatforms = "macOS"
	SearchBookmarkPlatforms_Unknown                            SearchBookmarkPlatforms = "unknown"
	SearchBookmarkPlatforms_Windows10AndLater                  SearchBookmarkPlatforms = "windows10AndLater"
	SearchBookmarkPlatforms_Windows81AndLater                  SearchBookmarkPlatforms = "windows81AndLater"
	SearchBookmarkPlatforms_WindowsPhone81                     SearchBookmarkPlatforms = "windowsPhone81"
)

func PossibleValuesForSearchBookmarkPlatforms() []string {
	return []string{
		string(SearchBookmarkPlatforms_Android),
		string(SearchBookmarkPlatforms_AndroidAOSP),
		string(SearchBookmarkPlatforms_AndroidForWork),
		string(SearchBookmarkPlatforms_AndroidMobileApplicationManagement),
		string(SearchBookmarkPlatforms_AndroidWorkProfile),
		string(SearchBookmarkPlatforms_IOS),
		string(SearchBookmarkPlatforms_IOSMobileApplicationManagement),
		string(SearchBookmarkPlatforms_MacOS),
		string(SearchBookmarkPlatforms_Unknown),
		string(SearchBookmarkPlatforms_Windows10AndLater),
		string(SearchBookmarkPlatforms_Windows81AndLater),
		string(SearchBookmarkPlatforms_WindowsPhone81),
	}
}

func (s *SearchBookmarkPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchBookmarkPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchBookmarkPlatforms(input string) (*SearchBookmarkPlatforms, error) {
	vals := map[string]SearchBookmarkPlatforms{
		"android":                            SearchBookmarkPlatforms_Android,
		"androidaosp":                        SearchBookmarkPlatforms_AndroidAOSP,
		"androidforwork":                     SearchBookmarkPlatforms_AndroidForWork,
		"androidmobileapplicationmanagement": SearchBookmarkPlatforms_AndroidMobileApplicationManagement,
		"androidworkprofile":                 SearchBookmarkPlatforms_AndroidWorkProfile,
		"ios":                                SearchBookmarkPlatforms_IOS,
		"iosmobileapplicationmanagement":     SearchBookmarkPlatforms_IOSMobileApplicationManagement,
		"macos":                              SearchBookmarkPlatforms_MacOS,
		"unknown":                            SearchBookmarkPlatforms_Unknown,
		"windows10andlater":                  SearchBookmarkPlatforms_Windows10AndLater,
		"windows81andlater":                  SearchBookmarkPlatforms_Windows81AndLater,
		"windowsphone81":                     SearchBookmarkPlatforms_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchBookmarkPlatforms(input)
	return &out, nil
}
