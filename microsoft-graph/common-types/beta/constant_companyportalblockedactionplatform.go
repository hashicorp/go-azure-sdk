package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedActionPlatform string

const (
	CompanyPortalBlockedActionPlatform_Android                            CompanyPortalBlockedActionPlatform = "android"
	CompanyPortalBlockedActionPlatform_AndroidAOSP                        CompanyPortalBlockedActionPlatform = "androidAOSP"
	CompanyPortalBlockedActionPlatform_AndroidForWork                     CompanyPortalBlockedActionPlatform = "androidForWork"
	CompanyPortalBlockedActionPlatform_AndroidMobileApplicationManagement CompanyPortalBlockedActionPlatform = "androidMobileApplicationManagement"
	CompanyPortalBlockedActionPlatform_AndroidWorkProfile                 CompanyPortalBlockedActionPlatform = "androidWorkProfile"
	CompanyPortalBlockedActionPlatform_IOS                                CompanyPortalBlockedActionPlatform = "iOS"
	CompanyPortalBlockedActionPlatform_IOSMobileApplicationManagement     CompanyPortalBlockedActionPlatform = "iOSMobileApplicationManagement"
	CompanyPortalBlockedActionPlatform_MacOS                              CompanyPortalBlockedActionPlatform = "macOS"
	CompanyPortalBlockedActionPlatform_Unknown                            CompanyPortalBlockedActionPlatform = "unknown"
	CompanyPortalBlockedActionPlatform_Windows10AndLater                  CompanyPortalBlockedActionPlatform = "windows10AndLater"
	CompanyPortalBlockedActionPlatform_Windows81AndLater                  CompanyPortalBlockedActionPlatform = "windows81AndLater"
	CompanyPortalBlockedActionPlatform_WindowsPhone81                     CompanyPortalBlockedActionPlatform = "windowsPhone81"
)

func PossibleValuesForCompanyPortalBlockedActionPlatform() []string {
	return []string{
		string(CompanyPortalBlockedActionPlatform_Android),
		string(CompanyPortalBlockedActionPlatform_AndroidAOSP),
		string(CompanyPortalBlockedActionPlatform_AndroidForWork),
		string(CompanyPortalBlockedActionPlatform_AndroidMobileApplicationManagement),
		string(CompanyPortalBlockedActionPlatform_AndroidWorkProfile),
		string(CompanyPortalBlockedActionPlatform_IOS),
		string(CompanyPortalBlockedActionPlatform_IOSMobileApplicationManagement),
		string(CompanyPortalBlockedActionPlatform_MacOS),
		string(CompanyPortalBlockedActionPlatform_Unknown),
		string(CompanyPortalBlockedActionPlatform_Windows10AndLater),
		string(CompanyPortalBlockedActionPlatform_Windows81AndLater),
		string(CompanyPortalBlockedActionPlatform_WindowsPhone81),
	}
}

func (s *CompanyPortalBlockedActionPlatform) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseCompanyPortalBlockedActionPlatform(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseCompanyPortalBlockedActionPlatform(input string) (*CompanyPortalBlockedActionPlatform, error) {
	vals := map[string]CompanyPortalBlockedActionPlatform{
		"android":                            CompanyPortalBlockedActionPlatform_Android,
		"androidaosp":                        CompanyPortalBlockedActionPlatform_AndroidAOSP,
		"androidforwork":                     CompanyPortalBlockedActionPlatform_AndroidForWork,
		"androidmobileapplicationmanagement": CompanyPortalBlockedActionPlatform_AndroidMobileApplicationManagement,
		"androidworkprofile":                 CompanyPortalBlockedActionPlatform_AndroidWorkProfile,
		"ios":                                CompanyPortalBlockedActionPlatform_IOS,
		"iosmobileapplicationmanagement":     CompanyPortalBlockedActionPlatform_IOSMobileApplicationManagement,
		"macos":                              CompanyPortalBlockedActionPlatform_MacOS,
		"unknown":                            CompanyPortalBlockedActionPlatform_Unknown,
		"windows10andlater":                  CompanyPortalBlockedActionPlatform_Windows10AndLater,
		"windows81andlater":                  CompanyPortalBlockedActionPlatform_Windows81AndLater,
		"windowsphone81":                     CompanyPortalBlockedActionPlatform_WindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompanyPortalBlockedActionPlatform(input)
	return &out, nil
}
