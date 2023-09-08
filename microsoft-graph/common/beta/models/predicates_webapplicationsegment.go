package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebApplicationSegmentOperationPredicate struct {
	AlternateUrl *string
	ExternalUrl  *string
	Id           *string
	InternalUrl  *string
	ODataType    *string
}

func (p WebApplicationSegmentOperationPredicate) Matches(input WebApplicationSegment) bool {

	if p.AlternateUrl != nil && (input.AlternateUrl == nil || *p.AlternateUrl != *input.AlternateUrl) {
		return false
	}

	if p.ExternalUrl != nil && (input.ExternalUrl == nil || *p.ExternalUrl != *input.ExternalUrl) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InternalUrl != nil && (input.InternalUrl == nil || *p.InternalUrl != *input.InternalUrl) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
