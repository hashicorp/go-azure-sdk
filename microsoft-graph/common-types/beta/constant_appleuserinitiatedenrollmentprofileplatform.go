package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfilePlatform string

const (
	AppleUserInitiatedEnrollmentProfilePlatform_Android                            AppleUserInitiatedEnrollmentProfilePlatform = "android"
	AppleUserInitiatedEnrollmentProfilePlatform_AndroidAOSP                        AppleUserInitiatedEnrollmentProfilePlatform = "androidAOSP"
	AppleUserInitiatedEnrollmentProfilePlatform_AndroidForWork                     AppleUserInitiatedEnrollmentProfilePlatform = "androidForWork"
	AppleUserInitiatedEnrollmentProfilePlatform_AndroidMobileApplicationManagement AppleUserInitiatedEnrollmentProfilePlatform = "androidMobileApplicationManagement"
	AppleUserInitiatedEnrollmentProfilePlatform_AndroidWorkProfile                 AppleUserInitiatedEnrollmentProfilePlatform = "androidWorkProfile"
	AppleUserInitiatedEnrollmentProfilePlatform_IOS                                AppleUserInitiatedEnrollmentProfilePlatform = "iOS"
	AppleUserInitiatedEnrollmentProfilePlatform_IOSMobileApplicationManagement     AppleUserInitiatedEnrollmentProfilePlatform = "iOSMobileApplicationManagement"
	AppleUserInitiatedEnrollmentProfilePlatform_MacOS                              AppleUserInitiatedEnrollmentProfilePlatform = "macOS"
	AppleUserInitiatedEnrollmentProfilePlatform_Unknown                            AppleUserInitiatedEnrollmentProfilePlatform = "unknown"
	AppleUserInitiatedEnrollmentProfilePlatform_Windows10AndLater                  AppleUserInitiatedEnrollmentProfilePlatform = "windows10AndLater"
	AppleUserInitiatedEnrollmentProfilePlatform_Windows81AndLater                  AppleUserInitiatedEnrollmentProfilePlatform = "windows81AndLater"
	AppleUserInitiatedEnrollmentProfilePlatform_WindowsPhone81                     AppleUserInitiatedEnrollmentProfilePlatform = "windowsPhone81"
)

func PossibleValuesForAppleUserInitiatedEnrollmentProfilePlatform() []string {
	return []string{
		string(AppleUserInitiatedEnrollmentProfilePlatform_Android),
		string(AppleUserInitiatedEnrollmentProfilePlatform_AndroidAOSP),
		string(AppleUserInitiatedEnrollmentProfilePlatform_AndroidForWork),
		string(AppleUserInitiatedEnrollmentProfilePlatform_AndroidMobileApplicationManagement),
		string(AppleUserInitiatedEnrollmentProfilePlatform_AndroidWorkProfile),
		string(AppleUserInitiatedEnrollmentProfilePlatform_IOS),
		string(AppleUserInitiatedEnrollmentProfilePlatform_IOSMobileApplicationManagement),
		string(AppleUserInitiatedEnrollmentProfilePlatform_MacOS),
		string(AppleUserInitiatedEnrollmentProfilePlatform_Unknown),
		string(AppleUserInitiatedEnrollmentProfilePlatform_Windows10AndLater),
		string(AppleUserInitiatedEnrollmentProfilePlatform_Windows81AndLater),
		string(AppleUserInitiatedEnrollmentProfilePlatform_WindowsPhone81),
	}
}

func (s *AppleUserInitiatedEnrollmentProfilePlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseAppleUserInitiatedEnrollmentProfilePlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseAppleUserInitiatedEnrollmentProfilePlatform(input string) (*AppleUserInitiatedEnrollmentProfilePlatform, error) {
	vals := map[string]AppleUserInitiatedEnrollmentProfilePlatform{
		"android":                            AppleUserInitiatedEnrollmentProfilePlatform_Android,
		"androidaosp":                        AppleUserInitiatedEnrollmentProfilePlatform_AndroidAOSP,
		"androidforwork":                     AppleUserInitiatedEnrollmentProfilePlatform_AndroidForWork,
		"androidmobileapplicationmanagement": AppleUserInitiatedEnrollmentProfilePlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 AppleUserInitiatedEnrollmentProfilePlatform_AndroidWorkProfile,
		"ios":                                AppleUserInitiatedEnrollmentProfilePlatform_IOS,
		"iosmobileapplicationmanagement":     AppleUserInitiatedEnrollmentProfilePlatform_IOSMobileApplicationManagement,
		"macos":                              AppleUserInitiatedEnrollmentProfilePlatform_MacOS,
		"unknown":                            AppleUserInitiatedEnrollmentProfilePlatform_Unknown,
		"windows10andlater":                  AppleUserInitiatedEnrollmentProfilePlatform_Windows10AndLater,
		"windows81andlater":                  AppleUserInitiatedEnrollmentProfilePlatform_Windows81AndLater,
		"windowsphone81":                     AppleUserInitiatedEnrollmentProfilePlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleUserInitiatedEnrollmentProfilePlatform(input)
	return &out, nil
}
