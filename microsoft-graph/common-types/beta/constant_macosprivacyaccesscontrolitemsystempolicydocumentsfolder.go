package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder string

const (
	MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_Disabled      MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_Enabled       MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_NotConfigured MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder(input string) (*MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyDocumentsFolder(input)
	return &out, nil
}
