package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BookingWorkHours struct {
	Day       *BookingWorkHoursDay   `json:"day,omitempty"`
	ODataType *string                `json:"@odata.type,omitempty"`
	TimeSlots *[]BookingWorkTimeSlot `json:"timeSlots,omitempty"`
}
