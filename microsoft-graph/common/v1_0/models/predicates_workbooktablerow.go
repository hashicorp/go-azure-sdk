package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookTableRowOperationPredicate struct {
	Id        *string
	Index     *int64
	ODataType *string
}

func (p WorkbookTableRowOperationPredicate) Matches(input WorkbookTableRow) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Index != nil && (input.Index == nil || *p.Index != *input.Index) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
