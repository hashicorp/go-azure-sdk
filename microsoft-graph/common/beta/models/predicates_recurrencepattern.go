package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePatternOperationPredicate struct {
	DayOfMonth *int64
	Interval   *int64
	Month      *int64
	ODataType  *string
}

func (p RecurrencePatternOperationPredicate) Matches(input RecurrencePattern) bool {

	if p.DayOfMonth != nil && (input.DayOfMonth == nil || *p.DayOfMonth != *input.DayOfMonth) {
		return false
	}

	if p.Interval != nil && (input.Interval == nil || *p.Interval != *input.Interval) {
		return false
	}

	if p.Month != nil && (input.Month == nil || *p.Month != *input.Month) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
