package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlDWSink struct {
	AllowCopyCommand         *interface{}           `json:"allowCopyCommand,omitempty"`
	AllowPolyBase            *interface{}           `json:"allowPolyBase,omitempty"`
	CopyCommandSettings      *DWCopyCommandSettings `json:"copyCommandSettings,omitempty"`
	DisableMetricsCollection *interface{}           `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{}           `json:"maxConcurrentConnections,omitempty"`
	PolyBaseSettings         *PolybaseSettings      `json:"polyBaseSettings,omitempty"`
	PreCopyScript            *interface{}           `json:"preCopyScript,omitempty"`
	SinkRetryCount           *interface{}           `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{}           `json:"sinkRetryWait,omitempty"`
	SqlWriterUseTableLock    *interface{}           `json:"sqlWriterUseTableLock,omitempty"`
	TableOption              *interface{}           `json:"tableOption,omitempty"`
	Type                     string                 `json:"type"`
	UpsertSettings           *SqlDWUpsertSettings   `json:"upsertSettings,omitempty"`
	WriteBatchSize           *interface{}           `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{}           `json:"writeBatchTimeout,omitempty"`
	WriteBehavior            *interface{}           `json:"writeBehavior,omitempty"`
}
