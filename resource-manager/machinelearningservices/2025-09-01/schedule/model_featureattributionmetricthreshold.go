package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FeatureAttributionMetricThreshold struct {
	Metric    FeatureAttributionMetric `json:"metric"`
	Threshold *MonitoringThreshold     `json:"threshold,omitempty"`
}
