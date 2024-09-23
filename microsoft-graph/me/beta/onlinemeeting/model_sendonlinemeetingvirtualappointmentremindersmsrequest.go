package onlinemeeting

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendOnlineMeetingVirtualAppointmentReminderSmsRequest struct {
	Attendees                     *[]beta.AttendeeNotificationInfo    `json:"attendees,omitempty"`
	RemindBeforeTimeInMinutesType *beta.RemindBeforeTimeInMinutesType `json:"remindBeforeTimeInMinutesType,omitempty"`
}
