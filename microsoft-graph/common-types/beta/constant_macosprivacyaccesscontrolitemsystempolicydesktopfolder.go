package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder string

const (
	MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_Disabled      MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_Enabled       MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_NotConfigured MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyDesktopFolder() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicyDesktopFolder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicyDesktopFolder(input string) (*MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyDesktopFolder(input)
	return &out, nil
}
