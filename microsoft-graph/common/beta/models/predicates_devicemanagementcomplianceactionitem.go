package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementComplianceActionItemOperationPredicate struct {
	GracePeriodHours       *int64
	Id                     *string
	NotificationTemplateId *string
	ODataType              *string
}

func (p DeviceManagementComplianceActionItemOperationPredicate) Matches(input DeviceManagementComplianceActionItem) bool {

	if p.GracePeriodHours != nil && (input.GracePeriodHours == nil || *p.GracePeriodHours != *input.GracePeriodHours) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.NotificationTemplateId != nil && (input.NotificationTemplateId == nil || *p.NotificationTemplateId != *input.NotificationTemplateId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
