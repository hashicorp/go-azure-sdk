package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TeamworkOnlineMeetingInfo struct {
	CalendarEventId *string               `json:"calendarEventId,omitempty"`
	JoinWebUrl      *string               `json:"joinWebUrl,omitempty"`
	ODataType       *string               `json:"@odata.type,omitempty"`
	Organizer       *TeamworkUserIdentity `json:"organizer,omitempty"`
}
