package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyNotificationRuleOperationPredicate struct {
	Id                         *string
	IsDefaultRecipientsEnabled *bool
	NotificationLevel          *string
	NotificationType           *string
	ODataType                  *string
	RecipientType              *string
}

func (p UnifiedRoleManagementPolicyNotificationRuleOperationPredicate) Matches(input UnifiedRoleManagementPolicyNotificationRule) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefaultRecipientsEnabled != nil && (input.IsDefaultRecipientsEnabled == nil || *p.IsDefaultRecipientsEnabled != *input.IsDefaultRecipientsEnabled) {
		return false
	}

	if p.NotificationLevel != nil && (input.NotificationLevel == nil || *p.NotificationLevel != *input.NotificationLevel) {
		return false
	}

	if p.NotificationType != nil && (input.NotificationType == nil || *p.NotificationType != *input.NotificationType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecipientType != nil && (input.RecipientType == nil || *p.RecipientType != *input.RecipientType) {
		return false
	}

	return true
}
