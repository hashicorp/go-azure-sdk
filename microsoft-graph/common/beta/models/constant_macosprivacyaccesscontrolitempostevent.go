package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemPostEvent string

const (
	MacOSPrivacyAccessControlItemPostEventdisabled      MacOSPrivacyAccessControlItemPostEvent = "Disabled"
	MacOSPrivacyAccessControlItemPostEventenabled       MacOSPrivacyAccessControlItemPostEvent = "Enabled"
	MacOSPrivacyAccessControlItemPostEventnotConfigured MacOSPrivacyAccessControlItemPostEvent = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemPostEvent() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemPostEventdisabled),
		string(MacOSPrivacyAccessControlItemPostEventenabled),
		string(MacOSPrivacyAccessControlItemPostEventnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemPostEvent(input string) (*MacOSPrivacyAccessControlItemPostEvent, error) {
	vals := map[string]MacOSPrivacyAccessControlItemPostEvent{
		"disabled":      MacOSPrivacyAccessControlItemPostEventdisabled,
		"enabled":       MacOSPrivacyAccessControlItemPostEventenabled,
		"notconfigured": MacOSPrivacyAccessControlItemPostEventnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemPostEvent(input)
	return &out, nil
}
