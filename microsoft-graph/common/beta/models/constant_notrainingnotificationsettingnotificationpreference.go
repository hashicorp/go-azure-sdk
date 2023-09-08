package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingNotificationSettingNotificationPreference string

const (
	NoTrainingNotificationSettingNotificationPreferencecustom    NoTrainingNotificationSettingNotificationPreference = "Custom"
	NoTrainingNotificationSettingNotificationPreferencemicrosoft NoTrainingNotificationSettingNotificationPreference = "Microsoft"
	NoTrainingNotificationSettingNotificationPreferenceunknown   NoTrainingNotificationSettingNotificationPreference = "Unknown"
)

func PossibleValuesForNoTrainingNotificationSettingNotificationPreference() []string {
	return []string{
		string(NoTrainingNotificationSettingNotificationPreferencecustom),
		string(NoTrainingNotificationSettingNotificationPreferencemicrosoft),
		string(NoTrainingNotificationSettingNotificationPreferenceunknown),
	}
}

func parseNoTrainingNotificationSettingNotificationPreference(input string) (*NoTrainingNotificationSettingNotificationPreference, error) {
	vals := map[string]NoTrainingNotificationSettingNotificationPreference{
		"custom":    NoTrainingNotificationSettingNotificationPreferencecustom,
		"microsoft": NoTrainingNotificationSettingNotificationPreferencemicrosoft,
		"unknown":   NoTrainingNotificationSettingNotificationPreferenceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NoTrainingNotificationSettingNotificationPreference(input)
	return &out, nil
}
