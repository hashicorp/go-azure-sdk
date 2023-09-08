package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItemReminders string

const (
	MacOSPrivacyAccessControlItemRemindersdisabled      MacOSPrivacyAccessControlItemReminders = "Disabled"
	MacOSPrivacyAccessControlItemRemindersenabled       MacOSPrivacyAccessControlItemReminders = "Enabled"
	MacOSPrivacyAccessControlItemRemindersnotConfigured MacOSPrivacyAccessControlItemReminders = "NotConfigured"
)

func PossibleValuesForMacOSPrivacyAccessControlItemReminders() []string {
	return []string{
		string(MacOSPrivacyAccessControlItemRemindersdisabled),
		string(MacOSPrivacyAccessControlItemRemindersenabled),
		string(MacOSPrivacyAccessControlItemRemindersnotConfigured),
	}
}

func parseMacOSPrivacyAccessControlItemReminders(input string) (*MacOSPrivacyAccessControlItemReminders, error) {
	vals := map[string]MacOSPrivacyAccessControlItemReminders{
		"disabled":      MacOSPrivacyAccessControlItemRemindersdisabled,
		"enabled":       MacOSPrivacyAccessControlItemRemindersenabled,
		"notconfigured": MacOSPrivacyAccessControlItemRemindersnotConfigured,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := MacOSPrivacyAccessControlItemReminders(input)
	return &out, nil
}
