package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecording struct {
	Content             *string `json:"content,omitempty"`
	CreatedDateTime     *string `json:"createdDateTime,omitempty"`
	Id                  *string `json:"id,omitempty"`
	MeetingId           *string `json:"meetingId,omitempty"`
	MeetingOrganizerId  *string `json:"meetingOrganizerId,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	RecordingContentUrl *string `json:"recordingContentUrl,omitempty"`
}
