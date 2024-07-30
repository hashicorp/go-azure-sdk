package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyNotificationRule struct {
	Id                         *string                                `json:"id,omitempty"`
	IsDefaultRecipientsEnabled *bool                                  `json:"isDefaultRecipientsEnabled,omitempty"`
	NotificationLevel          *string                                `json:"notificationLevel,omitempty"`
	NotificationRecipients     *[]string                              `json:"notificationRecipients,omitempty"`
	NotificationType           *string                                `json:"notificationType,omitempty"`
	ODataType                  *string                                `json:"@odata.type,omitempty"`
	RecipientType              *string                                `json:"recipientType,omitempty"`
	Target                     *UnifiedRoleManagementPolicyRuleTarget `json:"target,omitempty"`
}
