package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemSpeechRecognition string

const (
	MacOSPrivacyAccessControlItemSpeechRecognitiondisabled      MacOSPrivacyAccessControlItemSpeechRecognition = "Disabled"
	MacOSPrivacyAccessControlItemSpeechRecognitionenabled       MacOSPrivacyAccessControlItemSpeechRecognition = "Enabled"
	MacOSPrivacyAccessControlItemSpeechRecognitionnotConfigured MacOSPrivacyAccessControlItemSpeechRecognition = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemSpeechRecognition() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemSpeechRecognitiondisabled),
		string(MacOSPrivacyAccessControlItemSpeechRecognitionenabled),
		string(MacOSPrivacyAccessControlItemSpeechRecognitionnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemSpeechRecognition(input string) (*MacOSPrivacyAccessControlItemSpeechRecognition, error) {
	vals := map[string]MacOSPrivacyAccessControlItemSpeechRecognition{
		"disabled":      MacOSPrivacyAccessControlItemSpeechRecognitiondisabled,
		"enabled":       MacOSPrivacyAccessControlItemSpeechRecognitionenabled,
		"notconfigured": MacOSPrivacyAccessControlItemSpeechRecognitionnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemSpeechRecognition(input)
	return &out, nil
}
