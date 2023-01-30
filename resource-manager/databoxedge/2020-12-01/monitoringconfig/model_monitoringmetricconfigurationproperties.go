package monitoringconfig

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringMetricConfigurationProperties struct {
	MetricConfigurations []MetricConfiguration `json:"metricConfigurations"`
}
