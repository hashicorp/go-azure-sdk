package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCAuditEventOperationPredicate struct {
	Activity         *string
	ActivityDateTime *string
	ActivityType     *string
	ComponentName    *string
	CorrelationId    *string
	DisplayName      *string
	Id               *string
	ODataType        *string
}

func (p CloudPCAuditEventOperationPredicate) Matches(input CloudPCAuditEvent) bool {

	if p.Activity != nil && (input.Activity == nil || *p.Activity != *input.Activity) {
		return false
	}

	if p.ActivityDateTime != nil && (input.ActivityDateTime == nil || *p.ActivityDateTime != *input.ActivityDateTime) {
		return false
	}

	if p.ActivityType != nil && (input.ActivityType == nil || *p.ActivityType != *input.ActivityType) {
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
