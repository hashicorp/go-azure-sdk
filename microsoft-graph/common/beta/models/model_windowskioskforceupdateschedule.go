package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskForceUpdateSchedule struct {
	DayofMonth                         *int64                                     `json:"dayofMonth,omitempty"`
	DayofWeek                          *WindowsKioskForceUpdateScheduleDayofWeek  `json:"dayofWeek,omitempty"`
	ODataType                          *string                                    `json:"@odata.type,omitempty"`
	Recurrence                         *WindowsKioskForceUpdateScheduleRecurrence `json:"recurrence,omitempty"`
	RunImmediatelyIfAfterStartDateTime *bool                                      `json:"runImmediatelyIfAfterStartDateTime,omitempty"`
	StartDateTime                      *string                                    `json:"startDateTime,omitempty"`
}
