package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BlobSource struct {
	DisableMetricsCollection *interface{} `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{} `json:"maxConcurrentConnections,omitempty"`
	Recursive                *interface{} `json:"recursive,omitempty"`
	SkipHeaderLineCount      *interface{} `json:"skipHeaderLineCount,omitempty"`
	SourceRetryCount         *interface{} `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{} `json:"sourceRetryWait,omitempty"`
	TreatEmptyAsNull         *interface{} `json:"treatEmptyAsNull,omitempty"`
	Type                     string       `json:"type"`
}
