package query

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryDataset struct {
	Aggregation   *map[string]QueryAggregation `json:"aggregation,omitempty"`
	Configuration *QueryDatasetConfiguration   `json:"configuration"`
	Filter        *QueryFilter                 `json:"filter"`
	Granularity   *GranularityType             `json:"granularity,omitempty"`
	Grouping      *[]QueryGrouping             `json:"grouping,omitempty"`
}
