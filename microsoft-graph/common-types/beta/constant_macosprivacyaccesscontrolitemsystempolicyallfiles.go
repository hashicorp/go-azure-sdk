package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyAllFiles string

const (
	MacOSPrivacyAccessControlItemSystemPolicyAllFiles_Disabled      MacOSPrivacyAccessControlItemSystemPolicyAllFiles = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicyAllFiles_Enabled       MacOSPrivacyAccessControlItemSystemPolicyAllFiles = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicyAllFiles_NotConfigured MacOSPrivacyAccessControlItemSystemPolicyAllFiles = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyAllFiles() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyAllFiles_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyAllFiles_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyAllFiles_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicyAllFiles) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicyAllFiles(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicyAllFiles(input string) (*MacOSPrivacyAccessControlItemSystemPolicyAllFiles, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyAllFiles{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyAllFiles_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyAllFiles_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyAllFiles_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyAllFiles(input)
	return &out, nil
}
