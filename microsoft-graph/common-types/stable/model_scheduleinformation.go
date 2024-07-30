package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleInformation struct {
	AvailabilityView *string         `json:"availabilityView,omitempty"`
	Error            *FreeBusyError  `json:"error,omitempty"`
	ODataType        *string         `json:"@odata.type,omitempty"`
	ScheduleId       *string         `json:"scheduleId,omitempty"`
	ScheduleItems    *[]ScheduleItem `json:"scheduleItems,omitempty"`
	WorkingHours     *WorkingHours   `json:"workingHours,omitempty"`
}
