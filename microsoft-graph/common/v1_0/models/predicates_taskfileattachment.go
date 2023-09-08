package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskFileAttachmentOperationPredicate struct {
	ContentBytes         *string
	ContentType          *string
	Id                   *string
	LastModifiedDateTime *string
	Name                 *string
	ODataType            *string
	Size                 *int64
}

func (p TaskFileAttachmentOperationPredicate) Matches(input TaskFileAttachment) bool {

	if p.ContentBytes != nil && (input.ContentBytes == nil || *p.ContentBytes != *input.ContentBytes) {
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
