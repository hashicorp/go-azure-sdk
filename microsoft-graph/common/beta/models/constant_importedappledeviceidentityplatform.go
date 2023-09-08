package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityPlatform string

const (
	ImportedAppleDeviceIdentityPlatformandroid       ImportedAppleDeviceIdentityPlatform = "Android"
	ImportedAppleDeviceIdentityPlatformios           ImportedAppleDeviceIdentityPlatform = "Ios"
	ImportedAppleDeviceIdentityPlatformmacOS         ImportedAppleDeviceIdentityPlatform = "MacOS"
	ImportedAppleDeviceIdentityPlatformunknown       ImportedAppleDeviceIdentityPlatform = "Unknown"
	ImportedAppleDeviceIdentityPlatformwindows       ImportedAppleDeviceIdentityPlatform = "Windows"
	ImportedAppleDeviceIdentityPlatformwindowsMobile ImportedAppleDeviceIdentityPlatform = "WindowsMobile"
)

func PossibleValuesForImportedAppleDeviceIdentityPlatform() []string {
	return []string{
		string(ImportedAppleDeviceIdentityPlatformandroid),
		string(ImportedAppleDeviceIdentityPlatformios),
		string(ImportedAppleDeviceIdentityPlatformmacOS),
		string(ImportedAppleDeviceIdentityPlatformunknown),
		string(ImportedAppleDeviceIdentityPlatformwindows),
		string(ImportedAppleDeviceIdentityPlatformwindowsMobile),
	}
}

func parseImportedAppleDeviceIdentityPlatform(input string) (*ImportedAppleDeviceIdentityPlatform, error) {
	vals := map[string]ImportedAppleDeviceIdentityPlatform{
		"android":       ImportedAppleDeviceIdentityPlatformandroid,
		"ios":           ImportedAppleDeviceIdentityPlatformios,
		"macos":         ImportedAppleDeviceIdentityPlatformmacOS,
		"unknown":       ImportedAppleDeviceIdentityPlatformunknown,
		"windows":       ImportedAppleDeviceIdentityPlatformwindows,
		"windowsmobile": ImportedAppleDeviceIdentityPlatformwindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityPlatform(input)
	return &out, nil
}
