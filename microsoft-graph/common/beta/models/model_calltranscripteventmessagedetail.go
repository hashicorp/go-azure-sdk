package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscriptEventMessageDetail struct {
	CallId                *string      `json:"callId,omitempty"`
	CallTranscriptICalUid *string      `json:"callTranscriptICalUid,omitempty"`
	MeetingOrganizer      *IdentitySet `json:"meetingOrganizer,omitempty"`
	ODataType             *string      `json:"@odata.type,omitempty"`
}
