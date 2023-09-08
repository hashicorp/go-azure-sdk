package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingNotificationSettingNotificationPreference string

const (
	TrainingNotificationSettingNotificationPreferencecustom    TrainingNotificationSettingNotificationPreference = "Custom"
	TrainingNotificationSettingNotificationPreferencemicrosoft TrainingNotificationSettingNotificationPreference = "Microsoft"
	TrainingNotificationSettingNotificationPreferenceunknown   TrainingNotificationSettingNotificationPreference = "Unknown"
)

func PossibleValuesForTrainingNotificationSettingNotificationPreference() []string {
	return []string{
		string(TrainingNotificationSettingNotificationPreferencecustom),
		string(TrainingNotificationSettingNotificationPreferencemicrosoft),
		string(TrainingNotificationSettingNotificationPreferenceunknown),
	}
}

func parseTrainingNotificationSettingNotificationPreference(input string) (*TrainingNotificationSettingNotificationPreference, error) {
	vals := map[string]TrainingNotificationSettingNotificationPreference{
		"custom":    TrainingNotificationSettingNotificationPreferencecustom,
		"microsoft": TrainingNotificationSettingNotificationPreferencemicrosoft,
		"unknown":   TrainingNotificationSettingNotificationPreferenceunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingNotificationSettingNotificationPreference(input)
	return &out, nil
}
