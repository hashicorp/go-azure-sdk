package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleUpdateProperties struct {
	Notes             *string            `json:"notes,omitempty"`
	RecurrencePattern *RecurrencePattern `json:"recurrencePattern,omitempty"`
	StartAt           *string            `json:"startAt,omitempty"`
	StopAt            *string            `json:"stopAt,omitempty"`
	TimeZoneId        *string            `json:"timeZoneId,omitempty"`
}
