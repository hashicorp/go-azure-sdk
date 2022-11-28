package entities

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EntityInsightItem struct {
	ChartQueryResults *[]InsightsTableResult              `json:"chartQueryResults,omitempty"`
	QueryId           *string                             `json:"queryId,omitempty"`
	QueryTimeInterval *EntityInsightItemQueryTimeInterval `json:"queryTimeInterval"`
	TableQueryResults *InsightsTableResult                `json:"tableQueryResults"`
}
