package pipelines

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OracleCloudStorageReadSettings struct {
	DeleteFilesAfterCompletion *interface{} `json:"deleteFilesAfterCompletion,omitempty"`
	DisableMetricsCollection   *interface{} `json:"disableMetricsCollection,omitempty"`
	EnablePartitionDiscovery   *interface{} `json:"enablePartitionDiscovery,omitempty"`
	FileListPath               *interface{} `json:"fileListPath,omitempty"`
	MaxConcurrentConnections   *interface{} `json:"maxConcurrentConnections,omitempty"`
	ModifiedDatetimeEnd        *interface{} `json:"modifiedDatetimeEnd,omitempty"`
	ModifiedDatetimeStart      *interface{} `json:"modifiedDatetimeStart,omitempty"`
	PartitionRootPath          *interface{} `json:"partitionRootPath,omitempty"`
	Prefix                     *interface{} `json:"prefix,omitempty"`
	Recursive                  *interface{} `json:"recursive,omitempty"`
	Type                       string       `json:"type"`
	WildcardFileName           *interface{} `json:"wildcardFileName,omitempty"`
	WildcardFolderPath         *interface{} `json:"wildcardFolderPath,omitempty"`
}
