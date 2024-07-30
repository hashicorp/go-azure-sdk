package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BucketAggregationDefinitionSortBy string

const (
	BucketAggregationDefinitionSortBy_Count       BucketAggregationDefinitionSortBy = "count"
	BucketAggregationDefinitionSortBy_KeyAsNumber BucketAggregationDefinitionSortBy = "keyAsNumber"
	BucketAggregationDefinitionSortBy_KeyAsString BucketAggregationDefinitionSortBy = "keyAsString"
)

func PossibleValuesForBucketAggregationDefinitionSortBy() []string {
	return []string{
		string(BucketAggregationDefinitionSortBy_Count),
		string(BucketAggregationDefinitionSortBy_KeyAsNumber),
		string(BucketAggregationDefinitionSortBy_KeyAsString),
	}
}

func (s *BucketAggregationDefinitionSortBy) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseBucketAggregationDefinitionSortBy(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseBucketAggregationDefinitionSortBy(input string) (*BucketAggregationDefinitionSortBy, error) {
	vals := map[string]BucketAggregationDefinitionSortBy{
		"count":       BucketAggregationDefinitionSortBy_Count,
		"keyasnumber": BucketAggregationDefinitionSortBy_KeyAsNumber,
		"keyasstring": BucketAggregationDefinitionSortBy_KeyAsString,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := BucketAggregationDefinitionSortBy(input)
	return &out, nil
}
