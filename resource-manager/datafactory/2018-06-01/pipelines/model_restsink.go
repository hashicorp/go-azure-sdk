package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestSink struct {
	AdditionalHeaders        *interface{} `json:"additionalHeaders,omitempty"`
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	HTTPCompressionType      *interface{} `json:"httpCompressionType,omitempty"`
	HTTPRequestTimeout       *interface{} `json:"httpRequestTimeout,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	RequestInterval          *interface{} `json:"requestInterval,omitempty"`
	RequestMethod            *interface{} `json:"requestMethod,omitempty"`
	SinkRetryCount           *interface{} `json:"sinkRetryCount,omitempty"`
	SinkRetryWait            *interface{} `json:"sinkRetryWait,omitempty"`
	Type                     string       `json:"type"`
	WriteBatchSize           *interface{} `json:"writeBatchSize,omitempty"`
	WriteBatchTimeout        *interface{} `json:"writeBatchTimeout,omitempty"`
}
