package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SapTableSource struct {
	AdditionalColumns                *interface{}               `json:"additionalColumns,omitempty"`
	BatchSize                        *interface{}               `json:"batchSize,omitempty"`
	CustomRfcReadTableFunctionModule *interface{}               `json:"customRfcReadTableFunctionModule,omitempty"`
	DisableMetricsCollection         *interface{}               `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections         *interface{}               `json:"maxConcurrentConnections,omitempty"`
	PartitionOption                  *interface{}               `json:"partitionOption,omitempty"`
	PartitionSettings                *SapTablePartitionSettings `json:"partitionSettings,omitempty"`
	QueryTimeout                     *interface{}               `json:"queryTimeout,omitempty"`
	RfcTableFields                   *interface{}               `json:"rfcTableFields,omitempty"`
	RfcTableOptions                  *interface{}               `json:"rfcTableOptions,omitempty"`
	RowCount                         *interface{}               `json:"rowCount,omitempty"`
	RowSkips                         *interface{}               `json:"rowSkips,omitempty"`
	SapDataColumnDelimiter           *interface{}               `json:"sapDataColumnDelimiter,omitempty"`
	SourceRetryCount                 *interface{}               `json:"sourceRetryCount,omitempty"`
	SourceRetryWait                  *interface{}               `json:"sourceRetryWait,omitempty"`
	Type                             string                     `json:"type"`
}
