package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscript struct {
	Content              *string      `json:"content,omitempty"`
	CreatedDateTime      *string      `json:"createdDateTime,omitempty"`
	Id                   *string      `json:"id,omitempty"`
	MeetingId            *string      `json:"meetingId,omitempty"`
	MeetingOrganizer     *IdentitySet `json:"meetingOrganizer,omitempty"`
	MeetingOrganizerId   *string      `json:"meetingOrganizerId,omitempty"`
	MetadataContent      *string      `json:"metadataContent,omitempty"`
	ODataType            *string      `json:"@odata.type,omitempty"`
	TranscriptContentUrl *string      `json:"transcriptContentUrl,omitempty"`
}
