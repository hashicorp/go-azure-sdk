package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NoTrainingNotificationSettingSettingType string

const (
	NoTrainingNotificationSettingSettingTypenoNotification   NoTrainingNotificationSettingSettingType = "NoNotification"
	NoTrainingNotificationSettingSettingTypenoTraining       NoTrainingNotificationSettingSettingType = "NoTraining"
	NoTrainingNotificationSettingSettingTypetrainingSelected NoTrainingNotificationSettingSettingType = "TrainingSelected"
	NoTrainingNotificationSettingSettingTypeunknown          NoTrainingNotificationSettingSettingType = "Unknown"
)

func PossibleValuesForNoTrainingNotificationSettingSettingType() []string {
	return []string{
		string(NoTrainingNotificationSettingSettingTypenoNotification),
		string(NoTrainingNotificationSettingSettingTypenoTraining),
		string(NoTrainingNotificationSettingSettingTypetrainingSelected),
		string(NoTrainingNotificationSettingSettingTypeunknown),
	}
}

func parseNoTrainingNotificationSettingSettingType(input string) (*NoTrainingNotificationSettingSettingType, error) {
	vals := map[string]NoTrainingNotificationSettingSettingType{
		"nonotification":   NoTrainingNotificationSettingSettingTypenoNotification,
		"notraining":       NoTrainingNotificationSettingSettingTypenoTraining,
		"trainingselected": NoTrainingNotificationSettingSettingTypetrainingSelected,
		"unknown":          NoTrainingNotificationSettingSettingTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := NoTrainingNotificationSettingSettingType(input)
	return &out, nil
}
