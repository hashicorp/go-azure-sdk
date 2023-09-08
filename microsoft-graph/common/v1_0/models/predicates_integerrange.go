package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IntegerRangeOperationPredicate struct {
	End       *int64
	ODataType *string
	Start     *int64
}

func (p IntegerRangeOperationPredicate) Matches(input IntegerRange) bool {

	if p.End != nil && (input.End == nil || *p.End != *input.End) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Start != nil && (input.Start == nil || *p.Start != *input.Start) {
		return false
	}

	return true
}
