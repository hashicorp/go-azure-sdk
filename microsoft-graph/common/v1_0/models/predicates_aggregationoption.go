package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AggregationOptionOperationPredicate struct {
	Field     *string
	ODataType *string
	Size      *int64
}

func (p AggregationOptionOperationPredicate) Matches(input AggregationOption) bool {

	if p.Field != nil && (input.Field == nil || *p.Field != *input.Field) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Size != nil && (input.Size == nil || *p.Size != *input.Size) {
		return false
	}

	return true
}
