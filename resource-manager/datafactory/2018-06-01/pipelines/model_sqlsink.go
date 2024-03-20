package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlSink struct {
	DisableMetricsCollection              *interface{}       `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections              *interface{}       `json:"maxConcurrentConnections,omitempty"`
	PreCopyScript                         *interface{}       `json:"preCopyScript,omitempty"`
	SinkRetryCount                        *interface{}       `json:"sinkRetryCount,omitempty"`
	SinkRetryWait                         *interface{}       `json:"sinkRetryWait,omitempty"`
	SqlWriterStoredProcedureName          *interface{}       `json:"sqlWriterStoredProcedureName,omitempty"`
	SqlWriterTableType                    *interface{}       `json:"sqlWriterTableType,omitempty"`
	SqlWriterUseTableLock                 *interface{}       `json:"sqlWriterUseTableLock,omitempty"`
	StoredProcedureParameters             *interface{}       `json:"storedProcedureParameters,omitempty"`
	StoredProcedureTableTypeParameterName *interface{}       `json:"storedProcedureTableTypeParameterName,omitempty"`
	TableOption                           *interface{}       `json:"tableOption,omitempty"`
	Type                                  string             `json:"type"`
	UpsertSettings                        *SqlUpsertSettings `json:"upsertSettings,omitempty"`
	WriteBatchSize                        *interface{}       `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout                     *interface{}       `json:"writeBatchTimeout,omitempty"`
	WriteBehavior                         *interface{}       `json:"writeBehavior,omitempty"`
}
