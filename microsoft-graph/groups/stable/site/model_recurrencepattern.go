package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePattern struct {
	DayOfMonth     *int64                           `json:"dayOfMonth,omitempty"`
	DaysOfWeek     *RecurrencePatternDaysOfWeek     `json:"daysOfWeek,omitempty"`
	FirstDayOfWeek *RecurrencePatternFirstDayOfWeek `json:"firstDayOfWeek,omitempty"`
	Index          *RecurrencePatternIndex          `json:"index,omitempty"`
	Interval       *int64                           `json:"interval,omitempty"`
	Month          *int64                           `json:"month,omitempty"`
	ODataType      *string                          `json:"@odata.type,omitempty"`
	Type           *RecurrencePatternType           `json:"type,omitempty"`
}
