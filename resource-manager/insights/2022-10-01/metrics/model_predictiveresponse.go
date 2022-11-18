package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictiveResponse struct {
	Data             *[]PredictiveValue `json:"data,omitempty"`
	Interval         *string            `json:"interval,omitempty"`
	MetricName       *string            `json:"metricName,omitempty"`
	TargetResourceId *string            `json:"targetResourceId,omitempty"`
	Timespan         *string            `json:"timespan,omitempty"`
}
