package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SftpWriteSettings struct {
	CopyBehavior             *interface{}    `json:"copyBehavior,omitempty"`
	DisableMetricsCollection *interface{}    `json:"disableMetricsCollection,omitempty"`
	MaxConcurrentConnections *interface{}    `json:"maxConcurrentConnections,omitempty"`
	Metadata                 *[]MetadataItem `json:"metadata,omitempty"`
	OperationTimeout         *interface{}    `json:"operationTimeout,omitempty"`
	Type                     string          `json:"type"`
	UseTempFileRename        *interface{}    `json:"useTempFileRename,omitempty"`
}
