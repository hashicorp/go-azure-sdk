package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StandardTimeZoneOffsetOperationPredicate struct {
	DayOccurrence *int64
	Month         *int64
	ODataType     *string
	Time          *string
	Year          *int64
}

func (p StandardTimeZoneOffsetOperationPredicate) Matches(input StandardTimeZoneOffset) bool {

	if p.DayOccurrence != nil && (input.DayOccurrence == nil || *p.DayOccurrence != *input.DayOccurrence) {
		return false
	}

	if p.Month != nil && (input.Month == nil || *p.Month != *input.Month) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Time != nil && (input.Time == nil || *p.Time != *input.Time) {
		return false
	}

	if p.Year != nil && (input.Year == nil || *p.Year != *input.Year) {
		return false
	}

	return true
}
