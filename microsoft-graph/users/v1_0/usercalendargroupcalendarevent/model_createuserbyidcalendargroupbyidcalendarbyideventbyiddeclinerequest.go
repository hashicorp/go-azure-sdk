package usercalendargroupcalendarevent

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateUserByIdCalendarGroupByIdCalendarByIdEventByIdDeclineRequest struct {
	Comment         *string   `json:"Comment,omitempty"`
	ProposedNewTime *TimeSlot `json:"ProposedNewTime,omitempty"`
	SendResponse    *bool     `json:"SendResponse,omitempty"`
}
