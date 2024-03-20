package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FtpReadSettings struct {
	DeleteFilesAfterCompletion *interface{} `json:"deleteFilesAfterCompletion,omitempty"`
	DisableChunking            *interface{} `json:"disableChunking,omitempty"`
	DisableMetricsCollection   *interface{} `json:"disableMetricsCollection,omitempty"`
	EnablePartitionDiscovery   *interface{} `json:"enablePartitionDiscovery,omitempty"`
	FileListPath               *interface{} `json:"fileListPath,omitempty"`
	MaxConcurrentConnections   *interface{} `json:"maxConcurrentConnections,omitempty"`
	PartitionRootPath          *interface{} `json:"partitionRootPath,omitempty"`
	Recursive                  *interface{} `json:"recursive,omitempty"`
	Type                       string       `json:"type"`
	UseBinaryTransfer          *interface{} `json:"useBinaryTransfer,omitempty"`
	WildcardFileName           *interface{} `json:"wildcardFileName,omitempty"`
	WildcardFolderPath         *interface{} `json:"wildcardFolderPath,omitempty"`
}
