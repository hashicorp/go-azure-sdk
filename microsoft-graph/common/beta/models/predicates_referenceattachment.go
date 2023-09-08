package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReferenceAttachmentOperationPredicate struct {
	ContentType          *string
	Id                   *string
	IsFolder             *bool
	IsInline             *bool
	LastModifiedDateTime *string
	Name                 *string
	ODataType            *string
	PreviewUrl           *string
	Size                 *int64
	SourceUrl            *string
	ThumbnailUrl         *string
}

func (p ReferenceAttachmentOperationPredicate) Matches(input ReferenceAttachment) bool {

	if p.ContentType != nil && (input.ContentType == nil || *p.ContentType != *input.ContentType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsFolder != nil && (input.IsFolder == nil || *p.IsFolder != *input.IsFolder) {
		return false
	}

	if p.IsInline != nil && (input.IsInline == nil || *p.IsInline != *input.IsInline) {
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

	if p.PreviewUrl != nil && (input.PreviewUrl == nil || *p.PreviewUrl != *input.PreviewUrl) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	if p.SourceUrl != nil && (input.SourceUrl == nil || *p.SourceUrl != *input.SourceUrl) {
		return false
	}

	if p.ThumbnailUrl != nil && (input.ThumbnailUrl == nil || *p.ThumbnailUrl != *input.ThumbnailUrl) {
		return false
	}

	return true
}
