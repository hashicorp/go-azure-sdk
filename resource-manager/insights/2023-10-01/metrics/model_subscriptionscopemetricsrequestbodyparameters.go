package metrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubscriptionScopeMetricsRequestBodyParameters struct {
	Aggregation         *string           `json:"aggregation,omitempty"`
	AutoAdjustTimegrain *bool             `json:"autoAdjustTimegrain,omitempty"`
	Filter              *string           `json:"filter,omitempty"`
	Interval            *string           `json:"interval,omitempty"`
	MetricNames         *string           `json:"metricNames,omitempty"`
	MetricNamespace     *string           `json:"metricNamespace,omitempty"`
	OrderBy             *string           `json:"orderBy,omitempty"`
	ResultType          *MetricResultType `json:"resultType,omitempty"`
	RollUpBy            *string           `json:"rollUpBy,omitempty"`
	Timespan            *string           `json:"timespan,omitempty"`
	Top                 *int64            `json:"top,omitempty"`
	ValidateDimensions  *bool             `json:"validateDimensions,omitempty"`
}
