package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChatMessageAttachmentOperationPredicate struct {
	Content      *string
	ContentType  *string
	ContentUrl   *string
	Id           *string
	Name         *string
	ODataType    *string
	TeamsAppId   *string
	ThumbnailUrl *string
}

func (p ChatMessageAttachmentOperationPredicate) Matches(input ChatMessageAttachment) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.ContentType != nil && (input.ContentType == nil || *p.ContentType != *input.ContentType) {
		return false
	}

	if p.ContentUrl != nil && (input.ContentUrl == nil || *p.ContentUrl != *input.ContentUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TeamsAppId != nil && (input.TeamsAppId == nil || *p.TeamsAppId != *input.TeamsAppId) {
		return false
	}

	if p.ThumbnailUrl != nil && (input.ThumbnailUrl == nil || *p.ThumbnailUrl != *input.ThumbnailUrl) {
		return false
	}

	return true
}
