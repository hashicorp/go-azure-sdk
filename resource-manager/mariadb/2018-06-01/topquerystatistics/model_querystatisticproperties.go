package topquerystatistics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryStatisticProperties struct {
	AggregationFunction *string   `json:"aggregationFunction,omitempty"`
	DatabaseNames       *[]string `json:"databaseNames,omitempty"`
	EndTime             *string   `json:"endTime,omitempty"`
	MetricDisplayName   *string   `json:"metricDisplayName,omitempty"`
	MetricName          *string   `json:"metricName,omitempty"`
	MetricValue         *float64  `json:"metricValue,omitempty"`
	MetricValueUnit     *string   `json:"metricValueUnit,omitempty"`
	QueryExecutionCount *int64    `json:"queryExecutionCount,omitempty"`
	QueryId             *string   `json:"queryId,omitempty"`
	StartTime           *string   `json:"startTime,omitempty"`
}
