package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceAnnouncementAttachmentOperationPredicate struct {
	Content              *string
	ContentType          *string
	Id                   *string
	LastModifiedDateTime *string
	Name                 *string
	ODataType            *string
	Size                 *int64
}

func (p ServiceAnnouncementAttachmentOperationPredicate) Matches(input ServiceAnnouncementAttachment) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.ContentType != nil && (input.ContentType == nil || *p.ContentType != *input.ContentType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	return true
}
