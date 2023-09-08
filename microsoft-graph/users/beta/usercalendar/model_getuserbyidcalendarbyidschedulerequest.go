package usercalendar

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GetUserByIdCalendarByIdScheduleRequest struct {
	AvailabilityViewInterval *int64            `json:"AvailabilityViewInterval,omitempty"`
	EndTime                  *DateTimeTimeZone `json:"EndTime,omitempty"`
	Schedules                *[]string         `json:"Schedules,omitempty"`
	StartTime                *DateTimeTimeZone `json:"StartTime,omitempty"`
}
