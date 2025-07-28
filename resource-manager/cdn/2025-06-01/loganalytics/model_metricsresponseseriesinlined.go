package loganalytics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricsResponseSeriesInlined struct {
	Data   *[]MetricsResponseSeriesInlinedDataInlined   `json:"data,omitempty"`
	Groups *[]MetricsResponseSeriesInlinedGroupsInlined `json:"groups,omitempty"`
	Metric *string                                      `json:"metric,omitempty"`
	Unit   *MetricsSeriesUnit                           `json:"unit,omitempty"`
}
