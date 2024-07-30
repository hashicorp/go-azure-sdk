package devicemanagement

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluateRequestPlatform string

const (
	AssignmentFilterEvaluateRequestPlatformAndroid                            AssignmentFilterEvaluateRequestPlatform = "android"
	AssignmentFilterEvaluateRequestPlatformAndroidAOSP                        AssignmentFilterEvaluateRequestPlatform = "androidAOSP"
	AssignmentFilterEvaluateRequestPlatformAndroidForWork                     AssignmentFilterEvaluateRequestPlatform = "androidForWork"
	AssignmentFilterEvaluateRequestPlatformAndroidMobileApplicationManagement AssignmentFilterEvaluateRequestPlatform = "androidMobileApplicationManagement"
	AssignmentFilterEvaluateRequestPlatformAndroidWorkProfile                 AssignmentFilterEvaluateRequestPlatform = "androidWorkProfile"
	AssignmentFilterEvaluateRequestPlatformIOS                                AssignmentFilterEvaluateRequestPlatform = "iOS"
	AssignmentFilterEvaluateRequestPlatformIOSMobileApplicationManagement     AssignmentFilterEvaluateRequestPlatform = "iOSMobileApplicationManagement"
	AssignmentFilterEvaluateRequestPlatformMacOS                              AssignmentFilterEvaluateRequestPlatform = "macOS"
	AssignmentFilterEvaluateRequestPlatformUnknown                            AssignmentFilterEvaluateRequestPlatform = "unknown"
	AssignmentFilterEvaluateRequestPlatformWindows10AndLater                  AssignmentFilterEvaluateRequestPlatform = "windows10AndLater"
	AssignmentFilterEvaluateRequestPlatformWindows81AndLater                  AssignmentFilterEvaluateRequestPlatform = "windows81AndLater"
	AssignmentFilterEvaluateRequestPlatformWindowsPhone81                     AssignmentFilterEvaluateRequestPlatform = "windowsPhone81"
)

func PossibleValuesForAssignmentFilterEvaluateRequestPlatform() []string {
	return []string{
		string(AssignmentFilterEvaluateRequestPlatformAndroid),
		string(AssignmentFilterEvaluateRequestPlatformAndroidAOSP),
		string(AssignmentFilterEvaluateRequestPlatformAndroidForWork),
		string(AssignmentFilterEvaluateRequestPlatformAndroidMobileApplicationManagement),
		string(AssignmentFilterEvaluateRequestPlatformAndroidWorkProfile),
		string(AssignmentFilterEvaluateRequestPlatformIOS),
		string(AssignmentFilterEvaluateRequestPlatformIOSMobileApplicationManagement),
		string(AssignmentFilterEvaluateRequestPlatformMacOS),
		string(AssignmentFilterEvaluateRequestPlatformUnknown),
		string(AssignmentFilterEvaluateRequestPlatformWindows10AndLater),
		string(AssignmentFilterEvaluateRequestPlatformWindows81AndLater),
		string(AssignmentFilterEvaluateRequestPlatformWindowsPhone81),
	}
}

func (s *AssignmentFilterEvaluateRequestPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterEvaluateRequestPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterEvaluateRequestPlatform(input string) (*AssignmentFilterEvaluateRequestPlatform, error) {
	vals := map[string]AssignmentFilterEvaluateRequestPlatform{
		"android":                            AssignmentFilterEvaluateRequestPlatformAndroid,
		"androidaosp":                        AssignmentFilterEvaluateRequestPlatformAndroidAOSP,
		"androidforwork":                     AssignmentFilterEvaluateRequestPlatformAndroidForWork,
		"androidmobileapplicationmanagement": AssignmentFilterEvaluateRequestPlatformAndroidMobileApplicationManagement,
		"androidworkprofile":                 AssignmentFilterEvaluateRequestPlatformAndroidWorkProfile,
		"ios":                                AssignmentFilterEvaluateRequestPlatformIOS,
		"iosmobileapplicationmanagement":     AssignmentFilterEvaluateRequestPlatformIOSMobileApplicationManagement,
		"macos":                              AssignmentFilterEvaluateRequestPlatformMacOS,
		"unknown":                            AssignmentFilterEvaluateRequestPlatformUnknown,
		"windows10andlater":                  AssignmentFilterEvaluateRequestPlatformWindows10AndLater,
		"windows81andlater":                  AssignmentFilterEvaluateRequestPlatformWindows81AndLater,
		"windowsphone81":                     AssignmentFilterEvaluateRequestPlatformWindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluateRequestPlatform(input)
	return &out, nil
}
