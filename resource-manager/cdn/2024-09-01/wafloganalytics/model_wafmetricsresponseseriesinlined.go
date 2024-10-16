package wafloganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WafMetricsResponseSeriesInlined struct {
	Data   *[]WafMetricsResponseSeriesInlinedDataInlined   `json:"data,omitempty"`
	Groups *[]WafMetricsResponseSeriesInlinedGroupsInlined `json:"groups,omitempty"`
	Metric *string                                         `json:"metric,omitempty"`
	Unit   *WafMetricsSeriesUnit                           `json:"unit,omitempty"`
}
