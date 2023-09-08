package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BrowserSharedCookieOperationPredicate struct {
	Comment              *string
	CreatedDateTime      *string
	DeletedDateTime      *string
	DisplayName          *string
	HostOnly             *bool
	HostOrDomain         *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	Path                 *string
}

func (p BrowserSharedCookieOperationPredicate) Matches(input BrowserSharedCookie) bool {

	if p.Comment != nil && (input.Comment == nil || *p.Comment != *input.Comment) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeletedDateTime != nil && (input.DeletedDateTime == nil || *p.DeletedDateTime != *input.DeletedDateTime) {
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

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Path != nil && (input.Path == nil || *p.Path != *input.Path) {
		return false
	}

	return true
}
