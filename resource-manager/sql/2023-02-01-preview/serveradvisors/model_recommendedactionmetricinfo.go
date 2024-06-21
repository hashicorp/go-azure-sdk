package serveradvisors

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecommendedActionMetricInfo struct {
	MetricName *string  `json:"metricName,omitempty"`
	StartTime  *string  `json:"startTime,omitempty"`
	TimeGrain  *string  `json:"timeGrain,omitempty"`
	Unit       *string  `json:"unit,omitempty"`
	Value      *float64 `json:"value,omitempty"`
}
