package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintMarginOperationPredicate struct {
	Bottom    *int64
	Left      *int64
	ODataType *string
	Right     *int64
	Top       *int64
}

func (p PrintMarginOperationPredicate) Matches(input PrintMargin) bool {

	if p.Bottom != nil && (input.Bottom == nil || *p.Bottom != *input.Bottom) {
		return false
	}

	if p.Left != nil && (input.Left == nil || *p.Left != *input.Left) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Right != nil && (input.Right == nil || *p.Right != *input.Right) {
		return false
	}

	if p.Top != nil && (input.Top == nil || *p.Top != *input.Top) {
		return false
	}

	return true
}
