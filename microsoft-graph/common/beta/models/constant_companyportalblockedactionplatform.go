package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CompanyPortalBlockedActionPlatform string

const (
	CompanyPortalBlockedActionPlatformandroid                            CompanyPortalBlockedActionPlatform = "Android"
	CompanyPortalBlockedActionPlatformandroidAOSP                        CompanyPortalBlockedActionPlatform = "AndroidAOSP"
	CompanyPortalBlockedActionPlatformandroidForWork                     CompanyPortalBlockedActionPlatform = "AndroidForWork"
	CompanyPortalBlockedActionPlatformandroidMobileApplicationManagement CompanyPortalBlockedActionPlatform = "AndroidMobileApplicationManagement"
	CompanyPortalBlockedActionPlatformandroidWorkProfile                 CompanyPortalBlockedActionPlatform = "AndroidWorkProfile"
	CompanyPortalBlockedActionPlatformiOS                                CompanyPortalBlockedActionPlatform = "IOS"
	CompanyPortalBlockedActionPlatformiOSMobileApplicationManagement     CompanyPortalBlockedActionPlatform = "IOSMobileApplicationManagement"
	CompanyPortalBlockedActionPlatformmacOS                              CompanyPortalBlockedActionPlatform = "MacOS"
	CompanyPortalBlockedActionPlatformunknown                            CompanyPortalBlockedActionPlatform = "Unknown"
	CompanyPortalBlockedActionPlatformwindows10AndLater                  CompanyPortalBlockedActionPlatform = "Windows10AndLater"
	CompanyPortalBlockedActionPlatformwindows81AndLater                  CompanyPortalBlockedActionPlatform = "Windows81AndLater"
	CompanyPortalBlockedActionPlatformwindowsPhone81                     CompanyPortalBlockedActionPlatform = "WindowsPhone81"
)

func PossibleValuesForCompanyPortalBlockedActionPlatform() []string {
	return []string{
		string(CompanyPortalBlockedActionPlatformandroid),
		string(CompanyPortalBlockedActionPlatformandroidAOSP),
		string(CompanyPortalBlockedActionPlatformandroidForWork),
		string(CompanyPortalBlockedActionPlatformandroidMobileApplicationManagement),
		string(CompanyPortalBlockedActionPlatformandroidWorkProfile),
		string(CompanyPortalBlockedActionPlatformiOS),
		string(CompanyPortalBlockedActionPlatformiOSMobileApplicationManagement),
		string(CompanyPortalBlockedActionPlatformmacOS),
		string(CompanyPortalBlockedActionPlatformunknown),
		string(CompanyPortalBlockedActionPlatformwindows10AndLater),
		string(CompanyPortalBlockedActionPlatformwindows81AndLater),
		string(CompanyPortalBlockedActionPlatformwindowsPhone81),
	}
}

func parseCompanyPortalBlockedActionPlatform(input string) (*CompanyPortalBlockedActionPlatform, error) {
	vals := map[string]CompanyPortalBlockedActionPlatform{
		"android":                            CompanyPortalBlockedActionPlatformandroid,
		"androidaosp":                        CompanyPortalBlockedActionPlatformandroidAOSP,
		"androidforwork":                     CompanyPortalBlockedActionPlatformandroidForWork,
		"androidmobileapplicationmanagement": CompanyPortalBlockedActionPlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 CompanyPortalBlockedActionPlatformandroidWorkProfile,
		"ios":                                CompanyPortalBlockedActionPlatformiOS,
		"iosmobileapplicationmanagement":     CompanyPortalBlockedActionPlatformiOSMobileApplicationManagement,
		"macos":                              CompanyPortalBlockedActionPlatformmacOS,
		"unknown":                            CompanyPortalBlockedActionPlatformunknown,
		"windows10andlater":                  CompanyPortalBlockedActionPlatformwindows10AndLater,
		"windows81andlater":                  CompanyPortalBlockedActionPlatformwindows81AndLater,
		"windowsphone81":                     CompanyPortalBlockedActionPlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := CompanyPortalBlockedActionPlatform(input)
	return &out, nil
}
