package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SalesforceServiceCloudV2Sink struct {
	DisableMetricsCollection *interface{}                   `json:"disableMetricsCollection,omitempty"`
	ExternalIdFieldName      *interface{}                   `json:"externalIdFieldName,omitempty"`
	IgnoreNullValues         *interface{}                   `json:"ignoreNullValues,omitempty"`
	MaxConcurrentConnections *interface{}                   `json:"maxConcurrentConnections,omitempty"`
	SinkRetryCount           *interface{}                   `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{}                   `json:"sinkRetryWait,omitempty"`
	Type                     string                         `json:"type"`
	WriteBatchSize           *interface{}                   `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{}                   `json:"writeBatchTimeout,omitempty"`
	WriteBehavior            *SalesforceV2SinkWriteBehavior `json:"writeBehavior,omitempty"`
}
