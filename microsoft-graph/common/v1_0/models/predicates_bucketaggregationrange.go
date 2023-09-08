package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationRangeOperationPredicate struct {
	From      *string
	ODataType *string
	To        *string
}

func (p BucketAggregationRangeOperationPredicate) Matches(input BucketAggregationRange) bool {

	if p.From != nil && (input.From == nil || *p.From != *input.From) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.To != nil && (input.To == nil || *p.To != *input.To) {
		return false
	}

	return true
}
