package diagnostics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenTelemetryConfiguration struct {
	DestinationsConfiguration *DestinationsConfiguration `json:"destinationsConfiguration,omitempty"`
	LogsConfiguration         *LogsConfiguration         `json:"logsConfiguration,omitempty"`
	MetricsConfiguration      *MetricsConfiguration      `json:"metricsConfiguration,omitempty"`
	TracesConfiguration       *TracesConfiguration       `json:"tracesConfiguration,omitempty"`
}
