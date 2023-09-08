package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookIconOperationPredicate struct {
	Index     *int64
	ODataType *string
	Set       *string
}

func (p WorkbookIconOperationPredicate) Matches(input WorkbookIcon) bool {

	if p.Index != nil && (input.Index == nil || *p.Index != *input.Index) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Set != nil && (input.Set == nil || *p.Set != *input.Set) {
		return false
	}

	return true
}
