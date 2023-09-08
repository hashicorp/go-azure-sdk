package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder string

const (
	MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolderdisabled      MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder = "Disabled"
	MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolderenabled       MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder = "Enabled"
	MacOSPrivacyAccessControlItemSystemPolicyDocumentsFoldernotConfigured MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolderdisabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolderenabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDocumentsFoldernotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder(input string) (*MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolderdisabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolderenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyDocumentsFoldernotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder(input)
	return &out, nil
}
