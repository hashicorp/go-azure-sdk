package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles string

const (
	MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_Disabled      MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_Enabled       MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_NotConfigured MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles(input string) (*MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicySystemAdminFiles(input)
	return &out, nil
}
