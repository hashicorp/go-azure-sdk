package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemAccessibility string

const (
	MacOSPrivacyAccessControlItemAccessibility_Disabled      MacOSPrivacyAccessControlItemAccessibility = "disabled"
	MacOSPrivacyAccessControlItemAccessibility_Enabled       MacOSPrivacyAccessControlItemAccessibility = "enabled"
	MacOSPrivacyAccessControlItemAccessibility_NotConfigured MacOSPrivacyAccessControlItemAccessibility = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemAccessibility() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemAccessibility_Disabled),
		string(MacOSPrivacyAccessControlItemAccessibility_Enabled),
		string(MacOSPrivacyAccessControlItemAccessibility_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemAccessibility) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemAccessibility(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemAccessibility(input string) (*MacOSPrivacyAccessControlItemAccessibility, error) {
	vals := map[string]MacOSPrivacyAccessControlItemAccessibility{
		"disabled":      MacOSPrivacyAccessControlItemAccessibility_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemAccessibility_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemAccessibility_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemAccessibility(input)
	return &out, nil
}
