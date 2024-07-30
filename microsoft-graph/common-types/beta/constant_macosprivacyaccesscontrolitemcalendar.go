package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemCalendar string

const (
	MacOSPrivacyAccessControlItemCalendar_Disabled      MacOSPrivacyAccessControlItemCalendar = "disabled"
	MacOSPrivacyAccessControlItemCalendar_Enabled       MacOSPrivacyAccessControlItemCalendar = "enabled"
	MacOSPrivacyAccessControlItemCalendar_NotConfigured MacOSPrivacyAccessControlItemCalendar = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemCalendar() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemCalendar_Disabled),
		string(MacOSPrivacyAccessControlItemCalendar_Enabled),
		string(MacOSPrivacyAccessControlItemCalendar_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemCalendar) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemCalendar(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemCalendar(input string) (*MacOSPrivacyAccessControlItemCalendar, error) {
	vals := map[string]MacOSPrivacyAccessControlItemCalendar{
		"disabled":      MacOSPrivacyAccessControlItemCalendar_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemCalendar_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemCalendar_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemCalendar(input)
	return &out, nil
}
