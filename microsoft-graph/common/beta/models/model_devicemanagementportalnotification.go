package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementPortalNotification struct {
	AlertImpact              *DeviceManagementAlertImpact                         `json:"alertImpact,omitempty"`
	AlertRecordId            *string                                              `json:"alertRecordId,omitempty"`
	AlertRuleId              *string                                              `json:"alertRuleId,omitempty"`
	AlertRuleName            *string                                              `json:"alertRuleName,omitempty"`
	AlertRuleTemplate        *DeviceManagementPortalNotificationAlertRuleTemplate `json:"alertRuleTemplate,omitempty"`
	Id                       *string                                              `json:"id,omitempty"`
	IsPortalNotificationSent *bool                                                `json:"isPortalNotificationSent,omitempty"`
	ODataType                *string                                              `json:"@odata.type,omitempty"`
	Severity                 *DeviceManagementPortalNotificationSeverity          `json:"severity,omitempty"`
}
