package datasets

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureBlobFSLocation struct {
	FileName   *interface{} `json:"fileName,omitempty"`
	FileSystem *interface{} `json:"fileSystem,omitempty"`
	FolderPath *interface{} `json:"folderPath,omitempty"`
	Type       string       `json:"type"`
}
