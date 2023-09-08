package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationNotificationType string

const (
	EndUserNotificationNotificationTypenoTraining            EndUserNotificationNotificationType = "NoTraining"
	EndUserNotificationNotificationTypepositiveReinforcement EndUserNotificationNotificationType = "PositiveReinforcement"
	EndUserNotificationNotificationTypetrainingAssignment    EndUserNotificationNotificationType = "TrainingAssignment"
	EndUserNotificationNotificationTypetrainingReminder      EndUserNotificationNotificationType = "TrainingReminder"
	EndUserNotificationNotificationTypeunknown               EndUserNotificationNotificationType = "Unknown"
)

func PossibleValuesForEndUserNotificationNotificationType() []string {
	return []string{
		string(EndUserNotificationNotificationTypenoTraining),
		string(EndUserNotificationNotificationTypepositiveReinforcement),
		string(EndUserNotificationNotificationTypetrainingAssignment),
		string(EndUserNotificationNotificationTypetrainingReminder),
		string(EndUserNotificationNotificationTypeunknown),
	}
}

func parseEndUserNotificationNotificationType(input string) (*EndUserNotificationNotificationType, error) {
	vals := map[string]EndUserNotificationNotificationType{
		"notraining":            EndUserNotificationNotificationTypenoTraining,
		"positivereinforcement": EndUserNotificationNotificationTypepositiveReinforcement,
		"trainingassignment":    EndUserNotificationNotificationTypetrainingAssignment,
		"trainingreminder":      EndUserNotificationNotificationTypetrainingReminder,
		"unknown":               EndUserNotificationNotificationTypeunknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EndUserNotificationNotificationType(input)
	return &out, nil
}
