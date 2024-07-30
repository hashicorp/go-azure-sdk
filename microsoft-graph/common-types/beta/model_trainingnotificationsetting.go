package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TrainingNotificationSetting struct {
	NotificationPreference *TrainingNotificationSettingNotificationPreference `json:"notificationPreference,omitempty"`
	ODataType              *string                                            `json:"@odata.type,omitempty"`
	PositiveReinforcement  *PositiveReinforcementNotification                 `json:"positiveReinforcement,omitempty"`
	SettingType            *TrainingNotificationSettingSettingType            `json:"settingType,omitempty"`
	TrainingAssignment     *BaseEndUserNotification                           `json:"trainingAssignment,omitempty"`
	TrainingReminder       *TrainingReminderNotification                      `json:"trainingReminder,omitempty"`
}
