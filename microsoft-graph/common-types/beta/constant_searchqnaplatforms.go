package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchQnaPlatforms string

const (
	SearchQnaPlatforms_Android                            SearchQnaPlatforms = "android"
	SearchQnaPlatforms_AndroidAOSP                        SearchQnaPlatforms = "androidAOSP"
	SearchQnaPlatforms_AndroidForWork                     SearchQnaPlatforms = "androidForWork"
	SearchQnaPlatforms_AndroidMobileApplicationManagement SearchQnaPlatforms = "androidMobileApplicationManagement"
	SearchQnaPlatforms_AndroidWorkProfile                 SearchQnaPlatforms = "androidWorkProfile"
	SearchQnaPlatforms_IOS                                SearchQnaPlatforms = "iOS"
	SearchQnaPlatforms_IOSMobileApplicationManagement     SearchQnaPlatforms = "iOSMobileApplicationManagement"
	SearchQnaPlatforms_MacOS                              SearchQnaPlatforms = "macOS"
	SearchQnaPlatforms_Unknown                            SearchQnaPlatforms = "unknown"
	SearchQnaPlatforms_Windows10AndLater                  SearchQnaPlatforms = "windows10AndLater"
	SearchQnaPlatforms_Windows81AndLater                  SearchQnaPlatforms = "windows81AndLater"
	SearchQnaPlatforms_WindowsPhone81                     SearchQnaPlatforms = "windowsPhone81"
)

func PossibleValuesForSearchQnaPlatforms() []string {
	return []string{
		string(SearchQnaPlatforms_Android),
		string(SearchQnaPlatforms_AndroidAOSP),
		string(SearchQnaPlatforms_AndroidForWork),
		string(SearchQnaPlatforms_AndroidMobileApplicationManagement),
		string(SearchQnaPlatforms_AndroidWorkProfile),
		string(SearchQnaPlatforms_IOS),
		string(SearchQnaPlatforms_IOSMobileApplicationManagement),
		string(SearchQnaPlatforms_MacOS),
		string(SearchQnaPlatforms_Unknown),
		string(SearchQnaPlatforms_Windows10AndLater),
		string(SearchQnaPlatforms_Windows81AndLater),
		string(SearchQnaPlatforms_WindowsPhone81),
	}
}

func (s *SearchQnaPlatforms) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchQnaPlatforms(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchQnaPlatforms(input string) (*SearchQnaPlatforms, error) {
	vals := map[string]SearchQnaPlatforms{
		"android":                            SearchQnaPlatforms_Android,
		"androidaosp":                        SearchQnaPlatforms_AndroidAOSP,
		"androidforwork":                     SearchQnaPlatforms_AndroidForWork,
		"androidmobileapplicationmanagement": SearchQnaPlatforms_AndroidMobileApplicationManagement,
		"androidworkprofile":                 SearchQnaPlatforms_AndroidWorkProfile,
		"ios":                                SearchQnaPlatforms_IOS,
		"iosmobileapplicationmanagement":     SearchQnaPlatforms_IOSMobileApplicationManagement,
		"macos":                              SearchQnaPlatforms_MacOS,
		"unknown":                            SearchQnaPlatforms_Unknown,
		"windows10andlater":                  SearchQnaPlatforms_Windows10AndLater,
		"windows81andlater":                  SearchQnaPlatforms_Windows81AndLater,
		"windowsphone81":                     SearchQnaPlatforms_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchQnaPlatforms(input)
	return &out, nil
}
