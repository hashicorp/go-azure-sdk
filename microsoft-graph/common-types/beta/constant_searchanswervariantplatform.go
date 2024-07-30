package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchAnswerVariantPlatform string

const (
	SearchAnswerVariantPlatform_Android                            SearchAnswerVariantPlatform = "android"
	SearchAnswerVariantPlatform_AndroidAOSP                        SearchAnswerVariantPlatform = "androidAOSP"
	SearchAnswerVariantPlatform_AndroidForWork                     SearchAnswerVariantPlatform = "androidForWork"
	SearchAnswerVariantPlatform_AndroidMobileApplicationManagement SearchAnswerVariantPlatform = "androidMobileApplicationManagement"
	SearchAnswerVariantPlatform_AndroidWorkProfile                 SearchAnswerVariantPlatform = "androidWorkProfile"
	SearchAnswerVariantPlatform_IOS                                SearchAnswerVariantPlatform = "iOS"
	SearchAnswerVariantPlatform_IOSMobileApplicationManagement     SearchAnswerVariantPlatform = "iOSMobileApplicationManagement"
	SearchAnswerVariantPlatform_MacOS                              SearchAnswerVariantPlatform = "macOS"
	SearchAnswerVariantPlatform_Unknown                            SearchAnswerVariantPlatform = "unknown"
	SearchAnswerVariantPlatform_Windows10AndLater                  SearchAnswerVariantPlatform = "windows10AndLater"
	SearchAnswerVariantPlatform_Windows81AndLater                  SearchAnswerVariantPlatform = "windows81AndLater"
	SearchAnswerVariantPlatform_WindowsPhone81                     SearchAnswerVariantPlatform = "windowsPhone81"
)

func PossibleValuesForSearchAnswerVariantPlatform() []string {
	return []string{
		string(SearchAnswerVariantPlatform_Android),
		string(SearchAnswerVariantPlatform_AndroidAOSP),
		string(SearchAnswerVariantPlatform_AndroidForWork),
		string(SearchAnswerVariantPlatform_AndroidMobileApplicationManagement),
		string(SearchAnswerVariantPlatform_AndroidWorkProfile),
		string(SearchAnswerVariantPlatform_IOS),
		string(SearchAnswerVariantPlatform_IOSMobileApplicationManagement),
		string(SearchAnswerVariantPlatform_MacOS),
		string(SearchAnswerVariantPlatform_Unknown),
		string(SearchAnswerVariantPlatform_Windows10AndLater),
		string(SearchAnswerVariantPlatform_Windows81AndLater),
		string(SearchAnswerVariantPlatform_WindowsPhone81),
	}
}

func (s *SearchAnswerVariantPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSearchAnswerVariantPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSearchAnswerVariantPlatform(input string) (*SearchAnswerVariantPlatform, error) {
	vals := map[string]SearchAnswerVariantPlatform{
		"android":                            SearchAnswerVariantPlatform_Android,
		"androidaosp":                        SearchAnswerVariantPlatform_AndroidAOSP,
		"androidforwork":                     SearchAnswerVariantPlatform_AndroidForWork,
		"androidmobileapplicationmanagement": SearchAnswerVariantPlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 SearchAnswerVariantPlatform_AndroidWorkProfile,
		"ios":                                SearchAnswerVariantPlatform_IOS,
		"iosmobileapplicationmanagement":     SearchAnswerVariantPlatform_IOSMobileApplicationManagement,
		"macos":                              SearchAnswerVariantPlatform_MacOS,
		"unknown":                            SearchAnswerVariantPlatform_Unknown,
		"windows10andlater":                  SearchAnswerVariantPlatform_Windows10AndLater,
		"windows81andlater":                  SearchAnswerVariantPlatform_Windows81AndLater,
		"windowsphone81":                     SearchAnswerVariantPlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SearchAnswerVariantPlatform(input)
	return &out, nil
}
