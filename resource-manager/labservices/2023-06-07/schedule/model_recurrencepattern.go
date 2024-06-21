package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecurrencePattern struct {
	ExpirationDate string              `json:"expirationDate"`
	Frequency      RecurrenceFrequency `json:"frequency"`
	Interval       *int64              `json:"interval,omitempty"`
	WeekDays       *[]WeekDay          `json:"weekDays,omitempty"`
}
