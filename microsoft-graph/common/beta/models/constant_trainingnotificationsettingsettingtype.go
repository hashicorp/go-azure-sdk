package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingNotificationSettingSettingType string

const (
	TrainingNotificationSettingSettingTypenoNotification   TrainingNotificationSettingSettingType = "NoNotification"
	TrainingNotificationSettingSettingTypenoTraining       TrainingNotificationSettingSettingType = "NoTraining"
	TrainingNotificationSettingSettingTypetrainingSelected TrainingNotificationSettingSettingType = "TrainingSelected"
	TrainingNotificationSettingSettingTypeunknown          TrainingNotificationSettingSettingType = "Unknown"
)

func PossibleValuesForTrainingNotificationSettingSettingType() []string {
	return []string{
		string(TrainingNotificationSettingSettingTypenoNotification),
		string(TrainingNotificationSettingSettingTypenoTraining),
		string(TrainingNotificationSettingSettingTypetrainingSelected),
		string(TrainingNotificationSettingSettingTypeunknown),
	}
}

func parseTrainingNotificationSettingSettingType(input string) (*TrainingNotificationSettingSettingType, error) {
	vals := map[string]TrainingNotificationSettingSettingType{
		"nonotification":   TrainingNotificationSettingSettingTypenoNotification,
		"notraining":       TrainingNotificationSettingSettingTypenoTraining,
		"trainingselected": TrainingNotificationSettingSettingTypetrainingSelected,
		"unknown":          TrainingNotificationSettingSettingTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := TrainingNotificationSettingSettingType(input)
	return &out, nil
}
