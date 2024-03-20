package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureTableSink struct {
	AzureTableDefaultPartitionKeyValue *interface{} `json:"azureTableDefaultPartitionKeyValue,omitempty"`
	AzureTableInsertType               *interface{} `json:"azureTableInsertType,omitempty"`
	AzureTablePartitionKeyName         *interface{} `json:"azureTablePartitionKeyName,omitempty"`
	AzureTableRowKeyName               *interface{} `json:"azureTableRowKeyName,omitempty"`
	DisableMetricsCollection           *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections           *interface{} `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount                     *interface{} `json:"sinkRetryCount,omitempty"`
	SinkRetryWait                      *interface{} `json:"sinkRetryWait,omitempty"`
	Type                               string       `json:"type"`
	WriteBatchSize                     *interface{} `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout                  *interface{} `json:"writeBatchTimeout,omitempty"`
}
