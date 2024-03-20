package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SapOpenHubSource struct {
	AdditionalColumns                *interface{} `json:"additionalColumns,omitempty"`
	BaseRequestId                    *interface{} `json:"baseRequestId,omitempty"`
	CustomRfcReadTableFunctionModule *interface{} `json:"customRfcReadTableFunctionModule,omitempty"`
	DisableMetricsCollection         *interface{} `json:"disableMetricsCollection,omitempty"`
	ExcludeLastRequest               *interface{} `json:"excludeLastRequest,omitempty"`
	MaxConcurrentConnections         *interface{} `json:"maxConcurrentConnections,omitempty"`
	QueryTimeout                     *interface{} `json:"queryTimeout,omitempty"`
	SapDataColumnDelimiter           *interface{} `json:"sapDataColumnDelimiter,omitempty"`
	SourceRetryCount                 *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait                  *interface{} `json:"sourceRetryWait,omitempty"`
	Type                             string       `json:"type"`
}
