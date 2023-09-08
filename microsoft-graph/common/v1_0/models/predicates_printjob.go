package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintJobOperationPredicate struct {
	CreatedDateTime *string
	Id              *string
	IsFetchable     *bool
	ODataType       *string
	RedirectedFrom  *string
	RedirectedTo    *string
}

func (p PrintJobOperationPredicate) Matches(input PrintJob) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsFetchable != nil && (input.IsFetchable == nil || *p.IsFetchable != *input.IsFetchable) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RedirectedFrom != nil && (input.RedirectedFrom == nil || *p.RedirectedFrom != *input.RedirectedFrom) {
		return false
	}

	if p.RedirectedTo != nil && (input.RedirectedTo == nil || *p.RedirectedTo != *input.RedirectedTo) {
		return false
	}

	return true
}
