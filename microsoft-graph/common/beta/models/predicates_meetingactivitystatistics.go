package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingActivityStatisticsOperationPredicate struct {
	AfterHours   *string
	Conflicting  *string
	Duration     *string
	EndDate      *string
	Id           *string
	Long         *string
	Multitasking *string
	ODataType    *string
	Organized    *string
	Recurring    *string
	StartDate    *string
	TimeZoneUsed *string
}

func (p MeetingActivityStatisticsOperationPredicate) Matches(input MeetingActivityStatistics) bool {

	if p.AfterHours != nil && (input.AfterHours == nil || *p.AfterHours != *input.AfterHours) {
		return false
	}

	if p.Conflicting != nil && (input.Conflicting == nil || *p.Conflicting != *input.Conflicting) {
		return false
	}

	if p.Duration != nil && (input.Duration == nil || *p.Duration != *input.Duration) {
		return false
	}

	if p.EndDate != nil && (input.EndDate == nil || *p.EndDate != *input.EndDate) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Long != nil && (input.Long == nil || *p.Long != *input.Long) {
		return false
	}

	if p.Multitasking != nil && (input.Multitasking == nil || *p.Multitasking != *input.Multitasking) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Organized != nil && (input.Organized == nil || *p.Organized != *input.Organized) {
		return false
	}

	if p.Recurring != nil && (input.Recurring == nil || *p.Recurring != *input.Recurring) {
		return false
	}

	if p.StartDate != nil && (input.StartDate == nil || *p.StartDate != *input.StartDate) {
		return false
	}

	if p.TimeZoneUsed != nil && (input.TimeZoneUsed == nil || *p.TimeZoneUsed != *input.TimeZoneUsed) {
		return false
	}

	return true
}
