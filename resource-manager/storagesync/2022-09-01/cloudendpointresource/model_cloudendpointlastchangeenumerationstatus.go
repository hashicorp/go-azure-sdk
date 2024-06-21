package cloudendpointresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudEndpointLastChangeEnumerationStatus struct {
	CompletedTimestamp        *string `json:"completedTimestamp,omitempty"`
	NamespaceDirectoriesCount *int64  `json:"namespaceDirectoriesCount,omitempty"`
	NamespaceFilesCount       *int64  `json:"namespaceFilesCount,omitempty"`
	NamespaceSizeBytes        *int64  `json:"namespaceSizeBytes,omitempty"`
	NextRunTimestamp          *string `json:"nextRunTimestamp,omitempty"`
	StartedTimestamp          *string `json:"startedTimestamp,omitempty"`
}
