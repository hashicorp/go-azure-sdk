package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationDefinition struct {
	IsDescending *bool                              `json:"isDescending,omitempty"`
	MinimumCount *int64                             `json:"minimumCount,omitempty"`
	ODataType    *string                            `json:"@odata.type,omitempty"`
	PrefixFilter *string                            `json:"prefixFilter,omitempty"`
	Ranges       *[]BucketAggregationRange          `json:"ranges,omitempty"`
	SortBy       *BucketAggregationDefinitionSortBy `json:"sortBy,omitempty"`
}
