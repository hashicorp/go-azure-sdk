package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteListOperationPredicate struct {
	Description          *string
	DisplayName          *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	PublishedDateTime    *string
	Revision             *string
}

func (p BrowserSiteListOperationPredicate) Matches(input BrowserSiteList) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
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

	if p.PublishedDateTime != nil && (input.PublishedDateTime == nil || *p.PublishedDateTime != *input.PublishedDateTime) {
		return false
	}

	if p.Revision != nil && (input.Revision == nil || *p.Revision != *input.Revision) {
		return false
	}

	return true
}
