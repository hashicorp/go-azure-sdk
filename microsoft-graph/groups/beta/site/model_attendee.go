package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Attendee struct {
	EmailAddress    *EmailAddress   `json:"emailAddress,omitempty"`
	ODataType       *string         `json:"@odata.type,omitempty"`
	ProposedNewTime *TimeSlot       `json:"proposedNewTime,omitempty"`
	Status          *ResponseStatus `json:"status,omitempty"`
	Type            *AttendeeType   `json:"type,omitempty"`
}
