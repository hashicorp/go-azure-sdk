package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumberRangeOperationPredicate struct {
	LowerNumber *int64
	ODataType   *string
	UpperNumber *int64
}

func (p NumberRangeOperationPredicate) Matches(input NumberRange) bool {

	if p.LowerNumber != nil && (input.LowerNumber == nil || *p.LowerNumber != *input.LowerNumber) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UpperNumber != nil && (input.UpperNumber == nil || *p.UpperNumber != *input.UpperNumber) {
		return false
	}

	return true
}
