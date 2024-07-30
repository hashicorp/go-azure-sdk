package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder string

const (
	MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_Disabled      MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder = "disabled"
	MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_Enabled       MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder = "enabled"
	MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_NotConfigured MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_Disabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_Enabled),
		string(MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder(input string) (*MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder{
		"disabled":      MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSystemPolicyDownloadsFolder(input)
	return &out, nil
}
