package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSettingSettingType string

const (
	EndUserNotificationSettingSettingTypenoNotification   EndUserNotificationSettingSettingType = "NoNotification"
	EndUserNotificationSettingSettingTypenoTraining       EndUserNotificationSettingSettingType = "NoTraining"
	EndUserNotificationSettingSettingTypetrainingSelected EndUserNotificationSettingSettingType = "TrainingSelected"
	EndUserNotificationSettingSettingTypeunknown          EndUserNotificationSettingSettingType = "Unknown"
)

func PossibleValuesForEndUserNotificationSettingSettingType() []string {
	return []string{
		string(EndUserNotificationSettingSettingTypenoNotification),
		string(EndUserNotificationSettingSettingTypenoTraining),
		string(EndUserNotificationSettingSettingTypetrainingSelected),
		string(EndUserNotificationSettingSettingTypeunknown),
	}
}

func parseEndUserNotificationSettingSettingType(input string) (*EndUserNotificationSettingSettingType, error) {
	vals := map[string]EndUserNotificationSettingSettingType{
		"nonotification":   EndUserNotificationSettingSettingTypenoNotification,
		"notraining":       EndUserNotificationSettingSettingTypenoTraining,
		"trainingselected": EndUserNotificationSettingSettingTypetrainingSelected,
		"unknown":          EndUserNotificationSettingSettingTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationSettingSettingType(input)
	return &out, nil
}
