package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationDefinitionOperationPredicate struct {
	IsDescending *bool
	MinimumCount *int64
	ODataType    *string
	PrefixFilter *string
}

func (p BucketAggregationDefinitionOperationPredicate) Matches(input BucketAggregationDefinition) bool {

	if p.IsDescending != nil && (input.IsDescending == nil || *p.IsDescending != *input.IsDescending) {
		return false
	}

	if p.MinimumCount != nil && (input.MinimumCount == nil || *p.MinimumCount != *input.MinimumCount) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrefixFilter != nil && (input.PrefixFilter == nil || *p.PrefixFilter != *input.PrefixFilter) {
		return false
	}

	return true
}
