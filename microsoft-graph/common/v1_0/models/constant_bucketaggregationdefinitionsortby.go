package models

import "strings"

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationDefinitionSortBy string

const (
	BucketAggregationDefinitionSortBycount       BucketAggregationDefinitionSortBy = "Count"
	BucketAggregationDefinitionSortBykeyAsNumber BucketAggregationDefinitionSortBy = "KeyAsNumber"
	BucketAggregationDefinitionSortBykeyAsString BucketAggregationDefinitionSortBy = "KeyAsString"
)

func PossibleValuesForBucketAggregationDefinitionSortBy() []string {
	return []string{
		string(BucketAggregationDefinitionSortBycount),
		string(BucketAggregationDefinitionSortBykeyAsNumber),
		string(BucketAggregationDefinitionSortBykeyAsString),
	}
}

func parseBucketAggregationDefinitionSortBy(input string) (*BucketAggregationDefinitionSortBy, error) {
	vals := map[string]BucketAggregationDefinitionSortBy{
		"count":       BucketAggregationDefinitionSortBycount,
		"keyasnumber": BucketAggregationDefinitionSortBykeyAsNumber,
		"keyasstring": BucketAggregationDefinitionSortBykeyAsString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BucketAggregationDefinitionSortBy(input)
	return &out, nil
}
