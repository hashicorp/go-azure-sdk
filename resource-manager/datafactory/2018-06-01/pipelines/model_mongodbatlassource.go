package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MongoDbAtlasSource struct {
	AdditionalColumns        *interface{}                    `json:"additionalColumns,omitempty"`
	BatchSize                *interface{}                    `json:"batchSize,omitempty"`
	CursorMethods            *MongoDbCursorMethodsProperties `json:"cursorMethods,omitempty"`
	DisableMetricsCollection *interface{}                    `json:"disableMetricsCollection,omitempty"`
	Filter                   *interface{}                    `json:"filter,omitempty"`
	MaxConcurrentConnections *interface{}                    `json:"maxConcurrentConnections,omitempty"`
	QueryTimeout             *interface{}                    `json:"queryTimeout,omitempty"`
	SourceRetryCount         *interface{}                    `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{}                    `json:"sourceRetryWait,omitempty"`
	Type                     string                          `json:"type"`
}
