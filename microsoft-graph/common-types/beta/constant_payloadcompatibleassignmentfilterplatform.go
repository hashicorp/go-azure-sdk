package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadCompatibleAssignmentFilterPlatform string

const (
	PayloadCompatibleAssignmentFilterPlatform_Android                            PayloadCompatibleAssignmentFilterPlatform = "android"
	PayloadCompatibleAssignmentFilterPlatform_AndroidAOSP                        PayloadCompatibleAssignmentFilterPlatform = "androidAOSP"
	PayloadCompatibleAssignmentFilterPlatform_AndroidForWork                     PayloadCompatibleAssignmentFilterPlatform = "androidForWork"
	PayloadCompatibleAssignmentFilterPlatform_AndroidMobileApplicationManagement PayloadCompatibleAssignmentFilterPlatform = "androidMobileApplicationManagement"
	PayloadCompatibleAssignmentFilterPlatform_AndroidWorkProfile                 PayloadCompatibleAssignmentFilterPlatform = "androidWorkProfile"
	PayloadCompatibleAssignmentFilterPlatform_IOS                                PayloadCompatibleAssignmentFilterPlatform = "iOS"
	PayloadCompatibleAssignmentFilterPlatform_IOSMobileApplicationManagement     PayloadCompatibleAssignmentFilterPlatform = "iOSMobileApplicationManagement"
	PayloadCompatibleAssignmentFilterPlatform_MacOS                              PayloadCompatibleAssignmentFilterPlatform = "macOS"
	PayloadCompatibleAssignmentFilterPlatform_Unknown                            PayloadCompatibleAssignmentFilterPlatform = "unknown"
	PayloadCompatibleAssignmentFilterPlatform_Windows10AndLater                  PayloadCompatibleAssignmentFilterPlatform = "windows10AndLater"
	PayloadCompatibleAssignmentFilterPlatform_Windows81AndLater                  PayloadCompatibleAssignmentFilterPlatform = "windows81AndLater"
	PayloadCompatibleAssignmentFilterPlatform_WindowsPhone81                     PayloadCompatibleAssignmentFilterPlatform = "windowsPhone81"
)

func PossibleValuesForPayloadCompatibleAssignmentFilterPlatform() []string {
	return []string{
		string(PayloadCompatibleAssignmentFilterPlatform_Android),
		string(PayloadCompatibleAssignmentFilterPlatform_AndroidAOSP),
		string(PayloadCompatibleAssignmentFilterPlatform_AndroidForWork),
		string(PayloadCompatibleAssignmentFilterPlatform_AndroidMobileApplicationManagement),
		string(PayloadCompatibleAssignmentFilterPlatform_AndroidWorkProfile),
		string(PayloadCompatibleAssignmentFilterPlatform_IOS),
		string(PayloadCompatibleAssignmentFilterPlatform_IOSMobileApplicationManagement),
		string(PayloadCompatibleAssignmentFilterPlatform_MacOS),
		string(PayloadCompatibleAssignmentFilterPlatform_Unknown),
		string(PayloadCompatibleAssignmentFilterPlatform_Windows10AndLater),
		string(PayloadCompatibleAssignmentFilterPlatform_Windows81AndLater),
		string(PayloadCompatibleAssignmentFilterPlatform_WindowsPhone81),
	}
}

func (s *PayloadCompatibleAssignmentFilterPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadCompatibleAssignmentFilterPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadCompatibleAssignmentFilterPlatform(input string) (*PayloadCompatibleAssignmentFilterPlatform, error) {
	vals := map[string]PayloadCompatibleAssignmentFilterPlatform{
		"android":                            PayloadCompatibleAssignmentFilterPlatform_Android,
		"androidaosp":                        PayloadCompatibleAssignmentFilterPlatform_AndroidAOSP,
		"androidforwork":                     PayloadCompatibleAssignmentFilterPlatform_AndroidForWork,
		"androidmobileapplicationmanagement": PayloadCompatibleAssignmentFilterPlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 PayloadCompatibleAssignmentFilterPlatform_AndroidWorkProfile,
		"ios":                                PayloadCompatibleAssignmentFilterPlatform_IOS,
		"iosmobileapplicationmanagement":     PayloadCompatibleAssignmentFilterPlatform_IOSMobileApplicationManagement,
		"macos":                              PayloadCompatibleAssignmentFilterPlatform_MacOS,
		"unknown":                            PayloadCompatibleAssignmentFilterPlatform_Unknown,
		"windows10andlater":                  PayloadCompatibleAssignmentFilterPlatform_Windows10AndLater,
		"windows81andlater":                  PayloadCompatibleAssignmentFilterPlatform_Windows81AndLater,
		"windowsphone81":                     PayloadCompatibleAssignmentFilterPlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadCompatibleAssignmentFilterPlatform(input)
	return &out, nil
}
