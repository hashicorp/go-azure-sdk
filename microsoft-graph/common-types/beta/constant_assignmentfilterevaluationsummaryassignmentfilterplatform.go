package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentFilterEvaluationSummaryAssignmentFilterPlatform string

const (
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Android                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "android"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidAOSP                        AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "androidAOSP"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidForWork                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "androidForWork"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidMobileApplicationManagement AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "androidMobileApplicationManagement"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidWorkProfile                 AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "androidWorkProfile"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_IOS                                AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "iOS"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_IOSMobileApplicationManagement     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "iOSMobileApplicationManagement"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_MacOS                              AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "macOS"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Unknown                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "unknown"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Windows10AndLater                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "windows10AndLater"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Windows81AndLater                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "windows81AndLater"
	AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_WindowsPhone81                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform = "windowsPhone81"
)

func PossibleValuesForAssignmentFilterEvaluationSummaryAssignmentFilterPlatform() []string {
	return []string{
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Android),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidAOSP),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidForWork),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidMobileApplicationManagement),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidWorkProfile),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_IOS),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_IOSMobileApplicationManagement),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_MacOS),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Unknown),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Windows10AndLater),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Windows81AndLater),
		string(AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_WindowsPhone81),
	}
}

func (s *AssignmentFilterEvaluationSummaryAssignmentFilterPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAssignmentFilterEvaluationSummaryAssignmentFilterPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAssignmentFilterEvaluationSummaryAssignmentFilterPlatform(input string) (*AssignmentFilterEvaluationSummaryAssignmentFilterPlatform, error) {
	vals := map[string]AssignmentFilterEvaluationSummaryAssignmentFilterPlatform{
		"android":                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Android,
		"androidaosp":                        AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidAOSP,
		"androidforwork":                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidForWork,
		"androidmobileapplicationmanagement": AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_AndroidWorkProfile,
		"ios":                                AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_IOS,
		"iosmobileapplicationmanagement":     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_IOSMobileApplicationManagement,
		"macos":                              AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_MacOS,
		"unknown":                            AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Unknown,
		"windows10andlater":                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Windows10AndLater,
		"windows81andlater":                  AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_Windows81AndLater,
		"windowsphone81":                     AssignmentFilterEvaluationSummaryAssignmentFilterPlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AssignmentFilterEvaluationSummaryAssignmentFilterPlatform(input)
	return &out, nil
}
