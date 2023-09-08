package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsManagedTenantEmailNotificationOperationPredicate struct {
	CreatedByUserId    *string
	CreatedDateTime    *string
	EmailBody          *string
	Id                 *string
	LastActionByUserId *string
	LastActionDateTime *string
	ODataType          *string
	Subject            *string
}

func (p ManagedTenantsManagedTenantEmailNotificationOperationPredicate) Matches(input ManagedTenantsManagedTenantEmailNotification) bool {

	if p.CreatedByUserId != nil && (input.CreatedByUserId == nil || *p.CreatedByUserId != *input.CreatedByUserId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.EmailBody != nil && (input.EmailBody == nil || *p.EmailBody != *input.EmailBody) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastActionByUserId != nil && (input.LastActionByUserId == nil || *p.LastActionByUserId != *input.LastActionByUserId) {
		return false
	}

	if p.LastActionDateTime != nil && (input.LastActionDateTime == nil || *p.LastActionDateTime != *input.LastActionDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Subject != nil && (input.Subject == nil || *p.Subject != *input.Subject) {
		return false
	}

	return true
}
