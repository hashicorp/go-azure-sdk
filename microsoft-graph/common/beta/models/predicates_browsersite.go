package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteOperationPredicate struct {
	AllowRedirect        *bool
	Comment              *string
	CreatedDateTime      *string
	DeletedDateTime      *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	WebUrl               *string
}

func (p BrowserSiteOperationPredicate) Matches(input BrowserSite) bool {

	if p.AllowRedirect != nil && (input.AllowRedirect == nil || *p.AllowRedirect != *input.AllowRedirect) {
		return false
	}

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
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

	if p.WebUrl != nil && (input.WebUrl == nil || *p.WebUrl != *input.WebUrl) {
		return false
	}

	return true
}
