package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DynamicsCrmSink struct {
	AlternateKeyName         *interface{}              `json:"alternateKeyName,omitempty"`
	DisableMetricsCollection *interface{}              `json:"disableMetricsCollection,omitempty"`
	IgnoreNullValues         *interface{}              `json:"ignoreNullValues,omitempty"`
	MaxConcurrentConnections *interface{}              `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *interface{}              `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{}              `json:"sinkRetryWait,omitempty"`
	Type                     string                    `json:"type"`
	WriteBatchSize           *interface{}              `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{}              `json:"writeBatchTimeout,omitempty"`
	WriteBehavior            DynamicsSinkWriteBehavior `json:"writeBehavior"`
}
