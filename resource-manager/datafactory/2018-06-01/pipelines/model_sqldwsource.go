package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlDWSource struct {
	AdditionalColumns            *interface{}          `json:"additionalColumns,omitempty"`
	DisableMetricsCollection     *interface{}          `json:"disableMetricsCollection,omitempty"`
	IsolationLevel               *interface{}          `json:"isolationLevel,omitempty"`
	MaxConcurrentConnections     *interface{}          `json:"maxConcurrentConnections,omitempty"`
	PartitionOption              *interface{}          `json:"partitionOption,omitempty"`
	PartitionSettings            *SqlPartitionSettings `json:"partitionSettings,omitempty"`
	QueryTimeout                 *interface{}          `json:"queryTimeout,omitempty"`
	SourceRetryCount             *interface{}          `json:"sourceRetryCount,omitempty"`
	SourceRetryWait              *interface{}          `json:"sourceRetryWait,omitempty"`
	SqlReaderQuery               *interface{}          `json:"sqlReaderQuery,omitempty"`
	SqlReaderStoredProcedureName *interface{}          `json:"sqlReaderStoredProcedureName,omitempty"`
	StoredProcedureParameters    *interface{}          `json:"storedProcedureParameters,omitempty"`
	Type                         string                `json:"type"`
}
