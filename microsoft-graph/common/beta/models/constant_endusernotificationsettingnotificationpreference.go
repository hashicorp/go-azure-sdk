package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSettingNotificationPreference string

const (
	EndUserNotificationSettingNotificationPreferencecustom    EndUserNotificationSettingNotificationPreference = "Custom"
	EndUserNotificationSettingNotificationPreferencemicrosoft EndUserNotificationSettingNotificationPreference = "Microsoft"
	EndUserNotificationSettingNotificationPreferenceunknown   EndUserNotificationSettingNotificationPreference = "Unknown"
)

func PossibleValuesForEndUserNotificationSettingNotificationPreference() []string {
	return []string{
		string(EndUserNotificationSettingNotificationPreferencecustom),
		string(EndUserNotificationSettingNotificationPreferencemicrosoft),
		string(EndUserNotificationSettingNotificationPreferenceunknown),
	}
}

func parseEndUserNotificationSettingNotificationPreference(input string) (*EndUserNotificationSettingNotificationPreference, error) {
	vals := map[string]EndUserNotificationSettingNotificationPreference{
		"custom":    EndUserNotificationSettingNotificationPreferencecustom,
		"microsoft": EndUserNotificationSettingNotificationPreferencemicrosoft,
		"unknown":   EndUserNotificationSettingNotificationPreferenceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSettingNotificationPreference(input)
	return &out, nil
}
