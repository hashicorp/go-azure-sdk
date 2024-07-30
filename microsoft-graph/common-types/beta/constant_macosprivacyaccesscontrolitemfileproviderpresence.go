package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemFileProviderPresence string

const (
	MacOSPrivacyAccessControlItemFileProviderPresence_Disabled      MacOSPrivacyAccessControlItemFileProviderPresence = "disabled"
	MacOSPrivacyAccessControlItemFileProviderPresence_Enabled       MacOSPrivacyAccessControlItemFileProviderPresence = "enabled"
	MacOSPrivacyAccessControlItemFileProviderPresence_NotConfigured MacOSPrivacyAccessControlItemFileProviderPresence = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemFileProviderPresence() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemFileProviderPresence_Disabled),
		string(MacOSPrivacyAccessControlItemFileProviderPresence_Enabled),
		string(MacOSPrivacyAccessControlItemFileProviderPresence_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemFileProviderPresence) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemFileProviderPresence(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemFileProviderPresence(input string) (*MacOSPrivacyAccessControlItemFileProviderPresence, error) {
	vals := map[string]MacOSPrivacyAccessControlItemFileProviderPresence{
		"disabled":      MacOSPrivacyAccessControlItemFileProviderPresence_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemFileProviderPresence_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemFileProviderPresence_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemFileProviderPresence(input)
	return &out, nil
}
