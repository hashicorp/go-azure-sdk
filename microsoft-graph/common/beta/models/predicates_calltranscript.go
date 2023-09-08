package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallTranscriptOperationPredicate struct {
	Content              *string
	CreatedDateTime      *string
	Id                   *string
	MeetingId            *string
	MeetingOrganizerId   *string
	MetadataContent      *string
	ODataType            *string
	TranscriptContentUrl *string
}

func (p CallTranscriptOperationPredicate) Matches(input CallTranscript) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MeetingId != nil && (input.MeetingId == nil || *p.MeetingId != *input.MeetingId) {
		return false
	}

	if p.MeetingOrganizerId != nil && (input.MeetingOrganizerId == nil || *p.MeetingOrganizerId != *input.MeetingOrganizerId) {
		return false
	}

	if p.MetadataContent != nil && (input.MetadataContent == nil || *p.MetadataContent != *input.MetadataContent) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TranscriptContentUrl != nil && (input.TranscriptContentUrl == nil || *p.TranscriptContentUrl != *input.TranscriptContentUrl) {
		return false
	}

	return true
}
