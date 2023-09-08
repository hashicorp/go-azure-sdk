package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSiteHistoryOperationPredicate struct {
	AllowRedirect     *bool
	Comment           *string
	ODataType         *string
	PublishedDateTime *string
}

func (p BrowserSiteHistoryOperationPredicate) Matches(input BrowserSiteHistory) bool {

	if p.AllowRedirect != nil && (input.AllowRedirect == nil || *p.AllowRedirect != *input.AllowRedirect) {
		return false
	}

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PublishedDateTime != nil && (input.PublishedDateTime == nil || *p.PublishedDateTime != *input.PublishedDateTime) {
		return false
	}

	return true
}
