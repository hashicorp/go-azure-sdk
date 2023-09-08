package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IpApplicationSegmentOperationPredicate struct {
	DestinationHost *string
	Id              *string
	ODataType       *string
	Port            *int64
}

func (p IpApplicationSegmentOperationPredicate) Matches(input IpApplicationSegment) bool {

	if p.DestinationHost != nil && (input.DestinationHost == nil || *p.DestinationHost != *input.DestinationHost) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Port != nil && (input.Port == nil || *p.Port != *input.Port) {
		return false
	}

	return true
}
