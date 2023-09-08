package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MeetingParticipants struct {
	Attendees *[]MeetingParticipantInfo `json:"attendees,omitempty"`
	ODataType *string                   `json:"@odata.type,omitempty"`
	Organizer *MeetingParticipantInfo   `json:"organizer,omitempty"`
}
