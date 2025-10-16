package datacollectionrules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DirectDataSourcesSpec struct {
	OtelLogs    *[]OtelLogsDirectDataSource    `json:"otelLogs,omitempty"`
	OtelMetrics *[]OtelMetricsDirectDataSource `json:"otelMetrics,omitempty"`
	OtelTraces  *[]OtelTracesDirectDataSource  `json:"otelTraces,omitempty"`
}
