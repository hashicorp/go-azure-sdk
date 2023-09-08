package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CustomSecurityAttributeAuditOperationPredicate struct {
	ActivityDateTime    *string
	ActivityDisplayName *string
	Category            *string
	CorrelationId       *string
	Id                  *string
	LoggedByService     *string
	ODataType           *string
	OperationType       *string
	ResultReason        *string
	UserAgent           *string
}

func (p CustomSecurityAttributeAuditOperationPredicate) Matches(input CustomSecurityAttributeAudit) bool {

	if p.ActivityDateTime != nil && (input.ActivityDateTime == nil || *p.ActivityDateTime != *input.ActivityDateTime) {
		return false
	}

	if p.ActivityDisplayName != nil && (input.ActivityDisplayName == nil || *p.ActivityDisplayName != *input.ActivityDisplayName) {
		return false
	}

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.CorrelationId != nil && (input.CorrelationId == nil || *p.CorrelationId != *input.CorrelationId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LoggedByService != nil && (input.LoggedByService == nil || *p.LoggedByService != *input.LoggedByService) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OperationType != nil && (input.OperationType == nil || *p.OperationType != *input.OperationType) {
		return false
	}

	if p.ResultReason != nil && (input.ResultReason == nil || *p.ResultReason != *input.ResultReason) {
		return false
	}

	if p.UserAgent != nil && (input.UserAgent == nil || *p.UserAgent != *input.UserAgent) {
		return false
	}

	return true
}
