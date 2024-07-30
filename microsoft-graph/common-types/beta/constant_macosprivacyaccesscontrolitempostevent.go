package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemPostEvent string

const (
	MacOSPrivacyAccessControlItemPostEvent_Disabled      MacOSPrivacyAccessControlItemPostEvent = "disabled"
	MacOSPrivacyAccessControlItemPostEvent_Enabled       MacOSPrivacyAccessControlItemPostEvent = "enabled"
	MacOSPrivacyAccessControlItemPostEvent_NotConfigured MacOSPrivacyAccessControlItemPostEvent = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemPostEvent() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemPostEvent_Disabled),
		string(MacOSPrivacyAccessControlItemPostEvent_Enabled),
		string(MacOSPrivacyAccessControlItemPostEvent_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemPostEvent) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemPostEvent(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemPostEvent(input string) (*MacOSPrivacyAccessControlItemPostEvent, error) {
	vals := map[string]MacOSPrivacyAccessControlItemPostEvent{
		"disabled":      MacOSPrivacyAccessControlItemPostEvent_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemPostEvent_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemPostEvent_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemPostEvent(input)
	return &out, nil
}
