package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemPhotos string

const (
	MacOSPrivacyAccessControlItemPhotosdisabled      MacOSPrivacyAccessControlItemPhotos = "Disabled"
	MacOSPrivacyAccessControlItemPhotosenabled       MacOSPrivacyAccessControlItemPhotos = "Enabled"
	MacOSPrivacyAccessControlItemPhotosnotConfigured MacOSPrivacyAccessControlItemPhotos = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemPhotos() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemPhotosdisabled),
		string(MacOSPrivacyAccessControlItemPhotosenabled),
		string(MacOSPrivacyAccessControlItemPhotosnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemPhotos(input string) (*MacOSPrivacyAccessControlItemPhotos, error) {
	vals := map[string]MacOSPrivacyAccessControlItemPhotos{
		"disabled":      MacOSPrivacyAccessControlItemPhotosdisabled,
		"enabled":       MacOSPrivacyAccessControlItemPhotosenabled,
		"notconfigured": MacOSPrivacyAccessControlItemPhotosnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemPhotos(input)
	return &out, nil
}
