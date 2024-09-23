package onlinemeeting

import (
	"github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SendOnlineMeetingVirtualAppointmentSmsRequest struct {
	Attendees   *[]beta.AttendeeNotificationInfo    `json:"attendees,omitempty"`
	MessageType *beta.VirtualAppointmentMessageType `json:"messageType,omitempty"`
}
