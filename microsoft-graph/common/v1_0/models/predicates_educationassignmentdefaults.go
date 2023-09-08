package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationAssignmentDefaultsOperationPredicate struct {
	DueTime                *string
	Id                     *string
	NotificationChannelUrl *string
	ODataType              *string
}

func (p EducationAssignmentDefaultsOperationPredicate) Matches(input EducationAssignmentDefaults) bool {

	if p.DueTime != nil && (input.DueTime == nil || *p.DueTime != *input.DueTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.NotificationChannelUrl != nil && (input.NotificationChannelUrl == nil || *p.NotificationChannelUrl != *input.NotificationChannelUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
