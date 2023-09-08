package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SearchBucketOperationPredicate struct {
	AggregationFilterToken *string
	Count                  *int64
	Key                    *string
	ODataType              *string
}

func (p SearchBucketOperationPredicate) Matches(input SearchBucket) bool {

	if p.AggregationFilterToken != nil && (input.AggregationFilterToken == nil || *p.AggregationFilterToken != *input.AggregationFilterToken) {
		return false
	}

	if p.Count != nil && (input.Count == nil || *p.Count != *input.Count) {
		return false
	}

	if p.Key != nil && (input.Key == nil || *p.Key != *input.Key) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
