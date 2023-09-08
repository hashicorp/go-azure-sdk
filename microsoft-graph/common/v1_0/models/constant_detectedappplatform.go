package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DetectedAppPlatform string

const (
	DetectedAppPlatformandroidDedicatedAndFullyManaged DetectedAppPlatform = "AndroidDedicatedAndFullyManaged"
	DetectedAppPlatformandroidDeviceAdministrator      DetectedAppPlatform = "AndroidDeviceAdministrator"
	DetectedAppPlatformandroidOSP                      DetectedAppPlatform = "AndroidOSP"
	DetectedAppPlatformandroidWorkProfile              DetectedAppPlatform = "AndroidWorkProfile"
	DetectedAppPlatformchromeOS                        DetectedAppPlatform = "ChromeOS"
	DetectedAppPlatformios                             DetectedAppPlatform = "Ios"
	DetectedAppPlatformmacOS                           DetectedAppPlatform = "MacOS"
	DetectedAppPlatformunknown                         DetectedAppPlatform = "Unknown"
	DetectedAppPlatformwindows                         DetectedAppPlatform = "Windows"
	DetectedAppPlatformwindowsHolographic              DetectedAppPlatform = "WindowsHolographic"
	DetectedAppPlatformwindowsMobile                   DetectedAppPlatform = "WindowsMobile"
)

func PossibleValuesForDetectedAppPlatform() []string {
	return []string{
		string(DetectedAppPlatformandroidDedicatedAndFullyManaged),
		string(DetectedAppPlatformandroidDeviceAdministrator),
		string(DetectedAppPlatformandroidOSP),
		string(DetectedAppPlatformandroidWorkProfile),
		string(DetectedAppPlatformchromeOS),
		string(DetectedAppPlatformios),
		string(DetectedAppPlatformmacOS),
		string(DetectedAppPlatformunknown),
		string(DetectedAppPlatformwindows),
		string(DetectedAppPlatformwindowsHolographic),
		string(DetectedAppPlatformwindowsMobile),
	}
}

func parseDetectedAppPlatform(input string) (*DetectedAppPlatform, error) {
	vals := map[string]DetectedAppPlatform{
		"androiddedicatedandfullymanaged": DetectedAppPlatformandroidDedicatedAndFullyManaged,
		"androiddeviceadministrator":      DetectedAppPlatformandroidDeviceAdministrator,
		"androidosp":                      DetectedAppPlatformandroidOSP,
		"androidworkprofile":              DetectedAppPlatformandroidWorkProfile,
		"chromeos":                        DetectedAppPlatformchromeOS,
		"ios":                             DetectedAppPlatformios,
		"macos":                           DetectedAppPlatformmacOS,
		"unknown":                         DetectedAppPlatformunknown,
		"windows":                         DetectedAppPlatformwindows,
		"windowsholographic":              DetectedAppPlatformwindowsHolographic,
		"windowsmobile":                   DetectedAppPlatformwindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DetectedAppPlatform(input)
	return &out, nil
}
