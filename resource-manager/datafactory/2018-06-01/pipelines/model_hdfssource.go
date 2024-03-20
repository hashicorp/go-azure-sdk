package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HdfsSource struct {
	DisableMetricsCollection *interface{}    `json:"disableMetricsCollection,omitempty"`
	DistcpSettings           *DistcpSettings `json:"distcpSettings,omitempty"`
	MaxConcurrentConnections *interface{}    `json:"maxConcurrentConnections,omitempty"`
	Recursive                *interface{}    `json:"recursive,omitempty"`
	SourceRetryCount         *interface{}    `json:"sourceRetryCount,omitempty"`
	SourceRetryWait          *interface{}    `json:"sourceRetryWait,omitempty"`
	Type                     string          `json:"type"`
}
