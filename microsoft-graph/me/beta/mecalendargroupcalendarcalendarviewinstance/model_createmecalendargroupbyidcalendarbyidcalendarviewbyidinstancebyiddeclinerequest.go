package mecalendargroupcalendarcalendarviewinstance

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreateMeCalendarGroupByIdCalendarByIdCalendarViewByIdInstanceByIdDeclineRequest struct {
	Comment         *string   `json:"Comment,omitempty"`
	ProposedNewTime *TimeSlot `json:"ProposedNewTime,omitempty"`
	SendResponse    *bool     `json:"SendResponse,omitempty"`
}
