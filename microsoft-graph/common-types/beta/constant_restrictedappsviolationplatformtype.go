package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsViolationPlatformType string

const (
	RestrictedAppsViolationPlatformType_All                RestrictedAppsViolationPlatformType = "all"
	RestrictedAppsViolationPlatformType_Android            RestrictedAppsViolationPlatformType = "android"
	RestrictedAppsViolationPlatformType_AndroidAOSP        RestrictedAppsViolationPlatformType = "androidAOSP"
	RestrictedAppsViolationPlatformType_AndroidForWork     RestrictedAppsViolationPlatformType = "androidForWork"
	RestrictedAppsViolationPlatformType_AndroidWorkProfile RestrictedAppsViolationPlatformType = "androidWorkProfile"
	RestrictedAppsViolationPlatformType_IOS                RestrictedAppsViolationPlatformType = "iOS"
	RestrictedAppsViolationPlatformType_MacOS              RestrictedAppsViolationPlatformType = "macOS"
	RestrictedAppsViolationPlatformType_Windows10AndLater  RestrictedAppsViolationPlatformType = "windows10AndLater"
	RestrictedAppsViolationPlatformType_Windows10XProfile  RestrictedAppsViolationPlatformType = "windows10XProfile"
	RestrictedAppsViolationPlatformType_Windows81AndLater  RestrictedAppsViolationPlatformType = "windows81AndLater"
	RestrictedAppsViolationPlatformType_WindowsPhone81     RestrictedAppsViolationPlatformType = "windowsPhone81"
)

func PossibleValuesForRestrictedAppsViolationPlatformType() []string {
	return []string{
		string(RestrictedAppsViolationPlatformType_All),
		string(RestrictedAppsViolationPlatformType_Android),
		string(RestrictedAppsViolationPlatformType_AndroidAOSP),
		string(RestrictedAppsViolationPlatformType_AndroidForWork),
		string(RestrictedAppsViolationPlatformType_AndroidWorkProfile),
		string(RestrictedAppsViolationPlatformType_IOS),
		string(RestrictedAppsViolationPlatformType_MacOS),
		string(RestrictedAppsViolationPlatformType_Windows10AndLater),
		string(RestrictedAppsViolationPlatformType_Windows10XProfile),
		string(RestrictedAppsViolationPlatformType_Windows81AndLater),
		string(RestrictedAppsViolationPlatformType_WindowsPhone81),
	}
}

func (s *RestrictedAppsViolationPlatformType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRestrictedAppsViolationPlatformType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRestrictedAppsViolationPlatformType(input string) (*RestrictedAppsViolationPlatformType, error) {
	vals := map[string]RestrictedAppsViolationPlatformType{
		"all":                RestrictedAppsViolationPlatformType_All,
		"android":            RestrictedAppsViolationPlatformType_Android,
		"androidaosp":        RestrictedAppsViolationPlatformType_AndroidAOSP,
		"androidforwork":     RestrictedAppsViolationPlatformType_AndroidForWork,
		"androidworkprofile": RestrictedAppsViolationPlatformType_AndroidWorkProfile,
		"ios":                RestrictedAppsViolationPlatformType_IOS,
		"macos":              RestrictedAppsViolationPlatformType_MacOS,
		"windows10andlater":  RestrictedAppsViolationPlatformType_Windows10AndLater,
		"windows10xprofile":  RestrictedAppsViolationPlatformType_Windows10XProfile,
		"windows81andlater":  RestrictedAppsViolationPlatformType_Windows81AndLater,
		"windowsphone81":     RestrictedAppsViolationPlatformType_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RestrictedAppsViolationPlatformType(input)
	return &out, nil
}
