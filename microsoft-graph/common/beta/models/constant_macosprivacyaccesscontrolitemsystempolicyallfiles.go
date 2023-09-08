package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyAllFiles string

const (
	MacOSPrivacyAccessControlItemSystemPolicyAllFilesdisabled      MacOSPrivacyAccessControlItemSystemPolicyAllFiles = "Disabled"
	MacOSPrivacyAccessControlItemSystemPolicyAllFilesenabled       MacOSPrivacyAccessControlItemSystemPolicyAllFiles = "Enabled"
	MacOSPrivacyAccessControlItemSystemPolicyAllFilesnotConfigured MacOSPrivacyAccessControlItemSystemPolicyAllFiles = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyAllFiles() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyAllFilesdisabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyAllFilesenabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyAllFilesnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSystemPolicyAllFiles(input string) (*MacOSPrivacyAccessControlItemSystemPolicyAllFiles, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyAllFiles{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyAllFilesdisabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyAllFilesenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyAllFilesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyAllFiles(input)
	return &out, nil
}
