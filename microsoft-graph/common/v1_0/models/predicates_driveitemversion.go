package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DriveItemVersionOperationPredicate struct {
	Content              *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	Size                 *int64
}

func (p DriveItemVersionOperationPredicate) Matches(input DriveItemVersion) bool {

	if p.Content != nil && (input.Content == nil || *p.Content != *input.Content) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
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
