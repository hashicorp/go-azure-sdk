package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemCalendar string

const (
	MacOSPrivacyAccessControlItemCalendardisabled      MacOSPrivacyAccessControlItemCalendar = "Disabled"
	MacOSPrivacyAccessControlItemCalendarenabled       MacOSPrivacyAccessControlItemCalendar = "Enabled"
	MacOSPrivacyAccessControlItemCalendarnotConfigured MacOSPrivacyAccessControlItemCalendar = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemCalendar() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemCalendardisabled),
		string(MacOSPrivacyAccessControlItemCalendarenabled),
		string(MacOSPrivacyAccessControlItemCalendarnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemCalendar(input string) (*MacOSPrivacyAccessControlItemCalendar, error) {
	vals := map[string]MacOSPrivacyAccessControlItemCalendar{
		"disabled":      MacOSPrivacyAccessControlItemCalendardisabled,
		"enabled":       MacOSPrivacyAccessControlItemCalendarenabled,
		"notconfigured": MacOSPrivacyAccessControlItemCalendarnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemCalendar(input)
	return &out, nil
}
