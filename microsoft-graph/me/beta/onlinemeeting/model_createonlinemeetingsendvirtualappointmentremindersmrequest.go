package onlinemeeting

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateOnlineMeetingSendVirtualAppointmentReminderSmRequest struct {
	PhoneNumbers                  *[]string                                                                                `json:"phoneNumbers,omitempty"`
	RemindBeforeTimeInMinutesType *CreateOnlineMeetingSendVirtualAppointmentReminderSmRequestRemindBeforeTimeInMinutesType `json:"remindBeforeTimeInMinutesType,omitempty"`
}
