package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCertificateStateDevicePlatform string

const (
	ManagedDeviceCertificateStateDevicePlatformandroid                            ManagedDeviceCertificateStateDevicePlatform = "Android"
	ManagedDeviceCertificateStateDevicePlatformandroidAOSP                        ManagedDeviceCertificateStateDevicePlatform = "AndroidAOSP"
	ManagedDeviceCertificateStateDevicePlatformandroidForWork                     ManagedDeviceCertificateStateDevicePlatform = "AndroidForWork"
	ManagedDeviceCertificateStateDevicePlatformandroidMobileApplicationManagement ManagedDeviceCertificateStateDevicePlatform = "AndroidMobileApplicationManagement"
	ManagedDeviceCertificateStateDevicePlatformandroidWorkProfile                 ManagedDeviceCertificateStateDevicePlatform = "AndroidWorkProfile"
	ManagedDeviceCertificateStateDevicePlatformiOS                                ManagedDeviceCertificateStateDevicePlatform = "IOS"
	ManagedDeviceCertificateStateDevicePlatformiOSMobileApplicationManagement     ManagedDeviceCertificateStateDevicePlatform = "IOSMobileApplicationManagement"
	ManagedDeviceCertificateStateDevicePlatformmacOS                              ManagedDeviceCertificateStateDevicePlatform = "MacOS"
	ManagedDeviceCertificateStateDevicePlatformunknown                            ManagedDeviceCertificateStateDevicePlatform = "Unknown"
	ManagedDeviceCertificateStateDevicePlatformwindows10AndLater                  ManagedDeviceCertificateStateDevicePlatform = "Windows10AndLater"
	ManagedDeviceCertificateStateDevicePlatformwindows81AndLater                  ManagedDeviceCertificateStateDevicePlatform = "Windows81AndLater"
	ManagedDeviceCertificateStateDevicePlatformwindowsPhone81                     ManagedDeviceCertificateStateDevicePlatform = "WindowsPhone81"
)

func PossibleValuesForManagedDeviceCertificateStateDevicePlatform() []string {
	return []string{
		string(ManagedDeviceCertificateStateDevicePlatformandroid),
		string(ManagedDeviceCertificateStateDevicePlatformandroidAOSP),
		string(ManagedDeviceCertificateStateDevicePlatformandroidForWork),
		string(ManagedDeviceCertificateStateDevicePlatformandroidMobileApplicationManagement),
		string(ManagedDeviceCertificateStateDevicePlatformandroidWorkProfile),
		string(ManagedDeviceCertificateStateDevicePlatformiOS),
		string(ManagedDeviceCertificateStateDevicePlatformiOSMobileApplicationManagement),
		string(ManagedDeviceCertificateStateDevicePlatformmacOS),
		string(ManagedDeviceCertificateStateDevicePlatformunknown),
		string(ManagedDeviceCertificateStateDevicePlatformwindows10AndLater),
		string(ManagedDeviceCertificateStateDevicePlatformwindows81AndLater),
		string(ManagedDeviceCertificateStateDevicePlatformwindowsPhone81),
	}
}

func parseManagedDeviceCertificateStateDevicePlatform(input string) (*ManagedDeviceCertificateStateDevicePlatform, error) {
	vals := map[string]ManagedDeviceCertificateStateDevicePlatform{
		"android":                            ManagedDeviceCertificateStateDevicePlatformandroid,
		"androidaosp":                        ManagedDeviceCertificateStateDevicePlatformandroidAOSP,
		"androidforwork":                     ManagedDeviceCertificateStateDevicePlatformandroidForWork,
		"androidmobileapplicationmanagement": ManagedDeviceCertificateStateDevicePlatformandroidMobileApplicationManagement,
		"androidworkprofile":                 ManagedDeviceCertificateStateDevicePlatformandroidWorkProfile,
		"ios":                                ManagedDeviceCertificateStateDevicePlatformiOS,
		"iosmobileapplicationmanagement":     ManagedDeviceCertificateStateDevicePlatformiOSMobileApplicationManagement,
		"macos":                              ManagedDeviceCertificateStateDevicePlatformmacOS,
		"unknown":                            ManagedDeviceCertificateStateDevicePlatformunknown,
		"windows10andlater":                  ManagedDeviceCertificateStateDevicePlatformwindows10AndLater,
		"windows81andlater":                  ManagedDeviceCertificateStateDevicePlatformwindows81AndLater,
		"windowsphone81":                     ManagedDeviceCertificateStateDevicePlatformwindowsPhone81,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ManagedDeviceCertificateStateDevicePlatform(input)
	return &out, nil
}
