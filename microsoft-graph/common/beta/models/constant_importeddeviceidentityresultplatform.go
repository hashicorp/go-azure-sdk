package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ImportedDeviceIdentityResultPlatform string

const (
	ImportedDeviceIdentityResultPlatformandroid       ImportedDeviceIdentityResultPlatform = "Android"
	ImportedDeviceIdentityResultPlatformios           ImportedDeviceIdentityResultPlatform = "Ios"
	ImportedDeviceIdentityResultPlatformmacOS         ImportedDeviceIdentityResultPlatform = "MacOS"
	ImportedDeviceIdentityResultPlatformunknown       ImportedDeviceIdentityResultPlatform = "Unknown"
	ImportedDeviceIdentityResultPlatformwindows       ImportedDeviceIdentityResultPlatform = "Windows"
	ImportedDeviceIdentityResultPlatformwindowsMobile ImportedDeviceIdentityResultPlatform = "WindowsMobile"
)

func PossibleValuesForImportedDeviceIdentityResultPlatform() []string {
	return []string{
		string(ImportedDeviceIdentityResultPlatformandroid),
		string(ImportedDeviceIdentityResultPlatformios),
		string(ImportedDeviceIdentityResultPlatformmacOS),
		string(ImportedDeviceIdentityResultPlatformunknown),
		string(ImportedDeviceIdentityResultPlatformwindows),
		string(ImportedDeviceIdentityResultPlatformwindowsMobile),
	}
}

func parseImportedDeviceIdentityResultPlatform(input string) (*ImportedDeviceIdentityResultPlatform, error) {
	vals := map[string]ImportedDeviceIdentityResultPlatform{
		"android":       ImportedDeviceIdentityResultPlatformandroid,
		"ios":           ImportedDeviceIdentityResultPlatformios,
		"macos":         ImportedDeviceIdentityResultPlatformmacOS,
		"unknown":       ImportedDeviceIdentityResultPlatformunknown,
		"windows":       ImportedDeviceIdentityResultPlatformwindows,
		"windowsmobile": ImportedDeviceIdentityResultPlatformwindowsMobile,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ImportedDeviceIdentityResultPlatform(input)
	return &out, nil
}
