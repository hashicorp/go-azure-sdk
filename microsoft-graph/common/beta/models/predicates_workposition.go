package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkPositionOperationPredicate struct {
	CreatedDateTime      *string
	Id                   *string
	IsCurrent            *bool
	IsSearchable         *bool
	LastModifiedDateTime *string
	ODataType            *string
}

func (p WorkPositionOperationPredicate) Matches(input WorkPosition) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsCurrent != nil && (input.IsCurrent == nil || *p.IsCurrent != *input.IsCurrent) {
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

	return true
}
