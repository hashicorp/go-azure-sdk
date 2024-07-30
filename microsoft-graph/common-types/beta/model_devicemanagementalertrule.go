package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementAlertRule struct {
	AlertRuleTemplate    *DeviceManagementAlertRuleAlertRuleTemplate `json:"alertRuleTemplate,omitempty"`
	Description          *string                                     `json:"description,omitempty"`
	DisplayName          *string                                     `json:"displayName,omitempty"`
	Enabled              *bool                                       `json:"enabled,omitempty"`
	Id                   *string                                     `json:"id,omitempty"`
	IsSystemRule         *bool                                       `json:"isSystemRule,omitempty"`
	NotificationChannels *[]DeviceManagementNotificationChannel      `json:"notificationChannels,omitempty"`
	ODataType            *string                                     `json:"@odata.type,omitempty"`
	Severity             *DeviceManagementAlertRuleSeverity          `json:"severity,omitempty"`
	Threshold            *DeviceManagementRuleThreshold              `json:"threshold,omitempty"`
}
