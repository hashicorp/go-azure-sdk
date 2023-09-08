package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder string

const (
	MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolderdisabled      MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder = "Disabled"
	MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolderenabled       MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder = "Enabled"
	MacOSPrivacyAccessControlItemSystemPolicyDownloadsFoldernotConfigured MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolderdisabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolderenabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDownloadsFoldernotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder(input string) (*MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolderdisabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolderenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyDownloadsFoldernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder(input)
	return &out, nil
}
