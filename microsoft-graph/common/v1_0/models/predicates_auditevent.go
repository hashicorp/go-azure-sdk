package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AuditEventOperationPredicate struct {
	Activity              *string
	ActivityDateTime      *string
	ActivityOperationType *string
	ActivityResult        *string
	ActivityType          *string
	Category              *string
	ComponentName         *string
	CorrelationId         *string
	DisplayName           *string
	Id                    *string
	ODataType             *string
}

func (p AuditEventOperationPredicate) Matches(input AuditEvent) bool {

	if p.Activity != nil && (input.Activity == nil || *p.Activity != *input.Activity) {
		return false
	}

	if p.ActivityDateTime != nil && (input.ActivityDateTime == nil || *p.ActivityDateTime != *input.ActivityDateTime) {
		return false
	}

	if p.ActivityOperationType != nil && (input.ActivityOperationType == nil || *p.ActivityOperationType != *input.ActivityOperationType) {
		return false
	}

	if p.ActivityResult != nil && (input.ActivityResult == nil || *p.ActivityResult != *input.ActivityResult) {
		return false
	}

	if p.ActivityType != nil && (input.ActivityType == nil || *p.ActivityType != *input.ActivityType) {
		return false
	}

	if p.Category != nil && (input.Category == nil || *p.Category != *input.Category) {
		return false
	}

	if p.ComponentName != nil && (input.ComponentName == nil || *p.ComponentName != *input.ComponentName) {
		return false
	}

	if p.CorrelationId != nil && (input.CorrelationId == nil || *p.CorrelationId != *input.CorrelationId) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
