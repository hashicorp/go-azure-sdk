package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewRecurrenceSettingsOperationPredicate struct {
	DurationInDays    *int64
	ODataType         *string
	RecurrenceCount   *int64
	RecurrenceEndType *string
	RecurrenceType    *string
}

func (p AccessReviewRecurrenceSettingsOperationPredicate) Matches(input AccessReviewRecurrenceSettings) bool {

	if p.DurationInDays != nil && (input.DurationInDays == nil || *p.DurationInDays != *input.DurationInDays) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RecurrenceCount != nil && (input.RecurrenceCount == nil || *p.RecurrenceCount != *input.RecurrenceCount) {
		return false
	}

	if p.RecurrenceEndType != nil && (input.RecurrenceEndType == nil || *p.RecurrenceEndType != *input.RecurrenceEndType) {
		return false
	}

	if p.RecurrenceType != nil && (input.RecurrenceType == nil || *p.RecurrenceType != *input.RecurrenceType) {
		return false
	}

	return true
}
