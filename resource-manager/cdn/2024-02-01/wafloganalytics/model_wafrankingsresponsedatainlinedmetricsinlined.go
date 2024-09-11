package wafloganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WafRankingsResponseDataInlinedMetricsInlined struct {
	Metric     *string  `json:"metric,omitempty"`
	Percentage *float64 `json:"percentage,omitempty"`
	Value      *int64   `json:"value,omitempty"`
}
