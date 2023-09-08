package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleItemOperationPredicate struct {
	IsPrivate *bool
	Location  *string
	ODataType *string
	Subject   *string
}

func (p ScheduleItemOperationPredicate) Matches(input ScheduleItem) bool {

	if p.IsPrivate != nil && (input.IsPrivate == nil || *p.IsPrivate != *input.IsPrivate) {
		return false
	}

	if p.Location != nil && (input.Location == nil || *p.Location != *input.Location) {
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
