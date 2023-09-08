package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPublicationOperationPredicate struct {
	CreatedDateTime      *string
	Description          *string
	DisplayName          *string
	Id                   *string
	IsSearchable         *bool
	LastModifiedDateTime *string
	ODataType            *string
	PublishedDate        *string
	Publisher            *string
	ThumbnailUrl         *string
	WebUrl               *string
}

func (p ItemPublicationOperationPredicate) Matches(input ItemPublication) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.PublishedDate != nil && (input.PublishedDate == nil || *p.PublishedDate != *input.PublishedDate) {
		return false
	}

	if p.Publisher != nil && (input.Publisher == nil || *p.Publisher != *input.Publisher) {
		return false
	}

	if p.ThumbnailUrl != nil && (input.ThumbnailUrl == nil || *p.ThumbnailUrl != *input.ThumbnailUrl) {
		return false
	}

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
