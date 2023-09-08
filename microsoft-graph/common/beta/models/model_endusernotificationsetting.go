package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EndUserNotificationSetting struct {
	NotificationPreference *EndUserNotificationSettingNotificationPreference `json:"notificationPreference,omitempty"`
	ODataType              *string                                           `json:"@odata.type,omitempty"`
	PositiveReinforcement  *PositiveReinforcementNotification                `json:"positiveReinforcement,omitempty"`
	SettingType            *EndUserNotificationSettingSettingType            `json:"settingType,omitempty"`
}
