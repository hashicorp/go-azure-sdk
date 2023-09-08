package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefaultColumnValueOperationPredicate struct {
	Formula   *string
	ODataType *string
	Value     *string
}

func (p DefaultColumnValueOperationPredicate) Matches(input DefaultColumnValue) bool {

	if p.Formula != nil && (input.Formula == nil || *p.Formula != *input.Formula) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Value != nil && (input.Value == nil || *p.Value != *input.Value) {
		return false
	}

	return true
}
