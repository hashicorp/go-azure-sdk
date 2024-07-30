package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemMediaLibrary string

const (
	MacOSPrivacyAccessControlItemMediaLibrary_Disabled      MacOSPrivacyAccessControlItemMediaLibrary = "disabled"
	MacOSPrivacyAccessControlItemMediaLibrary_Enabled       MacOSPrivacyAccessControlItemMediaLibrary = "enabled"
	MacOSPrivacyAccessControlItemMediaLibrary_NotConfigured MacOSPrivacyAccessControlItemMediaLibrary = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemMediaLibrary() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemMediaLibrary_Disabled),
		string(MacOSPrivacyAccessControlItemMediaLibrary_Enabled),
		string(MacOSPrivacyAccessControlItemMediaLibrary_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemMediaLibrary) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemMediaLibrary(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemMediaLibrary(input string) (*MacOSPrivacyAccessControlItemMediaLibrary, error) {
	vals := map[string]MacOSPrivacyAccessControlItemMediaLibrary{
		"disabled":      MacOSPrivacyAccessControlItemMediaLibrary_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemMediaLibrary_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemMediaLibrary_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemMediaLibrary(input)
	return &out, nil
}
