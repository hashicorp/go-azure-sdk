package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AppleUserInitiatedEnrollmentProfilePlatform string

const (
	AppleUserInitiatedEnrollmentProfilePlatformandroid                            AppleUserInitiatedEnrollmentProfilePlatform = "Android"
	AppleUserInitiatedEnrollmentProfilePlatformandroidAOSP                        AppleUserInitiatedEnrollmentProfilePlatform = "AndroidAOSP"
	AppleUserInitiatedEnrollmentProfilePlatformandroidForWork                     AppleUserInitiatedEnrollmentProfilePlatform = "AndroidForWork"
	AppleUserInitiatedEnrollmentProfilePlatformandroidMobileApplicationManagement AppleUserInitiatedEnrollmentProfilePlatform = "AndroidMobileApplicationManagement"
	AppleUserInitiatedEnrollmentProfilePlatformandroidWorkProfile                 AppleUserInitiatedEnrollmentProfilePlatform = "AndroidWorkProfile"
	AppleUserInitiatedEnrollmentProfilePlatformiOS                                AppleUserInitiatedEnrollmentProfilePlatform = "IOS"
	AppleUserInitiatedEnrollmentProfilePlatformiOSMobileApplicationManagement     AppleUserInitiatedEnrollmentProfilePlatform = "IOSMobileApplicationManagement"
	AppleUserInitiatedEnrollmentProfilePlatformmacOS                              AppleUserInitiatedEnrollmentProfilePlatform = "MacOS"
	AppleUserInitiatedEnrollmentProfilePlatformunknown                            AppleUserInitiatedEnrollmentProfilePlatform = "Unknown"
	AppleUserInitiatedEnrollmentProfilePlatformwindows10AndLater                  AppleUserInitiatedEnrollmentProfilePlatform = "Windows10AndLater"
	AppleUserInitiatedEnrollmentProfilePlatformwindows81AndLater                  AppleUserInitiatedEnrollmentProfilePlatform = "Windows81AndLater"
	AppleUserInitiatedEnrollmentProfilePlatformwindowsPhone81                     AppleUserInitiatedEnrollmentProfilePlatform = "WindowsPhone81"
)

func PossibleValuesForAppleUserInitiatedEnrollmentProfilePlatform() []string {
	return []string{
		string(AppleUserInitiatedEnrollmentProfilePlatformandroid),
		string(AppleUserInitiatedEnrollmentProfilePlatformandroidAOSP),
		string(AppleUserInitiatedEnrollmentProfilePlatformandroidForWork),
		string(AppleUserInitiatedEnrollmentProfilePlatformandroidMobileApplicationManagement),
		string(AppleUserInitiatedEnrollmentProfilePlatformandroidWorkProfile),
		string(AppleUserInitiatedEnrollmentProfilePlatformiOS),
		string(AppleUserInitiatedEnrollmentProfilePlatformiOSMobileApplicationManagement),
		string(AppleUserInitiatedEnrollmentProfilePlatformmacOS),
		string(AppleUserInitiatedEnrollmentProfilePlatformunknown),
		string(AppleUserInitiatedEnrollmentProfilePlatformwindows10AndLater),
		string(AppleUserInitiatedEnrollmentProfilePlatformwindows81AndLater),
		string(AppleUserInitiatedEnrollmentProfilePlatformwindowsPhone81),
	}
}

func parseAppleUserInitiatedEnrollmentProfilePlatform(input string) (*AppleUserInitiatedEnrollmentProfilePlatform, error) {
	vals := map[string]AppleUserInitiatedEnrollmentProfilePlatform{
		"android":                            AppleUserInitiatedEnrollmentProfilePlatformandroid,
		"androidaosp":                        AppleUserInitiatedEnrollmentProfilePlatformandroidAOSP,
		"androidforwork":                     AppleUserInitiatedEnrollmentProfilePlatformandroidForWork,
		"androidmobileapplicationmanagement": AppleUserInitiatedEnrollmentProfilePlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 AppleUserInitiatedEnrollmentProfilePlatformandroidWorkProfile,
		"ios":                                AppleUserInitiatedEnrollmentProfilePlatformiOS,
		"iosmobileapplicationmanagement":     AppleUserInitiatedEnrollmentProfilePlatformiOSMobileApplicationManagement,
		"macos":                              AppleUserInitiatedEnrollmentProfilePlatformmacOS,
		"unknown":                            AppleUserInitiatedEnrollmentProfilePlatformunknown,
		"windows10andlater":                  AppleUserInitiatedEnrollmentProfilePlatformwindows10AndLater,
		"windows81andlater":                  AppleUserInitiatedEnrollmentProfilePlatformwindows81AndLater,
		"windowsphone81":                     AppleUserInitiatedEnrollmentProfilePlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := AppleUserInitiatedEnrollmentProfilePlatform(input)
	return &out, nil
}
