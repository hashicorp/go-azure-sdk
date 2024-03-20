package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SnowflakeSource struct {
	DisableMetricsCollection *interface{}               `json:"disableMetricsCollection,omitempty"`
	ExportSettings           SnowflakeExportCopyCommand `json:"exportSettings"`
	MaxConcurrentConnections *interface{}               `json:"maxConcurrentConnections,omitempty"`
	Query                    *interface{}               `json:"query,omitempty"`
	SourceRetryCount         *interface{}               `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{}               `json:"sourceRetryWait,omitempty"`
	Type                     string                     `json:"type"`
}
