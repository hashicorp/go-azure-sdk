package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieHistoryOperationPredicate struct {
	Comment           *string
	DisplayName       *string
	HostOnly          *bool
	HostOrDomain      *string
	ODataType         *string
	Path              *string
	PublishedDateTime *string
}

func (p BrowserSharedCookieHistoryOperationPredicate) Matches(input BrowserSharedCookieHistory) bool {

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HostOnly != nil && (input.HostOnly == nil || *p.HostOnly != *input.HostOnly) {
		return false
	}

	if p.HostOrDomain != nil && (input.HostOrDomain == nil || *p.HostOrDomain != *input.HostOrDomain) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Path != nil && (input.Path == nil || *p.Path != *input.Path) {
		return false
	}

	if p.PublishedDateTime != nil && (input.PublishedDateTime == nil || *p.PublishedDateTime != *input.PublishedDateTime) {
		return false
	}

	return true
}
