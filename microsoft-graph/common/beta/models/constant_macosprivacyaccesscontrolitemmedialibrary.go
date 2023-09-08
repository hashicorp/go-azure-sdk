package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemMediaLibrary string

const (
	MacOSPrivacyAccessControlItemMediaLibrarydisabled      MacOSPrivacyAccessControlItemMediaLibrary = "Disabled"
	MacOSPrivacyAccessControlItemMediaLibraryenabled       MacOSPrivacyAccessControlItemMediaLibrary = "Enabled"
	MacOSPrivacyAccessControlItemMediaLibrarynotConfigured MacOSPrivacyAccessControlItemMediaLibrary = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemMediaLibrary() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemMediaLibrarydisabled),
		string(MacOSPrivacyAccessControlItemMediaLibraryenabled),
		string(MacOSPrivacyAccessControlItemMediaLibrarynotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemMediaLibrary(input string) (*MacOSPrivacyAccessControlItemMediaLibrary, error) {
	vals := map[string]MacOSPrivacyAccessControlItemMediaLibrary{
		"disabled":      MacOSPrivacyAccessControlItemMediaLibrarydisabled,
		"enabled":       MacOSPrivacyAccessControlItemMediaLibraryenabled,
		"notconfigured": MacOSPrivacyAccessControlItemMediaLibrarynotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemMediaLibrary(input)
	return &out, nil
}
