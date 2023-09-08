package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityPlatform string

const (
	ImportedDeviceIdentityPlatformandroid       ImportedDeviceIdentityPlatform = "Android"
	ImportedDeviceIdentityPlatformios           ImportedDeviceIdentityPlatform = "Ios"
	ImportedDeviceIdentityPlatformmacOS         ImportedDeviceIdentityPlatform = "MacOS"
	ImportedDeviceIdentityPlatformunknown       ImportedDeviceIdentityPlatform = "Unknown"
	ImportedDeviceIdentityPlatformwindows       ImportedDeviceIdentityPlatform = "Windows"
	ImportedDeviceIdentityPlatformwindowsMobile ImportedDeviceIdentityPlatform = "WindowsMobile"
)

func PossibleValuesForImportedDeviceIdentityPlatform() []string {
	return []string{
		string(ImportedDeviceIdentityPlatformandroid),
		string(ImportedDeviceIdentityPlatformios),
		string(ImportedDeviceIdentityPlatformmacOS),
		string(ImportedDeviceIdentityPlatformunknown),
		string(ImportedDeviceIdentityPlatformwindows),
		string(ImportedDeviceIdentityPlatformwindowsMobile),
	}
}

func parseImportedDeviceIdentityPlatform(input string) (*ImportedDeviceIdentityPlatform, error) {
	vals := map[string]ImportedDeviceIdentityPlatform{
		"android":       ImportedDeviceIdentityPlatformandroid,
		"ios":           ImportedDeviceIdentityPlatformios,
		"macos":         ImportedDeviceIdentityPlatformmacOS,
		"unknown":       ImportedDeviceIdentityPlatformunknown,
		"windows":       ImportedDeviceIdentityPlatformwindows,
		"windowsmobile": ImportedDeviceIdentityPlatformwindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityPlatform(input)
	return &out, nil
}
