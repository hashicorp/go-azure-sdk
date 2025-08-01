package monitoredsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MetricRules struct {
	FilteringTags  *[]FilteringTag       `json:"filteringTags,omitempty"`
	SendingMetrics *SendingMetricsStatus `json:"sendingMetrics,omitempty"`
}
