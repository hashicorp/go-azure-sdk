package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluateRequestPlatform string

const (
	AssignmentFilterEvaluateRequestPlatform_Android                            AssignmentFilterEvaluateRequestPlatform = "android"
	AssignmentFilterEvaluateRequestPlatform_AndroidAOSP                        AssignmentFilterEvaluateRequestPlatform = "androidAOSP"
	AssignmentFilterEvaluateRequestPlatform_AndroidForWork                     AssignmentFilterEvaluateRequestPlatform = "androidForWork"
	AssignmentFilterEvaluateRequestPlatform_AndroidMobileApplicationManagement AssignmentFilterEvaluateRequestPlatform = "androidMobileApplicationManagement"
	AssignmentFilterEvaluateRequestPlatform_AndroidWorkProfile                 AssignmentFilterEvaluateRequestPlatform = "androidWorkProfile"
	AssignmentFilterEvaluateRequestPlatform_IOS                                AssignmentFilterEvaluateRequestPlatform = "iOS"
	AssignmentFilterEvaluateRequestPlatform_IOSMobileApplicationManagement     AssignmentFilterEvaluateRequestPlatform = "iOSMobileApplicationManagement"
	AssignmentFilterEvaluateRequestPlatform_MacOS                              AssignmentFilterEvaluateRequestPlatform = "macOS"
	AssignmentFilterEvaluateRequestPlatform_Unknown                            AssignmentFilterEvaluateRequestPlatform = "unknown"
	AssignmentFilterEvaluateRequestPlatform_Windows10AndLater                  AssignmentFilterEvaluateRequestPlatform = "windows10AndLater"
	AssignmentFilterEvaluateRequestPlatform_Windows81AndLater                  AssignmentFilterEvaluateRequestPlatform = "windows81AndLater"
	AssignmentFilterEvaluateRequestPlatform_WindowsPhone81                     AssignmentFilterEvaluateRequestPlatform = "windowsPhone81"
)

func PossibleValuesForAssignmentFilterEvaluateRequestPlatform() []string {
	return []string{
		string(AssignmentFilterEvaluateRequestPlatform_Android),
		string(AssignmentFilterEvaluateRequestPlatform_AndroidAOSP),
		string(AssignmentFilterEvaluateRequestPlatform_AndroidForWork),
		string(AssignmentFilterEvaluateRequestPlatform_AndroidMobileApplicationManagement),
		string(AssignmentFilterEvaluateRequestPlatform_AndroidWorkProfile),
		string(AssignmentFilterEvaluateRequestPlatform_IOS),
		string(AssignmentFilterEvaluateRequestPlatform_IOSMobileApplicationManagement),
		string(AssignmentFilterEvaluateRequestPlatform_MacOS),
		string(AssignmentFilterEvaluateRequestPlatform_Unknown),
		string(AssignmentFilterEvaluateRequestPlatform_Windows10AndLater),
		string(AssignmentFilterEvaluateRequestPlatform_Windows81AndLater),
		string(AssignmentFilterEvaluateRequestPlatform_WindowsPhone81),
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
		"android":                            AssignmentFilterEvaluateRequestPlatform_Android,
		"androidaosp":                        AssignmentFilterEvaluateRequestPlatform_AndroidAOSP,
		"androidforwork":                     AssignmentFilterEvaluateRequestPlatform_AndroidForWork,
		"androidmobileapplicationmanagement": AssignmentFilterEvaluateRequestPlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 AssignmentFilterEvaluateRequestPlatform_AndroidWorkProfile,
		"ios":                                AssignmentFilterEvaluateRequestPlatform_IOS,
		"iosmobileapplicationmanagement":     AssignmentFilterEvaluateRequestPlatform_IOSMobileApplicationManagement,
		"macos":                              AssignmentFilterEvaluateRequestPlatform_MacOS,
		"unknown":                            AssignmentFilterEvaluateRequestPlatform_Unknown,
		"windows10andlater":                  AssignmentFilterEvaluateRequestPlatform_Windows10AndLater,
		"windows81andlater":                  AssignmentFilterEvaluateRequestPlatform_Windows81AndLater,
		"windowsphone81":                     AssignmentFilterEvaluateRequestPlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluateRequestPlatform(input)
	return &out, nil
}
