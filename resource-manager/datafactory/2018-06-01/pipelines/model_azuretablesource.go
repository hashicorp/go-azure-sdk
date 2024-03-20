package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureTableSource struct {
	AdditionalColumns                   *interface{} `json:"additionalColumns,omitempty"`
	AzureTableSourceIgnoreTableNotFound *interface{} `json:"azureTableSourceIgnoreTableNotFound,omitempty"`
	AzureTableSourceQuery               *interface{} `json:"azureTableSourceQuery,omitempty"`
	DisableMetricsCollection            *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections            *interface{} `json:"maxConcurrentConnections,omitempty"`
	QueryTimeout                        *interface{} `json:"queryTimeout,omitempty"`
	SourceRetryCount                    *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait                     *interface{} `json:"sourceRetryWait,omitempty"`
	Type                                string       `json:"type"`
}
