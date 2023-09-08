package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedAppleDeviceIdentityResultPlatform string

const (
	ImportedAppleDeviceIdentityResultPlatformandroid       ImportedAppleDeviceIdentityResultPlatform = "Android"
	ImportedAppleDeviceIdentityResultPlatformios           ImportedAppleDeviceIdentityResultPlatform = "Ios"
	ImportedAppleDeviceIdentityResultPlatformmacOS         ImportedAppleDeviceIdentityResultPlatform = "MacOS"
	ImportedAppleDeviceIdentityResultPlatformunknown       ImportedAppleDeviceIdentityResultPlatform = "Unknown"
	ImportedAppleDeviceIdentityResultPlatformwindows       ImportedAppleDeviceIdentityResultPlatform = "Windows"
	ImportedAppleDeviceIdentityResultPlatformwindowsMobile ImportedAppleDeviceIdentityResultPlatform = "WindowsMobile"
)

func PossibleValuesForImportedAppleDeviceIdentityResultPlatform() []string {
	return []string{
		string(ImportedAppleDeviceIdentityResultPlatformandroid),
		string(ImportedAppleDeviceIdentityResultPlatformios),
		string(ImportedAppleDeviceIdentityResultPlatformmacOS),
		string(ImportedAppleDeviceIdentityResultPlatformunknown),
		string(ImportedAppleDeviceIdentityResultPlatformwindows),
		string(ImportedAppleDeviceIdentityResultPlatformwindowsMobile),
	}
}

func parseImportedAppleDeviceIdentityResultPlatform(input string) (*ImportedAppleDeviceIdentityResultPlatform, error) {
	vals := map[string]ImportedAppleDeviceIdentityResultPlatform{
		"android":       ImportedAppleDeviceIdentityResultPlatformandroid,
		"ios":           ImportedAppleDeviceIdentityResultPlatformios,
		"macos":         ImportedAppleDeviceIdentityResultPlatformmacOS,
		"unknown":       ImportedAppleDeviceIdentityResultPlatformunknown,
		"windows":       ImportedAppleDeviceIdentityResultPlatformwindows,
		"windowsmobile": ImportedAppleDeviceIdentityResultPlatformwindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedAppleDeviceIdentityResultPlatform(input)
	return &out, nil
}
