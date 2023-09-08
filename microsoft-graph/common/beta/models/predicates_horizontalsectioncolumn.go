package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSectionColumnOperationPredicate struct {
	Id        *string
	ODataType *string
	Width     *int64
}

func (p HorizontalSectionColumnOperationPredicate) Matches(input HorizontalSectionColumn) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Width != nil && (input.Width == nil || *p.Width != *input.Width) {
		return false
	}

	return true
}
