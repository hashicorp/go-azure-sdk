package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSpeechRecognition string

const (
	MacOSPrivacyAccessControlItemSpeechRecognition_Disabled      MacOSPrivacyAccessControlItemSpeechRecognition = "disabled"
	MacOSPrivacyAccessControlItemSpeechRecognition_Enabled       MacOSPrivacyAccessControlItemSpeechRecognition = "enabled"
	MacOSPrivacyAccessControlItemSpeechRecognition_NotConfigured MacOSPrivacyAccessControlItemSpeechRecognition = "notConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSpeechRecognition() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSpeechRecognition_Disabled),
		string(MacOSPrivacyAccessControlItemSpeechRecognition_Enabled),
		string(MacOSPrivacyAccessControlItemSpeechRecognition_NotConfigured),
	}
}

func (s *MacOSPrivacyAccessControlItemSpeechRecognition) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseMacOSPrivacyAccessControlItemSpeechRecognition(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseMacOSPrivacyAccessControlItemSpeechRecognition(input string) (*MacOSPrivacyAccessControlItemSpeechRecognition, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSpeechRecognition{
		"disabled":      MacOSPrivacyAccessControlItemSpeechRecognition_Disabled,
		"enabled":       MacOSPrivacyAccessControlItemSpeechRecognition_Enabled,
		"notconfigured": MacOSPrivacyAccessControlItemSpeechRecognition_NotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSpeechRecognition(input)
	return &out, nil
}
