package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemReminders string

const (
	MacOSPrivacyAccessControlItemReminders_Disabled      MacOSPrivacyAccessControlItemReminders = "disabled"
	MacOSPrivacyAccessControlItemReminders_Enabled       MacOSPrivacyAccessControlItemReminders = "enabled"
	MacOSPrivacyAccessControlItemReminders_NotConfigured MacOSPrivacyAccessControlItemReminders = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemReminders() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemReminders_Disabled),
		string(MacOSPrivacyAccessControlItemReminders_Enabled),
		string(MacOSPrivacyAccessControlItemReminders_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemReminders) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemReminders(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemReminders(input string) (*MacOSPrivacyAccessControlItemReminders, error) {
	vals := map[string]MacOSPrivacyAccessControlItemReminders{
		"disabled":      MacOSPrivacyAccessControlItemReminders_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemReminders_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemReminders_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemReminders(input)
	return &out, nil
}
