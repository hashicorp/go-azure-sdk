package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles string

const (
	MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesdisabled      MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles = "Disabled"
	MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesenabled       MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles = "Enabled"
	MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesnotConfigured MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesdisabled),
		string(MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesenabled),
		string(MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles(input string) (*MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesdisabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicySystemAdminFilesnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles(input)
	return &out, nil
}
