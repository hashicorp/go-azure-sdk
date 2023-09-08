package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebAccountOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	Id                   *string
	IsSearchable         *bool
	LastModifiedDateTime *string
	ODataType            *string
	StatusMessage        *string
	ThumbnailUrl         *string
	UserId               *string
	WebUrl               *string
}

func (p WebAccountOperationPredicate) Matches(input WebAccount) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSearchable != nil && (input.IsSearchable == nil || *p.IsSearchable != *input.IsSearchable) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StatusMessage != nil && (input.StatusMessage == nil || *p.StatusMessage != *input.StatusMessage) {
		return false
	}

	if p.ThumbnailUrl != nil && (input.ThumbnailUrl == nil || *p.ThumbnailUrl != *input.ThumbnailUrl) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
